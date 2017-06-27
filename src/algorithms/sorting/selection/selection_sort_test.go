package selection

import (
	"testing"
	"utils"
)

func TestSort(t *testing.T) {
	if !utils.SortTest(1000, Sort) {
		t.Error("Not Sorted!")
	}
}
