package sorted_arrays

import "testing"

func TestProblem(t *testing.T) {
	var tests = []struct {
		name     string
		inputA   []int
		inputB   []int
		expected float64
	}{
		{
			name:     "first",
			inputA:   []int{1, 3},
			inputB:   []int{2},
			expected: 2,
		},
		{
			name:     "second",
			inputA:   []int{1, 2},
			inputB:   []int{3, 4},
			expected: 2.5,
		},
		{
			name:     "empty",
			inputA:   []int{},
			inputB:   []int{1},
			expected: 1,
		},
		{
			name:     "negative_edge_case_1",
			inputA:   []int{0, 0, 0, 0, 0},
			inputB:   []int{-1, 0, 0, 0, 0, 0, 1},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findMedianSortedArrays(test.inputA, test.inputB)
			if result != test.expected {
				t.Log(result)
				t.FailNow()
			}
		})
	}
}
