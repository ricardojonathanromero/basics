package insertion

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		optimalWay bool
		expect     []int
		want       sort.ArrIntegerSorted
	}{
		{
			name:  "simple_set_insertion_sort",
			input: []int{43, 97, 46, 23, 1, 45},
			want:  sort.IntSorted,
		},
		{
			name:  "simple_set_insertion_sort_50",
			input: sort.GenerateRandomIntSlice(50),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_set_insertion_sort_100",
			input: sort.GenerateRandomIntSlice(100),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_set_insertion_sort_1000",
			input: sort.GenerateRandomIntSlice(1000),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_set_insertion_sort_10000",
			input: sort.GenerateRandomIntSlice(10000),
			want:  sort.IntSorted,
		},
		{
			name:       "simple_set_insertion_sort_optimized",
			input:      []int{43, 97, 46, 23, 1, 45},
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "simple_set_insertion_sort_50_optimized",
			input:      sort.GenerateRandomIntSlice(50),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "simple_set_insertion_sort_100_optimized",
			input:      sort.GenerateRandomIntSlice(100),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "simple_set_insertion_sort_1000_optimized",
			input:      sort.GenerateRandomIntSlice(1000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "simple_set_insertion_sort_10000_optimized",
			input:      sort.GenerateRandomIntSlice(10000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.optimalWay {
				insertionSortOptimized(test.input)
			} else {
				insertionSort(test.input)
			}

			if test.expect != nil && !assert.Equal(t, test.expect, test.input) {
				t.Errorf("the input is not sorted => %v", test.input)
				t.FailNow()
			}
			if test.want != nil && !test.want(test.input) {
				t.Errorf("the input is not sorted => %v", test.input)
				t.FailNow()
			}
		})
	}
}
