package main

import (
	"La-Lista/themoviedb"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// ciao := themoviedb.SearchMovie{Year: 2018, Query: "Black Panther", Language: "it-IT"}
	// x, _ := ciao.Search()

	// fmt.Print(x.Results[0].Overview)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {

			file := file.Name()
			ext := path.Ext(file)
			movie := strings.TrimSuffix(file, ext)
			year := ""

			r, _ := regexp.Compile(` \(([0-9]{4})\)`)
			i := r.FindStringIndex(movie)

			if i != nil {
				year = r.FindStringSubmatch(movie)[1]
				movie = strings.Replace(movie, r.FindString(movie), "", -1)
			}

			y, _ := strconv.Atoi(year)
			tmdb := themoviedb.SearchMovie{Query: movie, Year: y}
			x, _ := tmdb.Search()
			fmt.Println(x.Results[0].Overview)
		}
	}
}
