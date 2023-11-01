package selection

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  sort.ArrIntegerSorted
	}{
		{
			name:  "simple_test",
			input: []int{2, 8, 5, 3, 9, 4, 1},
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
			selectionSort(test.input)

			if test.want != nil && !test.want(test.input) {
				t.FailNow()
			}
		})
	}
}
