# Binary Search
Applicable to sorted lists, it compares the target value to the middle element of the list and eliminates half of the remaining elements in each iteration. This is a highly efficient search algorithm for sorted data.

## Description
Binary search is an efficient algorithm for finding a specific element within a sorted array or list. It operates by repeatedly dividing the search interval in half until the target element is found or it is determined that the element is not present in the array.

## How it works
1. **Initialize:** Begin with the entire sorted array. Set two pointers, low and high, to the first and last indices of the array, respectively.

2. **Compare:** Calculate the middle index as (low + high) / 2. Compare the middle element with the target element you're searching for.

3. **Target Found:** If the middle element is equal to the target element, you have found the target, and you can return its index.

4. **Narrow the Search:** If the middle element is less than the target element, it means the target must be in the right half of the array. Update low to middle + 1, effectively discarding the left half of the array.

5. **Narrow the Search (Left):** If the middle element is greater than the target element, the target must be in the left half of the array. Update high to middle - 1, effectively discarding the right half of the array.

6. **Repeat:** Continue steps 2-5 until the low pointer is greater than the high pointer. If you reach this point without finding the target, it means the target is not in the array.

Binary search is significantly faster than linear search for large datasets because it reduces the search space by half with each iteration. It has a time complexity of O(log n), which makes it a very efficient search algorithm for sorted data. However, it requires the data to be sorted, and it may not be the best choice if you frequently need to insert or delete elements in the array, as maintaining the sorted order can be costly.
