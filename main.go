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

//Language used to search film
var Language string

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
			Name:        "default, d",
			Usage:       "Default choice for single request",
			Destination: &Default,
		},
		cli.StringFlag{
			Name:        "folder,f",
			Usage:       "Folder selected withour argument",
			Value:       ".",
			Destination: &Folder,
		},
		cli.StringFlag{
			Name:        "language, l",
			Usage:       "Language used to search film",
			Value:       "it-IT",
			Destination: &Language,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			ScanDir(c.Args().First())
		} else {
			ScanDir(Folder)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	start()

	movieFile.RawFilename()
}

//ScanDir search in the selected directory using TMDB
func ScanDir(dir string) {

	m := movieFile.GetFiles(dir)

	if len(m) > 1 {
		for _, mov := range m {
			time.Sleep(time.Duration(Delay) * time.Millisecond)
			s := themoviedb.SearchMovie{
				Query:    mov.Name,
				Year:     mov.Year,
				Language: Language,
			}
			sch, err := s.Search()

			if err != nil {
				fmt.Printf("[ %v ] %v\n", red("FAIL"), mov.Name)
				continue
			}

			fmt.Printf("[  %v  ] %v\n", green("OK"), sch.Results[0].Title)
		}
	} else if len(m) == 1 {
		s := themoviedb.SearchMovie{
			Query:    m[0].Name,
			Year:     m[0].Year,
			Language: Language,
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
