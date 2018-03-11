package themoviedb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

//API info: https://developers.themoviedb.org/3/search/search-movies
var site = "https://api.themoviedb.org/3/search/movie?"
var apiKey = "api_key=da6f7ad5c8626cea07130cdf023dab07"

// SearchMovie Search for movies.
type SearchMovie struct {
	Language           string //Pass a ISO 639-1 value to display translated data
	Query              string //Pass a text query to search
	Page               string //Specify which page to query
	IncludeAdult       bool   //Choose whether to include adult content in the results
	Region             string //Specify a ISO 3166-1 code to filter release dates
	Year               int
	PrimaryReleaseYear int
}

// Schema Responses
type Schema struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}

// Result Movie List Result Object
type Result struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	GenreIds         string  `json:"genre_ids"`
	ID               int     `json:"Id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float32 `json:"popularity"`
	VoteCount        int     `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
}

// Compose return the URL for the quey
func (m SearchMovie) Compose() string {
	result := site + apiKey
	if m.Language != "" {
		result = result + "&language=" + m.Language
	}
	if m.Query != "" {
		result = result + "&query=" + url.QueryEscape(m.Query)
	}
	if m.Page != "" {
		result = result + "&page=" + m.Page
	}
	if m.IncludeAdult != false {
		result = result + "&include_adult=" + "true"
	}
	if m.Region != "" {
		result = result + "&region=" + m.Region
	}
	if m.Year != 0 {
		result = result + "&year=" + strconv.Itoa(m.Year)
	}
	if m.PrimaryReleaseYear != 0 {
		result = result + "&primary_release_year=" + strconv.Itoa(m.PrimaryReleaseYear)
	}

	return result
}

// Search do a movie search query to TheMovieDb
func (m SearchMovie) Search() (Schema, error) {

	var schema Schema

	url := m.Compose()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return schema, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&schema)
	return schema, nil
}
