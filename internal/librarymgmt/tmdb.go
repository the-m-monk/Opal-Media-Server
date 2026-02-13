package librarymgmt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type tmdbFindResponseMovieResult struct {
	//	Adult            bool    `json:"adult"`
	//	BackdropPath     string  `json:"backdrop_path"`
	ID    int    `json:"id"`
	Title string `json:"title"`
	// OriginalTitle    string  `json:"original_title"`
	// Overview         string  `json:"overview"`
	// PosterPath       string  `json:"poster_path"`
	// MediaType        string  `json:"media_type"`
	// OriginalLanguage string  `json:"original_language"`
	// GenreIds         []int   `json:"genre_ids"`
	// Popularity       float64 `json:"popularity"`
	// ReleaseDate      string  `json:"release_date"`
	// Video            bool    `json:"video"`
	// VoteAverage      float64 `json:"vote_average"`
	// VoteCount        int     `json:"vote_count"`
}

type tmdbFindTvResult struct {
	//Adult        bool   `json:"adult"`
	//BackdropPath string `json:"backdrop_path"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	//OriginalName     string   `json:"original_name"`
	//Overview         string   `json:"overview"`
	//PosterPath       string   `json:"poster_path"`
	//MediaType        string   `json:"media_type"`
	//OriginalLanguage string   `json:"original_language"`
	//GenreIds         []int    `json:"genre_ids"`
	//Popularity       float64  `json:"popularity"`
	//FirstAirDate     string   `json:"first_air_date"`
	//VoteAverage      float64  `json:"vote_average"`
	//VoteCount        int      `json:"vote_count"`
	//OriginCountry    []string `json:"origin_country"`
}

// https://api.themoviedb.org/3/find/{imdb id}?external_source=imdb_id
type tmdbFindResponse struct {
	MovieResults []tmdbFindResponseMovieResult `json:"movie_results"`
	//PersonResults    []any                         `json:"person_results"`
	TvResults []tmdbFindTvResult `json:"tv_results"`
	//TvEpisodeResults []any `json:"tv_episode_results"`
	//TvSeasonResults  []any                         `json:"tv_season_results"`
}

func tmdbFind(imdbId string) *tmdbFindResponse {
	client := &http.Client{Timeout: 10 * time.Second}
	apiPath := fmt.Sprintf("https://api.themoviedb.org/3/find/%s?external_source=imdb_id", imdbId)

	for attempts := 0; attempts < 3; attempts++ {
		req, err := http.NewRequest("GET", apiPath, nil)
		if err != nil {
			return nil
		}
		req.Header.Add("Authorization", "Bearer "+tmdbApiKey)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error: tmdb api returned http status code %d\n", resp.StatusCode)
			return nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			fmt.Printf("Non-Fatal Error: rate limited by tmdb when attempting to fetch resources for %s\n", imdbId)
			//TODO: actually check Retry-After
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: status code %d for %s\n", resp.StatusCode, imdbId)
			resp.Body.Close()
			return nil
		}

		defer resp.Body.Close()
		findRes := &tmdbFindResponse{}
		if err := json.NewDecoder(resp.Body).Decode(findRes); err != nil {
			return nil
		}
		return findRes
	}

	return nil
}

type tmdbImage struct {
	AspectRatio float64 `json:"aspect_ratio"`
	Height      int     `json:"height"`
	Iso31661    string  `json:"iso_3166_1"`
	Iso6391     string  `json:"iso_639_1"`
	FilePath    string  `json:"file_path"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

type tmdbMovie struct {
	Adult               bool        `json:"adult"`
	BackdropPath        string      `json:"backdrop_path"`
	BelongsToCollection interface{} `json:"belongs_to_collection"`
	Budget              int         `json:"budget"`
	Genres              []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string   `json:"homepage"`
	ID                  int      `json:"id"`
	ImdbID              string   `json:"imdb_id"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalTitle       string   `json:"original_title"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int    `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Images      struct {
		Backdrops []tmdbImage `json:"backdrops"`
		Logos     []tmdbImage `json:"logos"`
		Posters   []tmdbImage `json:"posters"`
	} `json:"images"`
}

