# Radix Sort

Sorts integers by processing individual digits.
It can be efficient for sorting integers with a
limited range of values.


## Complexity
Time Complexity: **O(nk)** where k is the number of 
digits in the largest number.

## Description
Radix Sort is a non-comparative sorting algorithm that operates on integers, treating them
as strings of digits. It is suitable for sorting integers, especially when they have the
same number of digits. Radix Sort works by sorting the input numbers digit by digit, from
the least significant digit (rightmost) to the most significant digit (leftmost).

## How it works
1. **Digit Extraction:** Radix Sort starts by extracting the least significant digit
  (rightmost) from each number in the input array. It groups the numbers into "buckets" based
  on the value of this digit. Each bucket represents a range of values for that digit (e.g.,
  all numbers with a 0 in the rightmost digit go into one bucket, all numbers with a 1 go into
  another, and so on).

2. **Sorting by Digit:** After the numbers are grouped into buckets based on the least significant
  digit, they are rearranged in the original array, preserving their order within each bucket.
  This step effectively sorts the numbers based on the least significant digit.

3. **Repeat for All Digits:** The process is then repeated for the next least significant digit
  (moving left) and continues until all digits have been considered. Each time, the numbers are
  sorted based on the current digit, which ensures that they are sorted correctly based on all digits.

4. **Final Sorted Array:** After sorting by the most significant digit (leftmost), the array is
  fully sorted.

## Details
**Key points to note about Radix Sort:**

* It is a stable sorting algorithm, meaning that the relative order of equal elements is preserved.

* Radix Sort can be applied to both positive and negative integers. For negative numbers,
it may require extra steps to handle the sign bit and sorting the absolute values.

* The time complexity of Radix Sort is O(n * k), where n is the number of elements and k is
the number of digits in the maximum number. If k is small, Radix Sort can be very efficient.

* Radix Sort is often used for sorting fixed-length strings, where each character
position represents a digit.

* For Radix Sort to be efficient, the range of values for each digit (0-9 in base 10) should
be relatively small.

Radix Sort is a unique sorting algorithm that can be highly efficient for specific cases,
but it is not the best choice for all scenarios. It excels when the numbers to be sorted
have a limited number of digits and when there are a large number of numbers with the same
number of digits.
