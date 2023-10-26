package counting

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzCountingSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	countingSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Counting Sort failed to sort the array: %v", input)
	}
}

func fuzzCountingSortOptimal(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	countingOptimizedSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Counting Sort failed to sort the array: %v", input)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("counting_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzCountingSortOptimal(t, s, input)
		fuzzCountingSort(t, s, input)
	})
}
