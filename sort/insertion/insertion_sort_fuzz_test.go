package insertion

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzInsertionSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	insertionSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Insertion Sort failed to sort the array: %v", input)
	}
}

func fuzzInsertionSortOptimal(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	insertionSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Insertion Sort failed to sort the array: %v", input)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("counting_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzInsertionSortOptimal(t, s, input)
		fuzzInsertionSort(t, s, input)
	})
}
