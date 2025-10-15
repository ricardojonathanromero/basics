package bfs

import (
	"testing"
)

func TestPower(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "power zero",
			input:    1,
			expected: true,
		},
		{
			name:     "power four",
			input:    16,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isPowerOfTwo(test.input)

			if result != test.expected {
				t.Error("failed")
				t.FailNow()
			}
		})
	}
}
