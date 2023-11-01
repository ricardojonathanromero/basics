# Selection Sort

Divides the input list into two parts: the sublist
of items already sorted and the sublist of items
to be sorted.


## Complexity
Time Complexity: **O(n^2)** in the worst and average cases.

## Description
Selection Sort is a simple and intuitive comparison-based sorting algorithm.
It works by repeatedly finding the minimum (or maximum, depending on the
desired order) element from the unsorted portion of the array and moving it
to the beginning of the sorted portion.

## How it works
1. **Initialization:** The algorithm starts with an empty sorted portion at the beginning
  and the entire unsorted portion at the end. Initially, the sorted portion is empty, and
  the unsorted portion contains all elements.

2. **Finding the Minimum:** Selection Sort iterates through the unsorted portion of the
  array to find the minimum element. It does this by comparing each element with the current
  minimum element found so far. If a smaller element is found, it becomes the new minimum.

3. **Swapping:** Once the minimum element is found, it is swapped with the first element in
  the unsorted portion. This effectively moves the minimum element to the end of the sorted
  portion.

4. **Expanding the Sorted Portion:** The sorted portion is expanded by one element, and the
  unsorted portion is reduced by one element.

5. **Repeat:** Steps 2-4 are repeated until the entire array is sorted. In each iteration,
  the minimum element is found and placed in its correct position in the sorted portion.
