# Quick Sort

A divide-and-conquer algorithm that works by selecting a
'pivot' element from the array and partitioning the other
elements into two sub-arrays.


## Complexity
Time Complexity: **O(n^2)** in the worst case, but **O(n log n)**
on average.

## Description
Quick Sort is a highly efficient, comparison-based sorting algorithm that
falls into the category of divide-and-conquer algorithms. It is known for
its speed and is widely used in practice.

## How it works
1. **Choosing a Pivot:** The first step in Quick Sort is to select a pivot element from the
  array. The pivot element serves as a reference point for the sorting process. The choice
  of the pivot can greatly influence the algorithm's performance. Common strategies include
  selecting the first, last, or middle element, or even using a random element.

2. **Partitioning:** The array is partitioned into two sub-arrays: elements less than the
  pivot (the left sub-array) and elements greater than the pivot (the right sub-array).
  The pivot is now in its final sorted position.

3. **Recursive Sorting:** Quick Sort is applied recursively to the left and right
  sub-arrays. The left sub-array is sorted in ascending order, and the right sub-array is
  sorted in ascending order. This process continues until the entire array is sorted.

4. **Combine Sorted Sub-Arrays:** As the recursive calls return, the sorted sub-arrays are
  combined to produce the final sorted array. This combination is achieved simply by
  concatenating the left sub-array, the pivot, and the right sub-array.


## More about
1. **Efficiency:** Quick Sort is known for its efficiency and often outperforms other
  sorting algorithms in practice. On average, it has a time complexity of O(n log n),
  where n is the number of elements in the array. In the worst case, it can degrade
  to O(n^2), but good pivot selection strategies and randomization can mitigate this risk.

2. **In-Place Sorting:** Quick Sort can be performed in-place, which means it doesn't
  require additional memory allocation. It only requires a small amount of extra memory
  for the recursive call stack.

3. **Stability:** Quick Sort is not a stable sorting algorithm, which means it may change
  the relative order of equal elements.

The choice of the pivot element and the partitioning strategy are crucial for Quick Sort's
performance. Several variations and optimizations of Quick Sort exist, including choosing
the pivot based on the median of three values, dual pivot Quick Sort, and hybrid sorting
strategies. These optimizations aim to make Quick Sort more robust and efficient.
