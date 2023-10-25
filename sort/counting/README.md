# Counting Sort

An integer sorting algorithm that operates by counting
the number of objects that have each distinct key value
and using arithmetic to determine the position of
each key.


## Complexity
Time Complexity: **O(n + k)** where n is the number of
elements and k is the range of input.

## Description
Counting Sort is a linear time sorting algorithm that is often used
when the range of input values is not significantly larger than the
number of elements to be sorted. It's an integer sorting algorithm
that works by counting the number of occurrences of each element in
the input array and using that information to place each element in
its correct sorted position.

## How it works
1. **Determine the Range:** Find the range of input values (i.e.,
  the minimum and maximum values in the array) to determine the size of
  the counting array.

2. **Count Occurrences:** Create a counting array (also called a
  frequency array) with a size equal to the range of values. For example,
  if the range is from min to max, the counting array should have a size
  of max - min + 1.

3. **Count Frequencies:** Traverse the input array and, for each element,
  increment the count in the counting array corresponding to that element's
  value. This step counts how many times each element appears in the input.

4. **Calculate Cumulative Counts:** Modify the counting array to store the cumulative
  count of elements. This cumulative count represents the position of the last
  occurrence of each element in the sorted output.

5. **Build the Sorted Output:** Create an output array of the same size as the input 
  array. Starting from the end of the input array (for stability), for each
  element, look up its count in the counting array, subtract one (to account
  for zero-based indexing), and use this value as the index to place the element
  in the sorted output array.

6. **Decrement Counts:** After placing an element in the output array, decrement its
  count in the counting array.

7. **Repeat for All Elements:** Repeat step 5 for all elements in the input array.
  This process ensures that elements are placed in the correct sorted order.

8. **Final Output:** The output array now contains the sorted elements, and the
  Counting Sort is complete.

Counting Sort is very efficient when the range of input values is small compared to
the number of elements and is particularly well-suited for sorting non-negative
integers or small integers. It's a stable sorting algorithm, which means that the
relative order of equal elements is preserved. However, Counting Sort is not a
comparison-based sort, and its performance can degrade when the range of input
values is large.
