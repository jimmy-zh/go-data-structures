package quick

import (
	"testing"
	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestCutOffToInsertionSort(t *testing.T) {
	if !utils.SortTest(5, Sort) {
		t.Error("Not Sorted")
	}
}
func TestSort(t *testing.T) {
	if !utils.SortTest(1000, Sort) {
		t.Error("Not Sorted")
	}
}

func TestMedian3(t *testing.T) {
	if Median3([]int{2, 1, 3}, 0, 1, 2) != 0 {
		t.Error("Not median number")
	}
	if Median3([]int{1, 2, 3}, 0, 1, 2) != 1 {
		t.Error("Not median number")
	}
	if Median3([]int{3, 1, 2}, 0, 1, 2) != 2 {
		t.Error("Not median number")
	}
}
