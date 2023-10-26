## Bubble Sort

Another divide-and-conquer algorithm that divides the input
array into two halves, recursively sorts them, and then
merges the two sorted halves.


## Complexity
Time Complexity: **O(n log n)** in all cases.

## Description
Merge Sort is a divide-and-conquer sorting algorithm that works by
dividing the input array into smaller sub-arrays, sorting these
sub-arrays, and then merging them back together to produce a fully
sorted array.

## How it works
1. **Divide:** The input array is divided into two equal halves (or nearly
  equal if the size is odd). This process continues recursively until each
  sub-array contains only one element. These one-element sub-arrays are
  considered sorted.

2. **Conquer:** Once you have reached the base case of one-element sub-arrays,
  you start merging them back together. Merging involves comparing elements in
  the sub-arrays and arranging them in sorted order.

3. **Merge:** Merging is done by creating a new array, and elements from the
  two sub-arrays are copied to the new array in sorted order. The merging process
  continues until you have combined all the sub-arrays into one fully sorted array.

4. **Recursion:** The merge operation is performed in a bottom-up manner, meaning
  that you first merge the one-element sub-arrays, then pairs of sub-arrays, and so
  on, until the entire array is sorted.
