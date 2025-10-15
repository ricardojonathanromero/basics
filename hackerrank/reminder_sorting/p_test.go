package reminder_sorting

import "testing"

func TestRemainderSorting(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "first",
			input:  []string{"a", "ab", "bc", "abc"},
			expect: []string{"abc", "a", "ab", "bc"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := remainderSorting(test.input)
			if len(result) != len(test.expect) {
				t.Errorf("error result lenght %d, expected %d", len(result), len(test.expect))
				t.FailNow()
			}
		})
	}
}
