package long_substring

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "simple",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "single",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "complex",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "space",
			input:    " ",
			expected: 1,
		},
		{
			name:     "edge_case_1",
			input:    "dvdf",
			expected: 3,
		},
		{
			name:     "edge_case_2",
			input:    "ckilbkd",
			expected: 5,
		},
		{
			name:     "edge_case_3",
			input:    "abba",
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(test.input)
			t.Log(result)

			if result != test.expected {
				t.FailNow()
			}
		})
	}
}
