package bubble

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single-element slice",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "sorted slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse-sorted slice",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random unsorted slice",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := bubbleSort(test.input)
			if !assert.Equal(t, test.expected, result) {
				t.Errorf("Test case %s failed. Expected: %v, Got: %v", test.name, test.expected, result)
				t.Fail()
			}
		})
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := bubbleOptimalSolution(test.input)
			if !assert.Equal(t, test.expected, result) {
				t.Errorf("Test case %s failed. Expected: %v, Got: %v", test.name, test.expected, result)
				t.Fail()
			}
		})
	}
}
