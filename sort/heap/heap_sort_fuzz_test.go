package heap

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzHeapSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	arr := heapSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(arr) {
		t.Errorf("Heap Sort failed to sort the array: %v", arr)
	}
}

func fuzzHeapSortOptimal(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	arr := heapSortOptimal(input)

	// Verify that the array is sorted
	if !sort.IntSorted(arr) {
		t.Errorf("Heap Sort failed to sort the array: %v", arr)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("counting_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzHeapSortOptimal(t, s, input)
		fuzzHeapSort(t, s, input)
	})
}
