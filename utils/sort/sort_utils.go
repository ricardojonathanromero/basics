package sort

import (
	"math/rand"
	"time"
)

type ArrIntegerSorted func(arr []int) bool

type ArrFloat64Sorted func(arr []float64) bool

func GenerateRandomIntSlice(size int) []int {
	rand.NewSource(time.Now().UnixNano())
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000) // Adjust the upper limit as needed
	}
	return slice
}

func IntSorted(arr []int) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func FloatSorted(arr []float64) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func GenerateRandomFloatSlice(size int) []float64 {
	rand.NewSource(time.Now().UnixNano())
	slice := make([]float64, size)
	for i := range slice {
		slice[i] = rand.Float64()
	}
	return slice
}
