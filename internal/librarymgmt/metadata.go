package librarymgmt

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
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
var ratingSystem string

var tvshowMovieDetectionRegex = regexp.MustCompile(`\[(movie|tvshow)-imdb-(tt\d+)\]`)

func initMetadata() {
	tmdbApiKey = config.FetchValue("/tmdb.cfg", "api_key", true)

	if strings.HasPrefix(tmdbApiKey, "$") {
		tmdbApiKey = os.Getenv(strings.TrimPrefix(tmdbApiKey, "$"))
	}

	MetadataDir = config.FetchValue("/server.cfg", "metadata_dir", true)
	if err := os.MkdirAll(MetadataDir, 0755); err != nil {
		slog.Error("failed to mkdir", "dir", MetadataDir, "reason", err.Error())
		os.Exit(1)
	}

	ratingSystem = config.FetchValue("/server.cfg", "rating_system", true)

	buildLibraryTree()

	for _, rootNode := range LibTree {
		slices.SortFunc(NewestMedia[rootNode.RootUuid], func(a, b MtimeSortEntry) int {
			return int(b.Mtime.Unix() - a.Mtime.Unix())
		})
	}

}

func buildLibraryTree() {
	for i, rootLib := range AllLibraries {
		slog.Info("Scanning library", "libDisplayName", rootLib.DisplayName)

		rootStat, err := os.Stat(rootLib.Path)
		if err != nil {
			slog.Error("failed to parse library", "rootLib.Path", rootLib.Path, "reason", err.Error())
			os.Exit(1)
		}

		if !rootStat.IsDir() {
			slog.Error("library path is not a directory", "offendingLibrary", rootLib.DisplayName, "offendingPath", rootLib.Path)
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

		imageCachePath := filepath.Join(CacheDir, "images")
		libNameCardPath := filepath.Join(imageCachePath, newLibNode.RootUuid+".png")

		os.MkdirAll(imageCachePath, 0755)
		err = RenderNameCard(rootLib.DisplayName, libNameCardPath)
		if err != nil {
			slog.Warn("failed to render library name card", "lib.DisplayName", rootLib.DisplayName, "reason", err.Error())
		}
	}
}

func searchDir(libRootPath string, relativePath string, libNamespace uuid.UUID) *TreeNode {
	slog.Info("Searching for media", "relPath", relativePath)

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
		slog.Error("failed to list library directory", "path", absPath, "reason", err.Error())
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
			slog.Error("failed to mkdir", "path", MetadataDirPath, "reason", err.Error())
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
			slog.Error("failed to open file", "path", metadataJsonPath, "reason", err.Error())
			return
		}
		defer imdbJsonFile.Close()

		switch item.Type {
		case "movie":
			err = json.NewDecoder(imdbJsonFile).Decode(&item.MovieMetadata)
			if err != nil {
				slog.Error("failed to decode json", "file", metadataJsonPath, "reason", err.Error())
				return
			}
		case "tvshow":
			err = json.NewDecoder(imdbJsonFile).Decode(&item.TvshowMetadata)
			if err != nil {
				slog.Error("failed to decode json", "file", metadataJsonPath, "reason", err.Error())
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
			slog.Error("failed to parse date", "offendingMedia", item.MovieMetadata.Title, "reason", err.Error())
			return
		}

		for _, r := range item.MovieMetadata.ReleaseDates.Results {
			if r.Iso31661 != ratingSystem {
				continue
			}
			item.Rating = r.ReleaseDates[0].Certification
		}

		slog.Info("Registering movie", "title", item.MovieMetadata.Title)
	case "tvshow":
		item.ReleasedTime, err = time.Parse(timeLayout, item.TvshowMetadata.FirstAirDate)
		if err != nil {
			slog.Error("failed to parse date", "offendingMedia", item.TvshowMetadata.Name, "reason", err.Error())
			return
		}

		//for _, r := range item.TvshowMetadata.ReleaseDates.Results {
		//	if r.Iso31661 != ratingSystem {
		//		continue
		//	}
		//	item.Rating = r.ReleaseDates[0].Certification
		//}

		slog.Info("Registering tvshow", "name", item.TvshowMetadata.Name)
	default:
		return
	}
}

