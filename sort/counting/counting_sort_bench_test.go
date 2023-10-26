package counting

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func BenchmarkCountingSort(b *testing.B) {
	tests := []struct {
		name       string
		input      []int
		optimalWay bool
	}{
		{
			name:  "complex_unsorted_slice_normal_sol",
			input: sort.GenerateRandomIntSlice(50),
		},
		{
			name:       "complex_unsorted_slice_optimal_sol",
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
					countingOptimizedSort(test.input)
				} else {
					countingSort(test.input)
				}

				if !sort.IntSorted(test.input) {
					b.Error("!result is not sorted")
					b.FailNow()
				}
			}
		})
	}
}
