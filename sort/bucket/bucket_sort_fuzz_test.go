package bucket

import (
	"github.com/google/gofuzz"
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzBucketSort(t *testing.T, title string, input []float64) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	arr := bucketSort(input)

	// Verify that the array is sorted
	if !sort.FloatSorted(arr) {
		t.Errorf("Bubble Sort failed to sort the array: %v", arr)
	}
}

func fuzzBucketSortOptimal(t *testing.T, title string, input []float64) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	arr := bucketSortOptimalSolution(input)

	// Verify that the array is sorted
	if !sort.FloatSorted(arr) {
		t.Errorf("Bubble Sort failed to sort the array: %v", arr)
	}
}

func FuzzAll(f *testing.F) {
	var input []float64
	ff := fuzz.New()

	f.Add("bucket_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		for i := 0; i < 5; i++ {
			ff.NumElements(20, 1000).Fuzz(&input)
			fuzzBucketSort(t, s, input)
			fuzzBucketSortOptimal(t, s, input)
		}
	})
}
