package bfs

import (
	"testing"
)

func BenchmarkBreadthFirstSearch(b *testing.B) {
	benchmarks := []struct {
		name     string
		input    map[int]Node
		start    int
		end      int
		expected string
	}{
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
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = breadthFirstSearch(bm.input, bm.start, bm.end)
			}
		})
	}
}
