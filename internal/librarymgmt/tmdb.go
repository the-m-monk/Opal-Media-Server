package librarymgmt

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
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
			slog.Error("tmdb api failed during find", "httpStatusCode", resp.StatusCode, "imdbID", imdbId)
			return nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			slog.Info("non-fatal error: rate limited by tmdb when attempting to fetch resources", "imdbID", imdbId)
			//TODO: actually check Retry-After
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			slog.Error("tmdb api returned non-ok status code", "statusCode", resp.StatusCode, "imdbID", imdbId)
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

func tmdbFetchMovie(tmdbId int) *tmdbMovie {
	client := &http.Client{Timeout: 10 * time.Second}
	apiPath := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?append_to_response=credits,images,videos,release_dates,keywords,recommendations,similar,translations,external_ids", tmdbId)

	for attempts := 0; attempts < 3; attempts++ {
		req, err := http.NewRequest("GET", apiPath, nil)
		if err != nil {
			return nil
		}
		req.Header.Add("Authorization", "Bearer "+tmdbApiKey)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			slog.Error("tmdb api failed during movie fetch", "httpStatusCode", resp.StatusCode, "tmdbID", tmdbId)
			return nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			slog.Info("non-fatal error: rate limited by tmdb when attempting to fetch resources", "tmdbID", tmdbId)
			//TODO: actually check Retry-After
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			slog.Error("tmdb api returned non-ok status code", "statusCode", resp.StatusCode, "tmdbID", tmdbId)
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

func tmdbFetchTvshow(tmdbId int) *tmdbTvshow {
	client := &http.Client{Timeout: 10 * time.Second}
	apiPath := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d?append_to_response=credits,images,videos,release_dates,keywords,recommendations,similar,translations,external_ids", tmdbId)

	for attempts := 0; attempts < 3; attempts++ {
		req, err := http.NewRequest("GET", apiPath, nil)
		if err != nil {
			return nil
		}
		req.Header.Add("Authorization", "Bearer "+tmdbApiKey)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			slog.Error("tmdb api failed during tvshow fetch", "httpStatusCode", resp.StatusCode, "tmdbID", tmdbId)
			return nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			slog.Info("non-fatal error: rate limited by tmdb when attempting to fetch resources", "tmdbID", tmdbId)
			//TODO: actually check Retry-After
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			slog.Error("tmdb api returned non-ok status code", "statusCode", resp.StatusCode, "tmdbID", tmdbId)
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

func tmdbGetBestImage(images []tmdbImage) (url string, aspectRatio float64) {
	//TODO: this logic can lead to some mediocre image selections, optimise in the future
	if len(images) == 0 {
		return "", 0
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
		return "", 0
	}
	/*
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
	*/
	return filtered[0].FilePath, filtered[0].AspectRatio
}

func tmdbFetchImage(tmdbPath string, filePath string) error {
	apiPath := fmt.Sprintf("https://image.tmdb.org/t/p/original%s", tmdbPath)

	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("could not create directory: %w", err)
	}

	resp, err := http.Get(apiPath)
	if err != nil {
		return fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer out.Close()

	encoder := png.Encoder{CompressionLevel: png.DefaultCompression}
	err = encoder.Encode(out, img)
	if err != nil {
		return fmt.Errorf("failed to encode png: %w", err)
	}

	return nil
}
