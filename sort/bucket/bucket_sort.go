package bucket

import (
	"math"
	"sort"
	"sync"
)

func bucketSort(arr []float64) {
	if len(arr) <= 1 {
		return
	}

	minVal, maxVal := arr[0], arr[0]

	// Determine the Range
	for _, currVal := range arr {
		if currVal < minVal {
			minVal = currVal
		}
		if currVal > maxVal {
			maxVal = currVal
		}
	}

	// Determine the range and the number of buckets
	// bucketSize = (max - min / arr.length)
	bucketSize := (maxVal-minVal)/float64(len(arr)) + 1 // bucket size based on max - min / arr.length + 1
	// numBuckets = int(max - min / bucketSize)
	numBuckets := int(math.Ceil((maxVal - minVal) / bucketSize))

	// Create and initialize the buckets
	buckets := make([][]float64, numBuckets)
	for i := 0; i < numBuckets; i++ {
		buckets[i] = []float64{}
	}

	// Distribute Elements
	for _, number := range arr {
		bucketIdx := int((number - minVal) / bucketSize)
		buckets[bucketIdx] = append(buckets[bucketIdx], number)
	}

	// Sort Each Bucket
	for _, bucket := range buckets {
		// implements insertion sort algorithm
		insertionSortHelper(bucket)
	}

	// Concatenate Buckets
	pointer := 0
	for _, bucket := range buckets {
		for _, number := range bucket {
			arr[pointer] = number
			pointer++
		}
	}

	// return the result sorted!
}

func insertionSortHelper(arr []float64) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}

// Define the number of goroutines to use for sorting
const numGoroutines = 4

func bucketSortOptimalSolution(arr []float64) {
	if len(arr) <= 1 {
		return
	}

	// Find the minimum and maximum values in the array
	minimum := arr[0]
	maximum := arr[0]
	for _, value := range arr {
		if value < minimum {
			minimum = value
		}
		if value > maximum {
			maximum = value
		}
	}

	// Determine the range and the number of buckets
	bucketSize := ((maximum - minimum) / float64(len(arr))) + 1
	numBuckets := int(math.Ceil((maximum - minimum) / bucketSize))

	// Create and initialize the buckets
	buckets := make([][]float64, numBuckets)
	for i := 0; i < numBuckets; i++ {
		buckets[i] = []float64{}
	}

	// Distribute elements into the buckets
	for _, value := range arr {
		bucketIndex := int((value - minimum) / bucketSize)
		buckets[bucketIndex] = append(buckets[bucketIndex], value)
	}

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Sort each bucket using goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * (numBuckets / numGoroutines)
			end := (i + 1) * (numBuckets / numGoroutines)
			if i == numGoroutines-1 {
				end = numBuckets
			}
			for j := start; j < end; j++ {
				sort.Float64s(buckets[j])
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Concatenate the buckets to get the final sorted array
	index := 0
	for i := 0; i < numBuckets; i++ {
		for _, value := range buckets[i] {
			arr[index] = value
			index++
		}
	}
}
