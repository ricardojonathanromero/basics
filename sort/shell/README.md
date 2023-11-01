# Shell Sort

A variation of insertion sort that breaks the original
list into smaller sub-lists and sorts those sub-lists
using insertion sort.


## Complexity
Time Complexity: **O(n^1.25)** in the average case.

## Description
Shell Sort is an efficient and adaptive sorting algorithm that's an extension of the
Insertion Sort algorithm. It was designed to improve upon the Insertion Sort's
performance by reducing the number of comparisons and swaps needed for sorting.

## How it works
1. **Gap Sequence:** The core idea of Shell Sort is to divide the array into smaller
  subarrays, and then sort these subarrays using Insertion Sort. To achieve this, Shell
  Sort uses a sequence of integers called the "gap sequence." This sequence determines
  the number of elements between the items compared and swapped during each pass.

2. **Initial Gap:** Shell Sort starts with a relatively large gap (often determined by a
  mathematical sequence) and divides the array into several subarrays. The elements within
  each subarray are then sorted using Insertion Sort. This step is often referred to as a
  "preliminary pass."

3. **Shrinking Gap:** After the preliminary pass, the gap is reduced (commonly halved) to
  create smaller subarrays. These smaller subarrays are again sorted using Insertion Sort.

4. **Repetition:** The process of reducing the gap and sorting the subarrays continues until
  the gap becomes 1. The final pass, with a gap of 1, is an ordinary Insertion Sort, but since
  the elements have been partially sorted during previous passes, it is more efficient than a
  regular Insertion Sort.
