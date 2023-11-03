package bfs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFSAlgorithm(t *testing.T) {
	tests := []struct {
		name       string
		input      map[int]Node
		optimalWay bool
		start      int
		end        int
		expected   string
	}{
		{
			name: "simple_test",
			input: map[int]Node{
				1: {Value: 1, Neighbors: []int{2, 3}},
				2: {Value: 2, Neighbors: []int{4, 5}},
				3: {Value: 3, Neighbors: []int{6}},
				4: {Value: 4, Neighbors: []int{}},
				5: {Value: 5, Neighbors: []int{}},
				6: {Value: 6, Neighbors: []int{}},
			},
			start:    1,
			end:      6,
			expected: "1->3->6", //1->3->6
		},
		{
			name: "not_found",
			input: map[int]Node{
				1: {Value: 1, Neighbors: []int{2, 3}},
			},
			start:    8,
			end:      6,
			expected: "not_found",
		},
		{
			name: "complex_test",
			input: map[int]Node{
				0:  {Value: 0, Neighbors: []int{9, 7, 11}},
				9:  {Value: 9, Neighbors: []int{10, 8}},
				7:  {Value: 7, Neighbors: []int{3, 6, 11}},
				11: {Value: 11, Neighbors: []int{}},
				10: {Value: 10, Neighbors: []int{1}},
				8:  {Value: 8, Neighbors: []int{1, 12}},
				3:  {Value: 3, Neighbors: []int{2, 4}},
				6:  {Value: 6, Neighbors: []int{5}},
				1:  {Value: 1, Neighbors: []int{}},
				12: {Value: 12, Neighbors: []int{2}},
				2:  {Value: 2, Neighbors: []int{}},
				4:  {Value: 4, Neighbors: []int{}},
				5:  {Value: 5, Neighbors: []int{}},
			},
			start:    0,
			end:      5,
			expected: "0->7->6->5",
		},
		{
			name: "optimized_test",
			input: map[int]Node{
				1: {Value: 1, Neighbors: []int{2, 3}},
				2: {Value: 2, Neighbors: []int{4, 5}},
				3: {Value: 3, Neighbors: []int{6}},
				4: {Value: 4, Neighbors: []int{}},
				5: {Value: 5, Neighbors: []int{}},
				6: {Value: 6, Neighbors: []int{}},
			},
			optimalWay: true,
			start:      1,
			end:        6,
			expected:   "1->3->6", //1->3->6
		},
		{
			name:       "not_found_optimal",
			optimalWay: true,
			input: map[int]Node{
				1: {Value: 1, Neighbors: []int{2, 3}},
			},
			start:    8,
			end:      6,
			expected: "not_found",
		},
		{
			name:       "complex_optimal_test",
			optimalWay: true,
			input: map[int]Node{
				0:  {Value: 0, Neighbors: []int{9, 7, 11}},
				9:  {Value: 9, Neighbors: []int{10, 8}},
				7:  {Value: 7, Neighbors: []int{3, 6, 11}},
				11: {Value: 11, Neighbors: []int{}},
				10: {Value: 10, Neighbors: []int{1}},
				8:  {Value: 8, Neighbors: []int{1, 12}},
				3:  {Value: 3, Neighbors: []int{2, 4}},
				6:  {Value: 6, Neighbors: []int{5}},
				1:  {Value: 1, Neighbors: []int{}},
				12: {Value: 12, Neighbors: []int{2}},
				2:  {Value: 2, Neighbors: []int{}},
				4:  {Value: 4, Neighbors: []int{}},
				5:  {Value: 5, Neighbors: []int{}},
			},
			start:    0,
			end:      5,
			expected: "0->7->6->5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string
			if test.optimalWay {
				result = breadthFirstSearchOptimized(test.input, test.start, test.end)
			} else {
				result = breadthFirstSearch(test.input, test.start, test.end)
			}

			if !assert.Equal(t, test.expected, result) {
				t.FailNow()
			}
		})
	}
}
