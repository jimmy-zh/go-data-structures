package selection

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestSort(t *testing.T) {
	if !utils.SortTest(1000, Sort) {
		t.Error("Not Sorted!")
	}
}
