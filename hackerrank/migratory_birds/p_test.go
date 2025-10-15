package migratory_birds

import "testing"

func TestMigratoryBirds(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int32
		expected int32
	}{
		{
			name:     "first",
			input:    []int32{1, 4, 4, 4, 5, 3},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := migratoryBirds(test.input)
			if result != test.expected {
				t.Errorf("error result [%d], expected [%d", result, test.expected)
				t.FailNow()
			}
		})
	}
}
