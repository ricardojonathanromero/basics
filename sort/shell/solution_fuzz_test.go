package shell

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzShellSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	shellSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(input) {
		t.Errorf("Shell Sort failed to sort the array: %v", input)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("shell_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzShellSort(t, s, input)
	})
}
