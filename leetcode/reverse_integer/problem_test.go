package reverse_integer

import "testing"

func TestReverseInteger(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "first",
			input:    123,
			expected: 321,
		},
		{
			name:     "second",
			input:    -123,
			expected: -321,
		},
		{
			name:     "edge_case_1",
			input:    1534236469,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := reverse(test.input)
			if result != test.expected {
				t.Errorf("invalid result [%d], expected [%d]", result, test.expected)
				t.FailNow()
			}
		})
	}
}
