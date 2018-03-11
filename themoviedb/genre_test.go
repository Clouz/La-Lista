package themoviedb_test

import (
	"La-Lista/themoviedb"
	"testing"
)

func TestGenre(t *testing.T) {
	x := themoviedb.Genre{Language: "IT-it"}
	s, err := x.Get()
	if err != nil {
		t.Errorf("Error in the Search %d", err)
	}

	t.Logf("%s", s.Genres[0].Name)
}
