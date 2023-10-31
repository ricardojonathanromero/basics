package radix

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		optimalWay bool
		want       sort.ArrIntegerSorted
	}{
		{
			name:  "simple_test",
			input: []int{170, 45, 75, 90, 802, 24, 2, 66},
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_10",
			input: sort.GenerateRandomIntSlice(10),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_100",
			input: sort.GenerateRandomIntSlice(100),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_1000",
			input: sort.GenerateRandomIntSlice(1000),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_10000",
			input: sort.GenerateRandomIntSlice(10000),
			want:  sort.IntSorted,
		},
		{
			name:       "optimal_test",
			input:      []int{170, 45, 75, 90, 802, 24, 2, 66},
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "optimal_test_10",
			input:      sort.GenerateRandomIntSlice(10),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "optimal_test_100",
			input:      sort.GenerateRandomIntSlice(100),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "optimal_test_1000",
			input:      sort.GenerateRandomIntSlice(1000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "optimal_test_10000",
			input:      sort.GenerateRandomIntSlice(10000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.optimalWay {
				radixSortOptimized(test.input)
			} else {
				radixSort(test.input)
			}

			if test.want != nil && !test.want(test.input) {
				t.FailNow()
			}
		})
	}
}
