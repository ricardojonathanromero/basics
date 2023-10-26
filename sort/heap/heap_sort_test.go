package heap

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
			name:  "simple_heap_sort",
			input: []int{23, 13, 56, 8, 3, 2, 9, 89},
			want:  sort.IntSorted,
		},
		{
			name:  "complex_heap_sort_100",
			input: sort.GenerateRandomIntSlice(100),
			want:  sort.IntSorted,
		},
		{
			name:  "complex_heap_sort_500",
			input: sort.GenerateRandomIntSlice(500),
			want:  sort.IntSorted,
		},
		{
			name:  "complex_heap_sort_1000",
			input: sort.GenerateRandomIntSlice(1000),
			want:  sort.IntSorted,
		},

		{
			name:       "optimal_heap_sort",
			input:      []int{23, 13, 56, 8, 3, 2, 9, 89},
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "complex_optimal_heap_sort_100",
			input:      sort.GenerateRandomIntSlice(100),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "complex_optimal_heap_sort_500",
			input:      sort.GenerateRandomIntSlice(500),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "complex_optimal_heap_sort_1000",
			input:      sort.GenerateRandomIntSlice(1000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.optimalWay {
				heapSortOptimal(test.input)
			} else {
				heapSort(test.input)
			}

			if test.want != nil && !test.want(test.input) {
				t.Errorf("array is not sorted! \n%v\n", test.input)
				t.FailNow()
			}
		})
	}
}
