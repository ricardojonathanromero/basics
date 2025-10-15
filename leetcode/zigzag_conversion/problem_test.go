package zigzag_conversion

import "testing"

func TestConvert(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		rows     int
		expected string
	}{
		{
			name:     "first",
			input:    "PAYPALISHIRING",
			rows:     3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "second",
			input:    "PAYPALISHIRING",
			rows:     4,
			expected: "PINALSIGYAHRPI",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := convert(test.input, test.rows)
			if result != test.expected {
				t.Errorf("result [%s] not expected [%s]", result, test.expected)
				t.FailNow()
			}
		})
	}
}
