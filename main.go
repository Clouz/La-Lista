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

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), "testMovie") {

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
			tmdb := themoviedb.SearchMovie{Query: movie, Year: y, Language: "IT-it"}
			x, _ := tmdb.Search()
			fmt.Println(x.Results)
			fmt.Println("\n----------------------\n\n\n\n")
		}
	}
}
