# Bucket Sort

Distributes elements into a number of buckets, each of
which is then sorted, either using a different sorting
algorithm or by recursively applying the bucket sort
algorithm.


## Complexity
Time Complexity: O(n^2) in the worst case, but often
linear O(n) when elements are uniformly distributed.


## Description
Bucket Sort is a sorting algorithm that is mainly useful
for sorting elements that are uniformly distributed over a
range. It is a distribution-based sorting algorithm that
works by distributing the elements into a number of buckets
and then sorting each bucket individually. The final sorted
list is obtained by concatenating the elements from all
the buckets.

### How this algorithm works
1. **Determine the Range**: First, you need to determine the range
   of the input elements. Find the minimum and maximum values in
   the input array.

2. **Create Buckets**: Divide the range of values into a fixed number
   of equal-sized buckets. The number of buckets is typically chosen
   based on the nature of the data and can vary.

3. **Distribute Elements**: Iterate through the input array, placing
   each element into the appropriate bucket. The formula for determining
   the bucket for an element depends on the range and number of buckets.
   It could be something like bucket_index = (value - min) / bucket_size,
   where min is the minimum value, and bucket_size is the size of each bucket.

4. **Sort Each Bucket**: Sort each bucket individually. This can be done
   using another sorting algorithm like Insertion Sort, Quick Sort, or any
   other suitable sorting algorithm. The choice of sorting algorithm depends
   on the specific implementation.

5. **Concatenate Buckets**: After all the buckets are sorted, concatenate
   them to obtain the final sorted array.
