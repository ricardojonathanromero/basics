package shell

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func BenchmarkShellSort(b *testing.B) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "complex_unsorted_slice",
			input: sort.GenerateRandomIntSlice(50),
		},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				shellSort(test.input)

				if !sort.IntSorted(test.input) {
					b.Error("!result is not sorted")
					b.FailNow()
				}
			}
		})
	}
}
