package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//API info: https://developers.themoviedb.org/3/search/search-movies
var site = "https://api.themoviedb.org/3/search/movie?"
var apiKey = "api_key=da6f7ad5c8626cea07130cdf023dab07"

type searchMovie struct {
	language           string //Pass a ISO 639-1 value to display translated data
	query              string //Pass a text query to search
	page               string //Specify which page to query
	includeAdult       bool   //Choose whether to inlcude adult content in the results
	region             string //Specify a ISO 3166-1 code to filter release dates
	year               int
	primaryReleaseYear int
}

type schema struct {
	page          int
	results       []result
	total_results int
	total_pages   int
}

type result struct {
	poster_path       string
	adult             bool
	overview          string
	release_date      string
	genre_ids         string
	id                int
	original_title    string
	original_language string
	title             string
	backdrop_path     string
	popularity        float32
	vote_count        int
	video             bool
	vote_average      float32
}

func (m searchMovie) compose() string {
	result := site + apiKey
	if m.language != "" {
		result = result + "&language=" + m.language
	}
	if m.query != "" {
		result = result + "&query=" + url.QueryEscape(m.query)
	}
	if m.page != "" {
		result = result + "&page=" + m.page
	}
	if m.includeAdult != false {
		result = result + "&include_adult=" + "true"
	}
	if m.region != "" {
		result = result + "&region=" + m.region
	}
	if m.year != 0 {
		result = result + "&year=" + strconv.Itoa(m.year)
	}
	if m.primaryReleaseYear != 0 {
		result = result + "&primary_release_year=" + strconv.Itoa(m.primaryReleaseYear)
	}

	return result
}

func main() {

	ciao := searchMovie{year: 2018, query: "Black Panther"}
	url := ciao.compose()
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("NewRequest: ", err)
		return
	}

	str, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(str))

	var x schema

	json.NewDecoder(resp.Body).Decode(&x)

	fmt.Print(x)

}
