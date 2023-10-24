package bubble

import (
	"testing"
)

type arrSorted func(arr []int) bool

func isSorted(arr []int) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func FuzzBubbleSort(f *testing.F) {
	tests := []struct {
		name string
		args []int
		want arrSorted
	}{
		{
			name: "random numbers",
			args: generateRandomSlice(1000),
			want: isSorted,
		},
	}

	for _, tt := range tests {
		f.Add(tt.name)
		f.Fuzz(func(t *testing.T, name string) {
			t.Logf("start %s test", tt.name)
			// Sort the array using Bubble Sort
			arr := bubbleSort(tt.args)

			// Verify that the array is sorted
			if !tt.want(arr) {
				t.Errorf("Bubble Sort failed to sort the array: %v", arr)
			}
		})
	}
}
