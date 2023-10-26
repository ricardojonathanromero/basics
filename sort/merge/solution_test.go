package merge

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		optimalWay bool
		expect     []int
		want       sort.ArrIntegerSorted
	}{
		{
			name:   "two_array",
			input:  []int{6, 2},
			expect: []int{2, 6},
		},
		{
			name:  "three_array",
			input: []int{6, 2, 89},
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test",
			input: []int{6, 21, 2, 45, 90, 0, 84, 60},
			want:  sort.IntSorted,
		},
		{
			name:  "simple_merge_sort_100",
			input: sort.GenerateRandomIntSlice(100),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_merge_sort_1000",
			input: sort.GenerateRandomIntSlice(100),
			want:  sort.IntSorted,
		},
		{
			name:       "two_array_optimized",
			input:      []int{6, 2},
			optimalWay: true,
			expect:     []int{2, 6},
		},
		{
			name:       "three_array_optimized",
			input:      []int{6, 2, 89},
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "optimized_test",
			input:      []int{6, 21, 2, 45, 90, 0, 84, 60},
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "optimized_merge_sort_100",
			input:      sort.GenerateRandomIntSlice(100),
			optimalWay: true,
			want:       sort.IntSorted,
		},
		{
			name:       "optimized_merge_sort_1000",
			input:      sort.GenerateRandomIntSlice(100),
			optimalWay: true,
			want:       sort.IntSorted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual []int
			if tt.optimalWay {
				actual = mergeSortOptimized(tt.input)
			} else {
				actual = mergeSort(tt.input)
			}

			if tt.expect != nil && !assert.Equal(t, tt.expect, actual) {
				t.FailNow()
			}

			if tt.want != nil && !tt.want(actual) {
				t.FailNow()
			}
		})
	}
}
