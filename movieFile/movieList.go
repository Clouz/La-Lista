package movieFile

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// validExt represent hello
var validExt = map[string]bool{
	".testMovie": true,
	// TODO: Add all valid file extension
}

// Movie contain all the info given in the filename
type Movie struct {
	File os.FileInfo //Movie location
	Name string      //Movie name
	Ext  string      //Movie file extension
	Year int         //Movie year if exist otherwise is 0
}

//GetFile return all the movies found in the directory
func GetFile(dir string) []Movie {

	//Read all the file in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	movies := make([]Movie, len(files))

	for i, file := range files {

		if file.IsDir() || !validExt[path.Ext(file.Name())] {
			continue
		}

		movies[i].File = file
		movies[i].Ext = path.Ext(file.Name())
		movies[i].Name = strings.TrimSuffix(file.Name(), movies[i].Ext)

		regex, _ := regexp.Compile(` \(([0-9]{4})\)`)

		if regex.MatchString(movies[i].Name) == true {
			movies[i].Year, _ = strconv.Atoi(regex.FindStringSubmatch(movies[i].Name)[1])
			movies[i].Name = strings.Replace(movies[i].Name, regex.FindString(movies[i].Name), "", -1)
		}

		fmt.Printf("[%v]\t%v\t[%v][%v]\n", i, movies[i].Name, movies[i].Year, movies[i].Ext)
	}

	return movies
}
