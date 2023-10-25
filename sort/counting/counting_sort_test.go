package counting

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingSort(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		optimalWay bool
		expected   []int
		want       sort.ArrIntegerSorted
	}{
		{
			name:     "one_item",
			input:    []int{90},
			expected: []int{90},
		},
		{
			name:     "random_unsorted_slice",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "random_unsorted_slice_20",
			input: sort.GenerateRandomIntSlice(20),
			want:  sort.IntSorted,
		},
		{
			name:  "random_unsorted_slice_50",
			input: sort.GenerateRandomIntSlice(50),
			want:  sort.IntSorted,
		},
		{
			name:  "random_unsorted_slice_100",
			input: sort.GenerateRandomIntSlice(100),
			want:  sort.IntSorted,
		},
		{
			name:  "random_unsorted_slice_1000",
			input: sort.GenerateRandomIntSlice(1000),
			want:  sort.IntSorted,
		},
		{
			name:  "random_unsorted_slice_10000",
			input: sort.GenerateRandomIntSlice(10000),
			want:  sort.IntSorted,
		},
		{
			name:  "random_unsorted_slice_100000",
			input: sort.GenerateRandomIntSlice(100000),
			want:  sort.IntSorted,
		},
		{
			name:  "random_unsorted_slice_1000000",
			input: sort.GenerateRandomIntSlice(1000000),
			want:  sort.IntSorted,
		},
		{
			name:       "one_item",
			input:      []int{},
			optimalWay: true,
			expected:   []int{},
		},
		{
			name:       "random_unsorted_slice_optimized",
			input:      []int{64, 34, 25, 12, 22, 11, 90},
			optimalWay: true,
			expected:   []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:       "random_unsorted_slice_20_optimized",
			input:      sort.GenerateRandomIntSlice(20),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "random_unsorted_slice_50_optimized",
			input:      sort.GenerateRandomIntSlice(50),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "random_unsorted_slice_100_optimized",
			input:      sort.GenerateRandomIntSlice(100),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "random_unsorted_slice_1000_optimized",
			input:      sort.GenerateRandomIntSlice(1000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "random_unsorted_slice_10000_optimized",
			input:      sort.GenerateRandomIntSlice(10000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "random_unsorted_slice_100000_optimized",
			input:      sort.GenerateRandomIntSlice(100000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "random_unsorted_slice_1000000_optimized",
			input:      sort.GenerateRandomIntSlice(1000000),
			optimalWay: true,
			want:       sort.IntSorted,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual []int
			if test.optimalWay {
				actual = countingOptimizedSort(test.input)
			} else {
				actual = countingSort(test.input)
			}

			if test.expected != nil && !assert.Equal(t, test.expected, actual) {
				t.Errorf("Result return a value not expected %v\n", actual)
				t.FailNow()
			}
			if test.want != nil && !test.want(actual) {
				t.Errorf("Result return a value not expected %v\n", actual)
				t.FailNow()
			}
		})
	}
}
