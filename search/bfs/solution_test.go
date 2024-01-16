package bfs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFSAlgorithm(t *testing.T) {
	tests := []struct {
		name        string
		input       map[int]Node
		inputStr    map[string]NodeStr
		optimalWay  bool
		strSolution bool
		start       int
		startStr    string
		end         int
		endStr      string
		expected    string
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
		{
			name:       "edge_optimal_test",
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
			start:    11,
			end:      0,
			expected: "not_found",
		},
		{
			name:        "test_string_example",
			strSolution: true,
			inputStr: map[string]NodeStr{
				"John":     {Value: "John", Neighbors: []string{"George", "Sam", "Edward"}},
				"George":   {Value: "George", Neighbors: []string{"Richard"}},
				"Sam":      {Value: "Sam", Neighbors: []string{"Richard", "Briana"}},
				"Edward":   {Value: "Edward", Neighbors: []string{"Anett", "Shaun"}},
				"Richard":  {Value: "Richard", Neighbors: []string{"Franklin"}},
				"Briana":   {Value: "Briana", Neighbors: []string{"Lynsey", "Karen"}},
				"Anett":    {Value: "Anett", Neighbors: []string{"Wilson"}},
				"Shaun":    {Value: "Shaun", Neighbors: []string{}},
				"Franklin": {Value: "Franklin", Neighbors: []string{}},
				"Lynsey":   {Value: "Lynsey", Neighbors: []string{}},
				"Karen":    {Value: "Karen", Neighbors: []string{}},
				"Wilson":   {Value: "Wilson", Neighbors: []string{}},
			},
			startStr: "John",
			endStr:   "Anett",
			expected: "John->Edward->Anett",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string
			if test.optimalWay {
				result = breadthFirstSearchOptimized(test.input, test.start, test.end)
			} else if test.strSolution {
				result = breadthFirstSearchStr(test.inputStr, test.startStr, test.endStr)
			} else {
				result = breadthFirstSearch(test.input, test.start, test.end)
			}

			if !assert.Equal(t, test.expected, result) {
				t.FailNow()
			}
		})
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	tree := map[string]NodeStr{
		"John":     {Value: "John", Neighbors: []string{"George", "Sam", "Edward"}},
		"George":   {Value: "George", Neighbors: []string{"Richard"}},
		"Sam":      {Value: "Sam", Neighbors: []string{"Richard", "Briana"}},
		"Edward":   {Value: "Edward", Neighbors: []string{"Anett", "Shaun"}},
		"Richard":  {Value: "Richard", Neighbors: []string{"Franklin"}},
		"Briana":   {Value: "Briana", Neighbors: []string{"Lynsey", "Karen"}},
		"Anett":    {Value: "Anett", Neighbors: []string{"Wilson"}},
		"Shaun":    {Value: "Shaun", Neighbors: []string{}},
		"Franklin": {Value: "Franklin", Neighbors: []string{}},
		"Lynsey":   {Value: "Lynsey", Neighbors: []string{}},
		"Karen":    {Value: "Karen", Neighbors: []string{}},
		"Wilson":   {Value: "Wilson", Neighbors: []string{}},
	}

	tests := []struct {
		name     string
		input    map[string]NodeStr
		start    string
		end      string
		expected string
	}{
		{
			name:     "john_anett_example",
			input:    tree,
			start:    "John",
			end:      "Anett",
			expected: "John->Edward->Anett",
		},
		{
			name:     "roman_root_not_found",
			input:    tree,
			start:    "Roman",
			end:      "Anett",
			expected: "not_found",
		},
		{
			name:     "emili_target_not_found",
			input:    tree,
			start:    "John",
			end:      "Emili",
			expected: "not_found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := breadthFirstSearchStr(test.input, test.start, test.end)

			if !assert.Equal(t, test.expected, result) {
				t.FailNow()
			}
		})
	}
}

func TestExample(t *testing.T) {
	performers := []Score{{Score: 99.9}, {Score: 3.1}, {Score: 1.0}}
	venues := []Score{{Score: 88.1}, {Score: 12.43}}
	events := []Score{{Score: 100.0}, {Score: 90.9}, {Score: 11.32}}

	// Assign SortOrder for each list
	assignCombinedSortOrder(performers, venues, events)

	// Display the modified lists
	fmt.Println("performers:", performers)
	fmt.Println("venues:", venues)
	fmt.Println("events:", events)
}
