package movieFile

import (
	"fmt"
	"regexp"
)

func RawFilename() string {
	var newName string

	content := "Justice.League.2017.DVD.Rip.AVC.x264.AC3.ITA.AAC.ENG.Subs.^MF^.mkv"

	pattern := regexp.MustCompile(`^(?P<name>.+)\.(?P<year>\d{4}).+\.(?P<ext>mp4|avi|mkv)$`)

	template := "$name ($year).$ext\n"

	result := []byte{}

	submatches := pattern.FindAllStringSubmatchIndex(content, -1)

	result = pattern.ExpandString(result, template, content, submatches[0])

	fmt.Println(string(result))

	return newName
}
