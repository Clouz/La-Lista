package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
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

			r, _ := regexp.Compile(` \([0-9]{4}\)`)
			i := r.FindStringIndex(movie)

			if i != nil {
				//TODO: da dividere l'anno dal nome film
				fmt.Println(movie)
			}
		}
	}
}
