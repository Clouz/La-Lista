package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func arg() {
	app := cli.NewApp()
	app.Name = "La Lista"
	app.Usage = "Try to search in themoviedb.org the selected files"
	app.Action = func(c *cli.Context) error {
		ScanDir(nil)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
