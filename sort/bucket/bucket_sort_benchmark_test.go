package bucket

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func BenchmarkBucketSort(b *testing.B) {
	tests := []struct {
		name       string
		input      []float64
		optimalWay bool
	}{
		{
			name:  "complex_unsorted_slice_normal_sol",
			input: sort.GenerateRandomFloatSlice(20),
		},
		{
			name:       "complex_unsorted_slice_optimal_sol",
			input:      sort.GenerateRandomFloatSlice(20),
			optimalWay: true,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if test.optimalWay {
					bucketSortOptimalSolution(test.input)
				} else {
					bucketSort(test.input)
				}

				if !sort.FloatSorted(test.input) {
					b.Error("!result is not sorted")
					b.FailNow()
				}
			}
		})
	}
}
