package shell

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func TestShellSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  sort.ArrIntegerSorted
	}{
		{
			name:  "simple_test",
			input: []int{7, 6, 8, 9, 3, 2, 10, 5, 1},
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_10",
			input: sort.GenerateRandomIntSlice(10),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_100",
			input: sort.GenerateRandomIntSlice(100),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_1000",
			input: sort.GenerateRandomIntSlice(1000),
			want:  sort.IntSorted,
		},
		{
			name:  "simple_test_10000",
			input: sort.GenerateRandomIntSlice(10000),
			want:  sort.IntSorted,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shellSort(test.input)

			if test.want != nil && !test.want(test.input) {
				t.FailNow()
			}
		})
	}
}