func tmdbFetchMovie(tmdbId int) *tmdbMovie {
	client := &http.Client{Timeout: 10 * time.Second}
	apiPath := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?append_to_response=images", tmdbId)

	for attempts := 0; attempts < 3; attempts++ {
		req, err := http.NewRequest("GET", apiPath, nil)
		if err != nil {
			return nil
		}
		req.Header.Add("Authorization", "Bearer "+tmdbApiKey)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error: tmdb api returned http status code %d\n", resp.StatusCode)
			return nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			fmt.Printf("Non-Fatal Error: rate limited by tmdb when attempting to fetch resources for tmdb-%d\n", tmdbId)
			//TODO: actually check Retry-After header
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: tmdb api returned http status code %d\n", resp.StatusCode)
			resp.Body.Close()
			return nil
		}

		defer resp.Body.Close()
		movieRes := &tmdbMovie{}
		if err := json.NewDecoder(resp.Body).Decode(movieRes); err != nil {
			return nil
		}
		return movieRes
	}

	return nil
}

type tmdbTvshow struct {
	Adult          bool          `json:"adult"`
	BackdropPath   string        `json:"backdrop_path"`
	CreatedBy      []interface{} `json:"created_by"`
	EpisodeRunTime []int         `json:"episode_run_time"`
	FirstAirDate   string        `json:"first_air_date"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int      `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		EpisodeType    string  `json:"episode_type"`
		ProductionCode string  `json:"production_code"`
		Runtime        int     `json:"runtime"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int     `json:"show_id"`
		StillPath      string  `json:"still_path"`
	} `json:"last_episode_to_air"`
	Name             string      `json:"name"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air"`
	Networks         []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	Seasons []struct {
		AirDate      string  `json:"air_date"`
		EpisodeCount int     `json:"episode_count"`
		ID           int     `json:"id"`
		Name         string  `json:"name"`
		Overview     string  `json:"overview"`
		PosterPath   string  `json:"poster_path"`
		SeasonNumber int     `json:"season_number"`
		VoteAverage  float64 `json:"vote_average"`
	} `json:"seasons"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Type        string  `json:"type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Images      struct {
		Backdrops []tmdbImage `json:"backdrops"`
		Logos     []tmdbImage `json:"logos"`
		Posters   []tmdbImage `json:"posters"`
	} `json:"images"`
}

func tmdbFetchTvshow(tmdbId int) *tmdbTvshow {
	client := &http.Client{Timeout: 10 * time.Second}
	apiPath := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d?append_to_response=images", tmdbId)

	for attempts := 0; attempts < 3; attempts++ {
		req, err := http.NewRequest("GET", apiPath, nil)
		if err != nil {
			return nil
		}
		req.Header.Add("Authorization", "Bearer "+tmdbApiKey)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error: tmdb api returned http status code %d\n", resp.StatusCode)
			return nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			fmt.Printf("Non-Fatal Error: rate limited by tmdb when attempting to fetch resources for tmdb-%d\n", tmdbId)
			//TODO: actually check Retry-After header
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: tmdb api returned http status code %d\n", resp.StatusCode)
			resp.Body.Close()
			return nil
		}

		defer resp.Body.Close()
		tvshowRes := &tmdbTvshow{}
		if err := json.NewDecoder(resp.Body).Decode(tvshowRes); err != nil {
			return nil
		}
		return tvshowRes
	}

	return nil
}

func tmdbGetBestImage(images []tmdbImage) string {
	//TODO: this logic can lead to some mediocre image selections, optimise in the future
	if len(images) == 0 {
		return ""
	}

	var filtered []tmdbImage
	for _, img := range images {
		//TODO: add supported for other language preferences
		lang := img.Iso6391
		if lang == "en" || lang == "" {
			filtered = append(filtered, img)
		}
	}

	if len(filtered) == 0 {
		return ""
	}

	sort.Slice(filtered, func(i, j int) bool {
		langI := filtered[i].Iso6391
		langJ := filtered[j].Iso6391

		if langI != langJ {
			if langI == "en" {
				return true
			}
			if langJ == "en" {
				return false
			}
		}

		scoreI := filtered[i].VoteAverage * float64(filtered[i].VoteCount)
		scoreJ := filtered[j].VoteAverage * float64(filtered[j].VoteCount)

		if scoreI != scoreJ {
			return scoreI > scoreJ
		}

		return filtered[i].VoteCount > filtered[j].VoteCount
	})

	return filtered[0].FilePath
}

func tmdbFetchImage(tmdbPath string, filePath string) error {
	apiPath := fmt.Sprintf("https://image.tmdb.org/t/p/original%s", tmdbPath)

	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("could not create directory: %w", err)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer out.Close()

	resp, err := http.Get(apiPath)
	if err != nil {
		return fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}
