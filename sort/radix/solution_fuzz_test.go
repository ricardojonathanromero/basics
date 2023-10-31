package radix

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzRadixSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	radixSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Radix Sort failed to sort the array: %v", input)
	}
}

func fuzzRadixSortOptimal(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	radixSortOptimized(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Radix Sort failed to sort the array: %v", input)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("radix_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzRadixSortOptimal(t, s, input)
		fuzzRadixSort(t, s, input)
	})
}
