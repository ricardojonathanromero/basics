package string_integer

import "testing"

func TestStringToInteger(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "edge_case_1",
			input:    "",
			expected: 0,
		},
		{
			name:     "edge_case_1",
			input:    "-",
			expected: 0,
		},
		{
			name:     "first",
			input:    "42",
			expected: 42,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := myAtoi(test.input)
			if test.expected != result {
				t.Errorf("error result [%d], expected [%d]", result, test.expected)
				t.FailNow()
			}
		})
	}
}
