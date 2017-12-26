package binary

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

type Int int

func (i Int) Less(item interface{}) bool {
	return i < item.(Int)
}

func (i Int) More(item interface{}) bool {
	return i > item.(Int)
}

var collection = []utils.Comparabler{Int(2), Int(5), Int(7), Int(34), Int(111), Int(333)}

func TestSearch(t *testing.T) {
	if Search(collection, Int(1)) != 0 {
		t.Error("search error")
	}
	if Search(collection, Int(2)) != 0 {
		t.Error("search error")
	}

	if Search(collection, Int(5)) != 1 {
		t.Error("search error")
	}

	if Search(collection, Int(333)) != 5 {
		t.Error("search error")
	}

	if Search(collection, Int(12344)) != 6 {
		t.Error("search error")
	}
}