func fetchMetadataMovie(findRes *tmdbFindResponse, imdbId string, item *TreeNode) {
	if len(findRes.MovieResults) > 1 {
		slog.Warn("tmdb api lookup yield multiple results, skipping", "offendingMedia", item.Path)
		return
	}

	if len(findRes.MovieResults) == 0 && len(findRes.TvResults) != 0 {
		rootLibraryName := AllLibrariesMap[item.RootUuid].DisplayName
		slog.Warn("item has likely been mislablled as a movie when it is actually a tvshow, skipping", "item", fmt.Sprintf("%s in %s", item.Path, rootLibraryName))
		return
	}

	if len(findRes.MovieResults) == 0 {
		slog.Warn("tmdb api lookup yield not results, skipping", "item", item.Path)
		return
	}

	tmdbId := findRes.MovieResults[0].ID
	movieInfo := tmdbFetchMovie(tmdbId)
	if movieInfo == nil {
		return
	}

	posterPath, posterPrimaryAspectRatio := tmdbGetBestImage(movieInfo.Images.Posters)
	logoPath, logoPrimaryAspectRatio := tmdbGetBestImage(movieInfo.Images.Logos)
	backdropPath, backdropPrimaryAspectRatio := tmdbGetBestImage(movieInfo.Images.Backdrops)

	movieInfo.PosterPrimaryAspectRatio = posterPrimaryAspectRatio
	movieInfo.LogoPrimaryAspectRatio = logoPrimaryAspectRatio
	movieInfo.BackdropPrimaryAspectRatio = backdropPrimaryAspectRatio

	item.MovieMetadata = movieInfo

	MetadataDirPath := path.Join(MetadataDir, imdbId)
	posterImagePath := path.Join(MetadataDirPath, "poster.png")
	logoImagePath := path.Join(MetadataDirPath, "logo.png")
	backdropImagePath := path.Join(MetadataDirPath, "backdrop.png")

	err := tmdbFetchImage(posterPath, posterImagePath)
	if err != nil {
		slog.Warn("failed to fetch poster image", "item", movieInfo.Title, "reason", err.Error())
	}

	err = tmdbFetchImage(logoPath, logoImagePath)
	if err != nil {
		slog.Warn("failed to fetch logo image", "item", movieInfo.Title, "reason", err.Error())
	}

	err = tmdbFetchImage(backdropPath, backdropImagePath)
	if err != nil {
		slog.Warn("failed to fetch backdrop image", "item", movieInfo.Title, "reason", err.Error())
	}

	metadataJsonPath := path.Join(MetadataDirPath, "imdb.json")
	metadataJsonFile, err := os.Create(metadataJsonPath)
	if err != nil {
		slog.Warn("failed to create imdb.json", "item", metadataJsonPath, "reason", err.Error())
		return
	}
	defer metadataJsonFile.Close()

	imdbJsonEncoder := json.NewEncoder(metadataJsonFile)
	imdbJsonEncoder.SetIndent("", "    ")
	if err = imdbJsonEncoder.Encode(movieInfo); err != nil {
		slog.Warn("failed to encode imdb.json", "item", movieInfo.Title, "reason", err.Error())
		return
	}
}

func fetchMetadataTvshow(findRes *tmdbFindResponse, imdbId string, item *TreeNode) {
	if len(findRes.TvResults) > 1 {
		slog.Warn("tmdb api lookup yield multiple results, skipping", "offendingMedia", item.Path)
		return
	}

	if len(findRes.TvResults) == 0 && len(findRes.MovieResults) != 0 {
		rootLibraryName := AllLibrariesMap[item.RootUuid].DisplayName
		slog.Warn("item has likely been mislablled as a tvshow when it is actually a movie, skipping", "item", fmt.Sprintf("%s in %s", item.Path, rootLibraryName))
		return
	}

	if len(findRes.TvResults) == 0 {
		slog.Warn("tmdb api lookup yield not results, skipping", "item", item.Path)
		return
	}

	tmdbId := findRes.TvResults[0].ID
	tvshowInfo := tmdbFetchTvshow(tmdbId)
	if tvshowInfo == nil {
		return
	}

	posterPath, posterPrimaryAspectRatio := tmdbGetBestImage(tvshowInfo.Images.Posters)
	logoPath, logoPrimaryAspectRatio := tmdbGetBestImage(tvshowInfo.Images.Logos)
	backdropPath, backdropPrimaryAspectRatio := tmdbGetBestImage(tvshowInfo.Images.Backdrops)

	tvshowInfo.PosterPrimaryAspectRatio = posterPrimaryAspectRatio
	tvshowInfo.LogoPrimaryAspectRatio = logoPrimaryAspectRatio
	tvshowInfo.BackdropPrimaryAspectRatio = backdropPrimaryAspectRatio

	item.TvshowMetadata = tvshowInfo

	MetadataDirPath := path.Join(MetadataDir, imdbId)
	posterImagePath := path.Join(MetadataDirPath, "poster.png")
	logoImagePath := path.Join(MetadataDirPath, "logo.png")
	backdropImagePath := path.Join(MetadataDirPath, "backdrop.png")

	err := tmdbFetchImage(posterPath, posterImagePath)
	if err != nil {
		slog.Warn("failed to fetch poster image", "item", tvshowInfo.Name, "reason", err.Error())
	}

	err = tmdbFetchImage(logoPath, logoImagePath)
	if err != nil {
		slog.Warn("failed to fetch logo image", "item", tvshowInfo.Name, "reason", err.Error())
	}

	err = tmdbFetchImage(backdropPath, backdropImagePath)
	if err != nil {
		slog.Warn("failed to fetch backdrop image", "item", tvshowInfo.Name, "reason", err.Error())
	}

	metadataJsonPath := path.Join(MetadataDirPath, "imdb.json")
	metadataJsonFile, err := os.Create(metadataJsonPath)
	if err != nil {
		slog.Warn("failed to create imdb.json", "item", metadataJsonPath, "reason", err.Error())
		return
	}
	defer metadataJsonFile.Close()

	imdbJsonEncoder := json.NewEncoder(metadataJsonFile)
	imdbJsonEncoder.SetIndent("", "    ")
	if err = imdbJsonEncoder.Encode(tvshowInfo); err != nil {
		slog.Warn("failed to encode imdb.json", "item", tvshowInfo.Name, "reason", err.Error())
		return
	}
}
