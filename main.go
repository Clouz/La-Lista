package main

import (
	"La-Lista/movieFile"
	"La-Lista/themoviedb"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func init() {
	app := cli.NewApp()
	app.Name = "La Lista"
	app.Usage = "Try to search in themoviedb.org the selected files"
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			ScanDir(c.Args().First())
		} else {
			ScanDir("movieFile")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}

//ScanDir search in the selected directory using TMDB
func ScanDir(dir string) {

	m := movieFile.GetFiles(dir)

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
