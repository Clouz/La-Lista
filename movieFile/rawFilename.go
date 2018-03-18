package movieFile

import (
	"fmt"
	"regexp"
)

func RawFilename() string {
	var newName string

	//f2 := "La.Forma.dell.acqua.2017.720p.ita.dts.eng.aac.subs.fd.mkw"
	content := "Justice.League.2017.DVD.Rip.AVC.x264.AC3.ITA.AAC.ENG.Subs.^MF^.mkv"
	regexp := regexp.MustCompile(`^(?P<name>.+)\.(?P<year>\d{4}).+\.(?P<ext>mp4|avi|mkv)$`)
	template := "$name ($year).$ext"
	result := []byte{}

	for _, submatches := range regexp.FindAllStringSubmatchIndex(content, -1) {
		regexp.ExpandString(result, template, content, submatches)
	}

	fmt.Println(string(result))

	return newName
}
