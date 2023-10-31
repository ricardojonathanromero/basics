package quick

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		optimizedWay bool
		want         sort.ArrIntegerSorted
	}{
		{
			name:  "simple_test",
			input: []int{2, 6, 5, 3, 8, 7, 1, 0},
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
			name:  "simple_test_100000",
			input: sort.GenerateRandomIntSlice(100000),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_1000000",
			input: sort.GenerateRandomIntSlice(1000000),
			want:  sort.IntSorted,
		},
		{
			name:         "optimized_test",
			input:        []int{2, 6, 5, 3, 8, 7, 1, 0},
			optimizedWay: true,
			want:         sort.IntSorted,
		},
		{
			name:         "optimized_test_10",
			input:        sort.GenerateRandomIntSlice(10),
			optimizedWay: true,
			want:         sort.IntSorted,
		},
		{
			name:         "optimized_test_100",
			input:        sort.GenerateRandomIntSlice(100),
			optimizedWay: true,
			want:         sort.IntSorted,
		},
		{
			name:         "optimized_test_1000",
			input:        sort.GenerateRandomIntSlice(1000),
			optimizedWay: true,
			want:         sort.IntSorted,
		},
		{
			name:         "optimized_test_10000",
			input:        sort.GenerateRandomIntSlice(10000),
			optimizedWay: true,
			want:         sort.IntSorted,
		},
		{
			name:         "optimized_test_100000",
			input:        sort.GenerateRandomIntSlice(100000),
			optimizedWay: true,
			want:         sort.IntSorted,
		},
		{
			name:         "optimized_test_1000000",
			input:        sort.GenerateRandomIntSlice(1000000),
			optimizedWay: true,
			want:         sort.IntSorted,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.optimizedWay {
				quickSortOptimized(test.input)
			} else {
				quickSort(test.input)
			}

			if test.want != nil && !test.want(test.input) {
				t.FailNow()
			}
		})
	}
}
