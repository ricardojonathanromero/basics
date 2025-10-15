package longest_palindromic

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "first",
			input:  "babad",
			expect: "aba",
		},
		{
			name:   "second",
			input:  "cbbd",
			expect: "bb",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := longestPalindrome(test.input)
			if result != test.expect {
				t.Errorf("result [%s] is diff as expected [%s]", result, test.expect)
				t.FailNow()
			}
		})
	}
}
