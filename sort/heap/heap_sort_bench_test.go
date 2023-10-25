package heap

import (
	"github.com/ricardojonathanromero/basics/utils/parallelism"
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func wrapHeapSort(args ...any) any {
	return heapSort(intToAny(args...))
}

func wrapHeapSortOptimalSolution(args ...any) any {
	return heapSortOptimal(intToAny(args...))
}

func intToAny(args ...any) []int {
	convertedArgs := make([]int, len(args))
	for i, v := range args {
		convertedArgs[i] = v.(int)
	}
	return convertedArgs
}

func BenchmarkHeapSort(b *testing.B) {
	tests := []struct {
		name       string
		input      []int
		optimalWay bool
	}{
		{
			name:  "complex_unsorted_slice",
			input: sort.GenerateRandomIntSlice(50),
		},
		{
			name:       "complex_unsorted_slice_optimal",
			input:      sort.GenerateRandomIntSlice(50),
			optimalWay: true,
		},
	}

	// create worker for all the benchmark tests
	worker := parallelism.NewParallel()

	// iterate tests
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			// iterate until count parameter. by default 1
			for i := 0; i < b.N; i++ {
				var to parallelism.To
				var next parallelism.Next

				if test.optimalWay {
					to = wrapHeapSortOptimalSolution
				} else {
					to = wrapHeapSort
				}

				next = func(arr any) {
					numbers, ok := arr.([]int)
					if !ok {
						b.Errorf("result is not type []int, %v\n", arr)
						b.FailNow()
					}

					if !sort.IntSorted(numbers) {
						b.Error("!result is not sorted")
						b.FailNow()
					}
				}

				worker.RunInParallel(to, next)
			}

			// wait for results
			worker.Wait()
		})
	}
}
