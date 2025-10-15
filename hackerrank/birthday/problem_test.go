package birthday

import "testing"

func TestBirthday(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int32
		day      int32
		month    int32
		expected int32
	}{
		{
			name:     "one",
			input:    []int32{1, 2, 1, 3, 2},
			day:      3,
			month:    2,
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := birthday(test.input, test.day, test.month)
			if result != test.expected {
				t.Errorf("error result [%d], expected [%d]", result, test.expected)
				t.FailNow()
			}
		})
	}
}
