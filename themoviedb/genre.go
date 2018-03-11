package themoviedb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Genre Get the list of official genres for movies
type Genre struct {
	Language string
}

// SchemaGenre Responses
type SchemaGenre struct {
	Genres []GenreResponse `json:"genres"`
}

//GenreResponse Responses
type GenreResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Compose return the URL for the quey
func (m Genre) Compose() string {
	//API info: https://developers.themoviedb.org/3/genres
	var site = "https://api.themoviedb.org/3/genre/movie/list?"

	result := site + APIKey
	if m.Language != "" {
		result = result + "&language=" + m.Language
	}

	return result
}

//Get the list of official genres for movies
func (m Genre) Get() (SchemaGenre, error) {

	var schema SchemaGenre

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
