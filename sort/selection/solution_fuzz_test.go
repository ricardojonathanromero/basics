package selection

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzSelectionSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	selectionSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Selection Sort failed to sort the array: %v", input)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("selection_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzSelectionSort(t, s, input)
	})
}
