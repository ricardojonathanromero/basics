package quick

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzQuickSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	quickSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Quick Sort failed to sort the array: %v", input)
	}
}

func fuzzQuickSortOptimal(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	quickSortOptimized(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Quick Sort failed to sort the array: %v", input)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("quick_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzQuickSortOptimal(t, s, input)
		fuzzQuickSort(t, s, input)
	})
}
