package themoviedb_test

import (
	"La-Lista/themoviedb"
	"testing"
)

func TestSearch(t *testing.T) {
	x := themoviedb.SearchMovie{Query: "Deadpool", Year: 2016, IncludeAdult: true, Language: "IT-it"}
	s, err := x.Search()
	if err != nil {
		t.Errorf("Error in the Search %d", err)
	}

	t.Logf("%s", s.Results[0].Title)
}
