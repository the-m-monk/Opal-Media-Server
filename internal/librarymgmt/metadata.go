package librarymgmt

import (
	"encoding/json"
	"errors"
	"fmt"
	"opal/internal/config"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

var tmdbApiKey string
var MetadataDir string
var tvshowMovieDetectionRegex = regexp.MustCompile(`\[(movie|tvshow)-imdb-(tt\d+)\]`)

func initMetadata() {
	tmdbApiKey = config.FetchValue("/tmdb.cfg", "api_key", true)

	if strings.HasPrefix(tmdbApiKey, "$") {
		tmdbApiKey = os.Getenv(strings.TrimPrefix(tmdbApiKey, "$"))
	}

	MetadataDir = config.FetchValue("/server.cfg", "metadata_dir", true)
	if err := os.MkdirAll(MetadataDir, 0755); err != nil {
		fmt.Println("Error: failed to mkdir ")
		os.Exit(1)
	}

	buildLibraryTree()

	for _, rootNode := range LibTree {
		slices.SortFunc(NewestMedia[rootNode.RootUuid], func(a, b MtimeSortEntry) int {
			return int(b.Mtime.Unix() - a.Mtime.Unix())
		})
	}

}

func buildLibraryTree() {
	for i, rootLib := range AllLibraries {
		fmt.Printf("Scanning %s\n", rootLib.DisplayName)

		rootStat, err := os.Stat(rootLib.Path)
		if err != nil {
			fmt.Printf("Error: failed to parse library @ %s: %s", rootLib.Path, err.Error())
			os.Exit(1)
		}

		if !rootStat.IsDir() {
			fmt.Printf("Error: %s is not a directory", rootLib.Path)
			os.Exit(1)
		}

		uuidNamespace := uuid.MustParse(uuidNamespaceString)
		newLibNamespace := uuid.NewSHA1(uuidNamespace, []byte(rootLib.DisplayName))

		newLibNode := searchDir(rootLib.Path, "/", newLibNamespace)
		newLibNode.Name = rootLib.DisplayName

		AllLibraries[i].Uuid = newLibNamespace.String()
		AllLibrariesMap[newLibNamespace.String()] = AllLibraries[i]
		LibTreeMap[newLibNamespace.String()] = newLibNode

		LibTree = append(LibTree, newLibNode)
	}
}

func searchDir(libRootPath string, relativePath string, libNamespace uuid.UUID) *TreeNode {
	fmt.Println(relativePath)

	nodeId := uuid.NewSHA1(libNamespace, []byte(filepath.ToSlash(relativePath))).String()

	retNode := &TreeNode{
		Path:     relativePath,
		Name:     filepath.Base(relativePath),
		Uuid:     nodeId,
		RootUuid: libNamespace.String(),
		Type:     "folder",
	}

	LibTreeMap[nodeId] = retNode

	absPath := filepath.Join(libRootPath, relativePath)
	subFiles, err := os.ReadDir(absPath)
	if err != nil {
		fmt.Printf("Error: failed to list directory: %s\n", err.Error())
		return retNode
	}

	for _, sF := range subFiles {
		childRelPath := filepath.Join(relativePath, sF.Name())
		/*
			metadataSubstrings[0] == full substring ie: [movie-imdb-<imdb id>]
			metadataSubstrings[1] == "movie" || "tvshow"
			metadataSubstrings[2] == imdb id
		*/
		metadataSubstrings := tvshowMovieDetectionRegex.FindStringSubmatch(sF.Name())

		if metadataSubstrings == nil && sF.IsDir() {
			retNode.Children = append(retNode.Children, searchDir(libRootPath, childRelPath, libNamespace))
			continue
		}

		if metadataSubstrings == nil { //Caused by random file contaminants in library, ignore
			continue
		}

		if len(metadataSubstrings) != 3 ||
			!(metadataSubstrings[1] == "movie" || metadataSubstrings[1] == "tvshow") ||
			!strings.HasPrefix(metadataSubstrings[2], "tt") {
			continue
		}

		fullPath := filepath.Join(libRootPath, childRelPath)
		st, _ := os.Stat(fullPath)

		child := &TreeNode{
			Name:     metadataSubstrings[0],
			Type:     metadataSubstrings[1],
			Uuid:     uuid.NewSHA1(libNamespace, []byte(filepath.ToSlash(childRelPath))).String(),
			Path:     childRelPath,
			RootUuid: string(libNamespace.String()),
			Mtime:    st.ModTime(),
			ImdbId:   metadataSubstrings[2],
		}

		childNewestMediaEntry := MtimeSortEntry{
			Node:  child,
			Mtime: child.Mtime,
		}

		NewestMedia[child.RootUuid] = append(NewestMedia[child.RootUuid], childNewestMediaEntry)
		retNode.Children = append(retNode.Children, child)
		LibTreeMap[child.Uuid] = child

		fetchMetadata(metadataSubstrings[2], child)
	}

	return retNode
}

// TODO: TMDB api terms of use 1.C.iii: must refresh metadata every 6 months
func fetchMetadata(imdbId string, item *TreeNode) {
	MetadataDirPath := path.Join(MetadataDir, imdbId)

	if _, err := os.Stat(MetadataDirPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(MetadataDirPath, 0755)
		if err != nil {
			fmt.Printf("Error: failed to create %s\n", MetadataDirPath)
			return
		}

		findRes := tmdbFind(imdbId)
		if findRes == nil {
			return
		}

		switch item.Type {
		case "movie":
			fetchMetadataMovie(findRes, imdbId, item)
		case "tvshow":
			fetchMetadataTvshow(findRes, imdbId, item)
		default:
			return
		}
	} else {
		metadataJsonPath := path.Join(MetadataDirPath, "imdb.json")
		imdbJsonFile, err := os.Open(metadataJsonPath)
		if err != nil {
			fmt.Printf("Error: failed to open %s\n", metadataJsonPath)
			return
		}
		defer imdbJsonFile.Close()

		switch item.Type {
		case "movie":
			err = json.NewDecoder(imdbJsonFile).Decode(&item.MovieMetadata)
			if err != nil {
				fmt.Printf("Error: failed to decode %s\n", metadataJsonPath)
				return
			}
		case "tvshow":
			err = json.NewDecoder(imdbJsonFile).Decode(&item.TvshowMetadata)
			if err != nil {
				fmt.Printf("Error: failed to decode %s\n", metadataJsonPath)
				return
			}
		default:
			return
		}
	}

	timeLayout := "2006-01-02"
	var err error

	switch item.Type {
	case "movie":
		item.ReleasedTime, err = time.Parse(timeLayout, item.MovieMetadata.ReleaseDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		fmt.Printf("Registering: %s\n", item.MovieMetadata.Title)
	case "tvshow":
		item.ReleasedTime, err = time.Parse(timeLayout, item.TvshowMetadata.FirstAirDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		fmt.Printf("Registering: %s\n", item.TvshowMetadata.Name)
	default:
		return
	}
}

func fetchMetadataMovie(findRes *tmdbFindResponse, imdbId string, item *TreeNode) {
	if len(findRes.MovieResults) > 1 {
		fmt.Printf("Error: tmdb api lookup for %s yielded multiple results, skipping", item.Path)
		return
	}

	if len(findRes.MovieResults) == 0 && len(findRes.TvResults) != 0 {
		rootLibraryName := AllLibrariesMap[item.RootUuid].DisplayName
		fmt.Printf("Error: %s in %s has likely been mislablled as a movie when it is actually a tvshow, skipping\n", item.Path, rootLibraryName)
		return
	}

	if len(findRes.MovieResults) == 0 {
		fmt.Printf("Error: tmdb api lookup for %s yielded no results, skipping", item.Path)
		return
	}

	tmdbId := findRes.MovieResults[0].ID
	movieInfo := tmdbFetchMovie(tmdbId)
	if movieInfo == nil {
		return
	}

	item.MovieMetadata = movieInfo

	posterPath := tmdbGetBestImage(movieInfo.Images.Posters)
	logoPath := tmdbGetBestImage(movieInfo.Images.Logos)
	backdropPath := tmdbGetBestImage(movieInfo.Images.Backdrops)

	MetadataDirPath := path.Join(MetadataDir, imdbId)
	posterImagePath := path.Join(MetadataDirPath, "poster.jpg")
	logoImagePath := path.Join(MetadataDirPath, "logo.jpg")
	backdropImagePath := path.Join(MetadataDirPath, "backdrop.jpg")

	err := tmdbFetchImage(posterPath, posterImagePath)
	if err != nil {
		fmt.Printf("Warning: failed to fetch poster image for %s\n", movieInfo.Title)
	}

	err = tmdbFetchImage(logoPath, logoImagePath)
	if err != nil {
		fmt.Printf("Warning: failed to fetch logo image for %s\n", movieInfo.Title)
	}

	err = tmdbFetchImage(backdropPath, backdropImagePath)
	if err != nil {
		fmt.Printf("Warning: failed to fetch backdrop image for %s\n", movieInfo.Title)
	}

	metadataJsonPath := path.Join(MetadataDirPath, "imdb.json")
	metadataJsonFile, err := os.Create(metadataJsonPath)
	if err != nil {
		fmt.Printf("Error: failed to create %s\n", metadataJsonPath)
		return
	}
	defer metadataJsonFile.Close()

	imdbJsonEncoder := json.NewEncoder(metadataJsonFile)
	imdbJsonEncoder.SetIndent("", "    ")
	if err = imdbJsonEncoder.Encode(movieInfo); err != nil {
		fmt.Printf("Error: failed to encode %s\n", metadataJsonPath)
		return
	}
}

func fetchMetadataTvshow(findRes *tmdbFindResponse, imdbId string, item *TreeNode) {
	if len(findRes.TvResults) > 1 {
		fmt.Printf("Error: tmdb api lookup for %s yielded multiple results, skipping", item.Path)
		return
	}

	if len(findRes.TvResults) == 0 && len(findRes.MovieResults) != 0 {
		rootLibraryName := AllLibrariesMap[item.RootUuid].DisplayName
		fmt.Printf("Error: %s in %s has likely been mislablled as a tvshow when it is actually a movie, skipping\n", item.Path, rootLibraryName)
		return
	}

	if len(findRes.TvResults) == 0 {
		fmt.Printf("Error: tmdb api lookup for %s yielded no results, skipping", item.Path)
		return
	}

	tmdbId := findRes.TvResults[0].ID
	tvshowInfo := tmdbFetchTvshow(tmdbId)
	if tvshowInfo == nil {
		return
	}

	item.TvshowMetadata = tvshowInfo

	posterPath := tmdbGetBestImage(tvshowInfo.Images.Posters)
	logoPath := tmdbGetBestImage(tvshowInfo.Images.Logos)
	backdropPath := tmdbGetBestImage(tvshowInfo.Images.Backdrops)

	MetadataDirPath := path.Join(MetadataDir, imdbId)
	posterImagePath := path.Join(MetadataDirPath, "poster.jpg")
	logoImagePath := path.Join(MetadataDirPath, "logo.jpg")
	backdropImagePath := path.Join(MetadataDirPath, "backdrop.jpg")

	err := tmdbFetchImage(posterPath, posterImagePath)
	if err != nil {
		fmt.Printf("Warning: failed to fetch poster image for %s\n", tvshowInfo.Name)
	}

	err = tmdbFetchImage(logoPath, logoImagePath)
	if err != nil {
		fmt.Printf("Warning: failed to fetch logo image for %s\n", tvshowInfo.Name)
	}

	err = tmdbFetchImage(backdropPath, backdropImagePath)
	if err != nil {
		fmt.Printf("Warning: failed to fetch backdrop image for %s\n", tvshowInfo.Name)
	}

	metadataJsonPath := path.Join(MetadataDirPath, "imdb.json")
	metadataJsonFile, err := os.Create(metadataJsonPath)
	if err != nil {
		fmt.Printf("Error: failed to create %s\n", metadataJsonPath)
		return
	}
	defer metadataJsonFile.Close()

	imdbJsonEncoder := json.NewEncoder(metadataJsonFile)
	imdbJsonEncoder.SetIndent("", "    ")
	if err = imdbJsonEncoder.Encode(tvshowInfo); err != nil {
		fmt.Printf("Error: failed to encode %s\n", metadataJsonPath)
		return
	}
}
