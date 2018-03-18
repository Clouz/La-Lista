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

//Delay each request to the movie db server
var Delay int

//Default choice for single request
var Default bool

//Folder selected without argument
var Folder string

func start() {
	app := cli.NewApp()
	app.Name = "La Lista"
	app.Usage = "Try to search in themoviedb.org the selected files"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Claudio Mola",
			Email: "Clouz85@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "delay",
			Value:       200,
			Usage:       "Delay each request to the movie db server",
			Destination: &Delay,
		},
		cli.BoolFlag{
			Name:        "default",
			Usage:       "Default choice for single request",
			Destination: &Default,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			ScanDir(c.Args().First())
		} else {
			ScanDir("movieFile/debug/Rogue One: A Star Wars Story (2016).testMovie")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Default: %v, Delay: %v", Default, Delay)
}

func main() {
	start()
}

//ScanDir search in the selected directory using TMDB
func ScanDir(dir string) {

	m := movieFile.GetFiles(dir)

	if len(m) > 1 {
		for _, mov := range m {
			time.Sleep(time.Duration(Delay) * time.Millisecond) //TODO: try to reduce delay
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
			os.Exit(1)
		}

		fmt.Println("Select correct film:")
		for i, r := range sch.Results {
			if i > 10 {
				break
			}
			fmt.Printf("[%v]\t%v (%v)\n", i, r.Title, r.ReleaseDate)
		}

		var i int
		if Default == false {
			fmt.Print("\n> ")
			_, err = fmt.Scanf("%d", &i)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			i = 0
		}

		fn := sch.Results[i].Title + "(" + sch.Results[0].ReleaseDate + ")" + m[0].Ext
		fmt.Println(fn)
	}
}
