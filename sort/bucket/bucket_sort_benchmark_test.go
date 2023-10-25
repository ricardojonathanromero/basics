package bucket

import (
	"github.com/ricardojonathanromero/basics/utils/parallelism"
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func wrapBucketSort(args ...any) any {
	return bucketSort(floatToAny(args...))
}

func wrapBucketSortOptimalSolution(args ...any) any {
	return bucketSortOptimalSolution(floatToAny(args...))
}

func floatToAny(args ...any) []float64 {
	convertedArgs := make([]float64, len(args))
	for i, v := range args {
		convertedArgs[i] = v.(float64)
	}
	return convertedArgs
}

func BenchmarkBucketSort(b *testing.B) {
	benchmarks := []struct {
		name       string
		input      []float64
		optimalWay bool
	}{
		{
			name:  "complex_unsorted_slice_normal_sol",
			input: sort.GenerateRandomFloatSlice(20),
		},
		{
			name:       "complex_unsorted_slice_optimal_sol",
			input:      sort.GenerateRandomFloatSlice(20),
			optimalWay: true,
		},
	}

	// create worker for all the benchmark tests
	worker := parallelism.NewParallel()

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// iterate until count parameter. by default 1
				var to parallelism.To
				var next parallelism.Next

				if bm.optimalWay {
					to = wrapBucketSortOptimalSolution
				} else {
					to = wrapBucketSort
				}

				next = func(arr any) {
					numbers, ok := arr.([]float64)
					if !ok {
						b.Errorf("result is not type []int, %v\n", arr)
						b.FailNow()
					}

					if !sort.FloatSorted(numbers) {
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
