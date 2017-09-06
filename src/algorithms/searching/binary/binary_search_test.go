package binary

import "testing"

type ColInt []int

func (ci ColInt) Len() int {
	return len(ci)
}

func (ci ColInt) Less(i int, t interface{}) bool {
	return ci[i] < t.(int)
}

func (ci ColInt) More(i int, t interface{}) bool {
	return ci[i] > t.(int)
}

var collection = ColInt{2, 5, 7, 34, 111, 333}

func TestSearch(t *testing.T) {
	if Search(collection, 1) != 0 {
		t.Error("search error")
	}
	if Search(collection, 2) != 0 {
		t.Error("search error")
	}

	if Search(collection, 5) != 1 {
		t.Error("search error")
	}

	if Search(collection, 333) != 5 {
		t.Error("search error")
	}

	if Search(collection, 12344) != 6 {
		t.Error("search error")
	}
}
