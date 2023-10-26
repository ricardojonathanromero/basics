package bucket

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := []struct {
		name       string
		input      []float64
		optimalWay bool
		expect     []float64
		want       sort.ArrFloat64Sorted
	}{
		{
			name:   "random input",
			input:  []float64{292, 23, 453, 76, 63, 8},
			expect: []float64{8, 23, 63, 76, 292, 453},
		},
		{
			name:   "empty slice",
			input:  []float64{},
			expect: []float64{},
		},
		{
			name:   "single-element slice",
			input:  []float64{42},
			expect: []float64{42},
		},
		{
			name:   "sorted slice",
			input:  []float64{1, 2, 3, 4, 5},
			expect: []float64{1, 2, 3, 4, 5},
		},
		{
			name:   "reverse-sorted slice",
			input:  []float64{5, 4, 3, 2, 1},
			expect: []float64{1, 2, 3, 4, 5},
		},
		{
			name:   "random unsorted slice",
			input:  []float64{64, 34, 25, 12, 22, 11, 90},
			expect: []float64{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "complex unsorted slice",
			input: sort.GenerateRandomFloatSlice(1000),
			want:  sort.FloatSorted,
		},
		{
			name:       "random input",
			input:      []float64{292, 23, 453, 76, 63, 8},
			optimalWay: true,
			expect:     []float64{8, 23, 63, 76, 292, 453},
		},
		{
			name:       "empty slice",
			input:      []float64{},
			optimalWay: true,
			expect:     []float64{},
		},
		{
			name:       "single-element slice",
			input:      []float64{42},
			optimalWay: true,
			expect:     []float64{42},
		},
		{
			name:       "sorted slice",
			input:      []float64{1, 2, 3, 4, 5},
			optimalWay: true,
			expect:     []float64{1, 2, 3, 4, 5},
		},
		{
			name:       "reverse-sorted slice",
			input:      []float64{5, 4, 3, 2, 1},
			optimalWay: true,
			expect:     []float64{1, 2, 3, 4, 5},
		},
		{
			name:       "random unsorted slice",
			input:      []float64{64, 34, 25, 12, 22, 11, 90},
			optimalWay: true,
			expect:     []float64{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:       "complex unsorted slice",
			input:      sort.GenerateRandomFloatSlice(1000),
			optimalWay: true,
			want:       sort.FloatSorted,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.optimalWay {
				bucketSortOptimalSolution(test.input)
			} else {
				bucketSort(test.input)
			}

			if test.expect != nil && !assert.Equal(t, test.expect, test.input) {
				t.Errorf("test %s does not return the expected value\n", test.name)
				t.FailNow()
			} else if test.want != nil && !test.want(test.input) {
				t.Errorf("test %s does not return the expected value\n", test.name)
				t.FailNow()
			}
		})
	}
}
