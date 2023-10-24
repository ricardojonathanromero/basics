package bubble

import (
	"math/rand"
	"testing"
	"time"
)

func generateRandomSlice(size int) []int {
	rand.NewSource(time.Now().UnixNano())
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000) // Adjust the upper limit as needed
	}
	return slice
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Create a random slice for each benchmark run
		slice := generateRandomSlice(1000) // Adjust the size as needed
		bubbleSort(slice)
	}
}
