package find_word

import "testing"

func TestFindWord(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		expect string
	}{
		{
			name:   "first",
			input:  []string{"P>E", "E>R", "R>U"},
			expect: "PERU",
		},
		{
			name:   "example",
			input:  []string{"U>N", "G>A", "R>Y", "H>U", "N>G", "A>R"},
			expect: "HUNGARY",
		},
		{
			name:   "example",
			input:  []string{"I>F", "W>I", "S>W", "F>T"},
			expect: "SWIFT",
		},
		{
			name:   "example",
			input:  []string{"R>T", "A>L", "P>O", "O>R", "G>A", "T>U", "U>G"},
			expect: "PORTUGAL",
		},
		{
			name:   "example",
			input:  []string{"W>I", "R>L", "T>Z", "Z>E", "S>W", "E>R", "L>A", "A>N", "N>D", "I>T"},
			expect: "SWITZERLAND",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findWord(test.input)
			if result != test.expect {
				t.Errorf("error result [%s], expect [%s]", result, test.expect)
				t.FailNow()
			}
		})
	}
}
