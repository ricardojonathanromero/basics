package heap

import "container/heap"

/*
1.build max heap
*/
func heapSort(arr []int) {
	// get length
	length := len(arr)
	middle := length/2 - 1

	for root := middle; root >= 0; root-- {
		heapify(arr, length, root)
	}

	// Extract elements one by one from the max-heap
	for last := length - 1; last > 0; last-- {
		// Swap the root (largest element) with the current last element
		arr[0], arr[last] = arr[last], arr[0]

		// Call heapify on the reduced heap (excluding the last element)
		heapify(arr, last, 0)
	}
}

func heapify(arr []int, length, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	// Find the largest element among root, left child, and right child
	if left < length && arr[left] > arr[largest] {
		largest = left
	}
	if right < length && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest element is not the root, swap and heapify the affected sub-tree
	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, length, largest)
	}
}

func heapSortOptimal(arr []int) {
	h := IntHeap(arr)
	//h := &IntHeap(arr)

	// Build a max-heap using the container/heap package
	heap.Init(&h)

	// Extract elements one by one from the max-heap to sort the array
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = heap.Pop(&h).(int)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
