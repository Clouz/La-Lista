package main

import (
	"La-Lista/movieFile"
	"La-Lista/themoviedb"
	"fmt"
	"time"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func main() {

	arg()

	// dir := flag.String("scan", "./movieFile/", "Scan directory to search movie file")
	// flag.Parse()

}

//ScanDir search in the selected directory using TMDB
func ScanDir(dir string) {

	m := movieFile.GetFile(dir)

	for _, mov := range m {
		time.Sleep(time.Duration(200) * time.Millisecond) //TODO: try to reduce delay
		s := themoviedb.SearchMovie{
			Query:    mov.Name,
			Year:     mov.Year,
			Language: "IT-it",
		}
		sch, err := s.Search()

		if err != nil {
			fmt.Printf("[ %v ] %v\n", red("FAIL"), mov.Name)
			continue
		}

		fmt.Printf("[  %v  ] %v\n", green("OK"), sch.Results[0].Title)
	}

}
