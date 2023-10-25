# Bubble Sort

A simple sorting algorithm that repeatedly steps through the list,
compares adjacent elements, and swaps them if they are in the
wrong order.


## Complexity
Time Complexity: **O(n^2)** in the worst and average cases.

## Exercise
Write a Go program that performs the following tasks:

1. Create an unsorted slice of integers.
2. Implement the Bubble Sort algorithm to sort the slice in ascending order.


## Explanation
Bubble Sort is a simple comparison-based sorting algorithm that repeatedly
steps through the list, compares adjacent elements, and swaps them if they
are in the wrong order. The pass through the list is repeated until the
list is sorted. It is called "Bubble Sort" because the smaller or larger
elements "bubble" to the top of the list during each pass.

1. **Start at the Beginning**: Bubble Sort begins by comparing the first two
    elements of the list.

2. **Compare and Swap**: It compares these two elements and swaps them if
    they are in the wrong order (e.g., if the first element is larger than
    the second one).

3. **Continue to the Next Pai**r: Bubble Sort then moves on to compare the
    second and third elements, and then the third and fourth, and so on,
    continuously comparing and swapping adjacent elements until it reaches
    the end of the list.

4. **Repeat the Pass**: Once it reaches the end of the list, Bubble Sort
    starts another pass through the list, repeating steps 1 to 3.

5. **Continue Until No Swaps Are Made**: Bubble Sort keeps repeating these
    passes through the list until no more swaps are made during a pass.
    When no swaps are made in a pass, the list is considered sorted, and
    the algorithm terminates.

6. **Time Complexity**: The time complexity of Bubble Sort is O(n^2) in
    the worst and average cases, where 'n' is the number of elements in
    the list. This is because, in the worst case, it has to make n-1 passes
    through the list, each requiring n-1 comparisons and possibly swaps.

Bubble Sort is straightforward to understand and implement, but it is not the most
efficient sorting algorithm, especially for large datasets. It is mainly used for
educational purposes and for situations where simplicity is more important than
sorting speed. There are more efficient sorting algorithms, such as Quick Sort and
Merge Sort, which are preferred for practical applications.
