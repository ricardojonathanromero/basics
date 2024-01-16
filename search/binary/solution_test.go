package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	dataSet := []int{2, 5, 7, 9, 11, 14, 16, 17, 80, 85, 99, 105, 108, 135, 125}

	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{
			name:     "single_test",
			input:    dataSet,
			target:   14,
			expected: 5,
		},
		{
			name:     "item_not_exists",
			input:    dataSet,
			target:   15,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := binarySearch(test.input, test.target)

			if !assert.Equal(t, test.expected, actual) {
				t.FailNow()
			}
		})
	}
}
