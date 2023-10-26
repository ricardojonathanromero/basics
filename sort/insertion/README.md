# Insertion Sort

Builds the final sorted array one item at a time.
It repeatedly takes an element from the unsorted
part and inserts it into its correct position in
the sorted part.


## Complexity
Time Complexity: **O(n^2)** in the worst and average cases.

## Description
Insertion Sort is a simple comparison-based sorting algorithm that
builds the final sorted array one item at a time. It is particularly
useful for small datasets and is known for its simplicity and
efficiency on small lists.

## How it works
1. **Initial Configuration:** The algorithm begins with a list of elements.
  The first element in the list is considered to be the sorted part, while
  the remaining elements are considered the unsorted part.

2. **Iterative Comparison:** The algorithm iterates through the unsorted part
  of the list one element at a time, starting from the second element.

3. **Insertion:** For each element in the unsorted part, the algorithm compares
  it to the elements in the sorted part. It repeatedly compares the current
  element with the elements immediately before it and shifts those elements to
  the right until it finds the correct position for the current element.

4. **Sorting Progress:** As the algorithm iterates through the unsorted part
  and performs comparisons and insertions, the sorted part of the list gradually
  grows. The result is that the list becomes progressively sorted.

5. **Completion:** The algorithm continues this process until it has iterated
  through all elements in the unsorted part. At this point, the entire list is sorted.

Insertion Sort is known for its simplicity and efficiency when dealing with
small lists or partially sorted lists. However, it has a time complexity of
O(n^2) in the worst case, which makes it less efficient than some other sorting
algorithms for large datasets. It is often used when the list is already
partially sorted or when simplicity and minimal memory usage are more important
than sorting speed.
