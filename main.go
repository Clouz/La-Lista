package main

import (
	"La-Lista/movieFile"
	"La-Lista/themoviedb"
	"flag"
	"fmt"
)

func main() {

	dir := flag.String("scan", "./movieFile/", "Scan directory to search movie file")
	flag.Parse()

	m := movieFile.GetFile(*dir)

	for _, mov := range m {
		s := themoviedb.SearchMovie{
			Query:    mov.Name,
			Year:     mov.Year,
			Language: "IT-it",
		}
		sch, err := s.Search()

		if err != nil {
			fmt.Printf("[ FAIL ] %v\n", mov.Name)
			continue
		}

		fmt.Printf("[  OK  ] %v\n", sch.Results[0].Title)
	}
}
