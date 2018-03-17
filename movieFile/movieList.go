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

//GetFiles return all the movies found in the directory
func GetFiles(dir string) []Movie {

	file, err := os.Stat(dir)
	if err != nil {
		log.Fatal(err)
	}

	var files []os.FileInfo

	if file.IsDir() {
		files, err = ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		files = []os.FileInfo{file}
	}

	movies := make([]Movie, len(files))
	cap := 0

	for _, file := range files {

		if file.IsDir() || !validExt[strings.ToLower(path.Ext(file.Name()))] {
			continue
		}

		movies[cap].File = file
		movies[cap].Ext = path.Ext(file.Name())
		movies[cap].Name = strings.TrimSuffix(file.Name(), movies[cap].Ext)

		if strings.Contains(movies[cap].Name, ".part") {
			n := strings.Split(movies[cap].Name, ".part")
			switch n[1] {
			case "1":
				movies[cap].Name = n[0]
			default:
				continue
			}
		}

		regex, _ := regexp.Compile(` \(([0-9]{4})\)`)

		if regex.MatchString(movies[cap].Name) == true {
			movies[cap].Year, _ = strconv.Atoi(regex.FindStringSubmatch(movies[cap].Name)[1])
			movies[cap].Name = strings.Replace(movies[cap].Name, regex.FindString(movies[cap].Name), "", -1)
		}

		fmt.Printf("[%v]\t%v\t[%v][%v]\n", cap, movies[cap].Name, movies[cap].Year, movies[cap].Ext)

		cap++
	}

	return movies[:cap]
}
