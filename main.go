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

var delay int = 200

func init() {
	app := cli.NewApp()
	app.Name = "La Lista"
	app.Usage = "Try to search in themoviedb.org the selected files"
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			ScanDir(c.Args().First())
		} else {
			ScanDir("movieFile/Rogue One: A Star Wars Story (2016).testMovie")
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

	if len(m) > 1 {
		for _, mov := range m {
			time.Sleep(time.Duration(delay) * time.Millisecond) //TODO: try to reduce delay
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
	} else {
		s := themoviedb.SearchMovie{
			Query:    m[0].Name,
			Year:     m[0].Year,
			Language: "IT-it",
		}
		sch, err := s.Search()

		if err != nil {
			fmt.Printf("[ %v ] %v\n", red("FAIL"), m[0].Name)
		}

		fmt.Println("Select correct film:")
		for i, r := range sch.Results {
			if i > 10 {
				break
			}
			fmt.Printf("[%v]\t%v (%v)\n", i, r.Title, r.ReleaseDate)
		}

		fmt.Print("\n> ")
		var i int
		_, err = fmt.Scanf("%d", &i)
		if err != nil {
			log.Fatal(err)
		}

		fn := sch.Results[i].Title + "(" + sch.Results[0].ReleaseDate + ")" + m[0].Ext
		fmt.Println(fn)
	}
}
