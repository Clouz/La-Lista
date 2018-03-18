package movieFile

import (
	"fmt"
	"regexp"
)

func RawFilename() string {
	var newName string

	//f2 := "La.Forma.dell.acqua.2017.720p.ita.dts.eng.aac.subs.fd.mkw"
	content := "Justice.League.2017.DVD.Rip.AVC.x264.AC3.ITA.AAC.ENG.Subs.^MF^.mkv"

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`^(?P<name>.+)\.(?P<year>\d{4}).+\.(?P<ext>mp4|avi|mkv)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := "$name ($year).$ext\n"

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.ExpandString(result, template, content, submatches)
	}
	fmt.Println(string(result))

	return newName
}
