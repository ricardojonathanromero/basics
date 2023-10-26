package heap

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func BenchmarkHeapSort(b *testing.B) {
	tests := []struct {
		name       string
		input      []int
		optimalWay bool
	}{
		{
			name:  "complex_unsorted_slice",
			input: sort.GenerateRandomIntSlice(50),
		},
		{
			name:       "complex_unsorted_slice_optimal",
			input:      sort.GenerateRandomIntSlice(50),
			optimalWay: true,
		},
	}

	// iterate tests
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			// iterate until count parameter. by default 1
			for i := 0; i < b.N; i++ {
				if test.optimalWay {
					heapSortOptimal(test.input)
				} else {
					heapSort(test.input)
				}

				if !sort.IntSorted(test.input) {
					b.Error("!result is not sorted")
					b.FailNow()
				}
			}
		})
	}
}
