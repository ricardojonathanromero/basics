# Heap Sort

A comparison-based sorting algorithm that uses a binary heap
data structure.


## Complexity
Time Complexity: **O(n log n)** in all cases.

## Description
Heap Sort is a comparison-based sorting algorithm that is based on
the data structure known as a binary heap. It is an in-place sorting
algorithm, which means it sorts the elements within the original
array without requiring additional memory allocation.

## How it works
1. **Heapify:** The first step in Heap Sort is to create a binary heap
  from the input array. A binary heap is a complete binary tree that
  satisfies the "heap property," which means that each parent node has
  a value less than or equal to the values of its children. This process
  is known as "heapify."

2. **Building the Heap:** The input array is considered to be a binary
  heap of size N. The heap is built by iteratively applying the heapify
  operation to subtrees of the array, starting from the bottom of the
  tree and moving upwards. This process ensures that the largest element
  in the array (the root of the heap) is at the top.

3. **Sorting the Heap:** Once the heap is built, the largest element
  (at the root) is swapped with the last element in the array, effectively
  moving it to its correct sorted position at the end of the array. The heap
  size is reduced by one, and the heapify operation is applied to the new root
  (which may violate the heap property).

4. **Repeat:** Steps 3 and 4 are repeated for the remaining elements of the
  heap until the entire array is sorted. In each iteration, the largest unsorted
  element is moved to the sorted part of the array.

5. **Final Sorted Array:** Once all iterations are complete, the original array
  is sorted in ascending order.

Heap Sort is an efficient sorting algorithm with a worst-case time complexity of 
O(n log n), making it suitable for large datasets. However, it is not a stable
sorting algorithm, meaning that the relative order of equal elements may not be
preserved. Additionally, it requires additional memory for the heap data structure.
Overall, Heap Sort is a practical choice when a stable sort is not required, and
you want a sorting algorithm with guaranteed time complexity.
