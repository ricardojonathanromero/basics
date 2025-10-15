package divisible_sum_pairs

import "testing"

func TestDivisibleSumPairs(t *testing.T) {
	var tests = []struct {
		name     string
		n        int32
		k        int32
		ar       []int32
		expected int32
	}{
		{
			name:     "first_test",
			n:        6,
			k:        3,
			ar:       []int32{1, 3, 2, 6, 1, 2},
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := divisibleSumPairs(test.n, test.k, test.ar)
			if result != int(test.expected) {
				t.Errorf("error result [%d], expected [%d]", result, test.expected)
				t.FailNow()
			}
		})
	}
}
