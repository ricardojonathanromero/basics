package bubble

import (
	fuzz "github.com/google/gofuzz"
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzBubbleSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	arr := bubbleSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(arr) {
		t.Errorf("Bubble Sort failed to sort the array: %v", arr)
	}
}

func fuzzBubbleSortOptimal(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	arr := bubbleOptimalSolution(input)

	// Verify that the array is sorted
	if !sort.IntSorted(arr) {
		t.Errorf("Bubble Sort failed to sort the array: %v", arr)
	}
}

func FuzzAll(f *testing.F) {
	var input []int
	ff := fuzz.New()

	f.Add("bubble_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		for i := 0; i < 5; i++ {
			ff.NumElements(20, 1000).Fuzz(&input)
			fuzzBubbleSort(t, s, input)
			fuzzBubbleSortOptimal(t, s, input)
		}
	})
}
