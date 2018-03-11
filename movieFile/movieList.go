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
	".testmovie": true, //used for testing
	".mkv":       true,
	".avi":       true,
	".mp4":       true,
	".mov":       true,
	".mpeg":      true,
	".divx":      true,
	".wmv":       true,
	".xvid":      true,
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
	cap := 0

	for i, file := range files {

		if file.IsDir() || !validExt[strings.ToLower(path.Ext(file.Name()))] {
			continue
		}

		cap++

		movies[i].File = file
		movies[i].Ext = path.Ext(file.Name())
		movies[i].Name = strings.TrimSuffix(file.Name(), movies[i].Ext)

		if strings.Contains(movies[i].Name, ".part") {
			n := strings.Split(movies[i].Name, ".part")
			switch n[1] {
			case "1":
				movies[i].Name = n[0]
			default:
				cap-- //TODO: remove other part
				continue
			}

		}

		regex, _ := regexp.Compile(` \(([0-9]{4})\)`)

		if regex.MatchString(movies[i].Name) == true {
			movies[i].Year, _ = strconv.Atoi(regex.FindStringSubmatch(movies[i].Name)[1])
			movies[i].Name = strings.Replace(movies[i].Name, regex.FindString(movies[i].Name), "", -1)
		}

		fmt.Printf("[%v]\t%v\t[%v][%v]\n", i, movies[i].Name, movies[i].Year, movies[i].Ext)
	}
	//
	//fmt.Printf("%v/%v\n", cap, len(movies))
	return movies[:cap]
}
