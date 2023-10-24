# Basics
Basics

## Algorithms in Go:

### Sorting Algorithms
Go includes a standard library that provides sorting functions for slices.
You can also implement sorting algorithms like Bubble Sort, Quick Sort, and Merge Sort.

```go
import "sort"

// Using built-in sort
data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
sort.Ints(data)

// Implementing Quick Sort
func quickSort(arr []int) []int {
// Implementation
}
```

---
### Searching Algorithms
Common searching algorithms include linear search and binary search.
```go
// Linear Search
func linearSearch(arr []int, target int) int {
// Implementation
}

// Binary Search
func binarySearch(arr []int, target int) int {
// Implementation
}
```


---
### Graph Algorithms
Go provides powerful libraries for working with graphs.
Common graph algorithms include Depth-First Search (DFS) and Breadth-First Search (BFS).

```go
// Example of DFS
func dfs(graph Graph, startNode Node) {
// Implementation
}

// Example of BFS
func bfs(graph Graph, startNode Node) {
// Implementation
}
```


---
### Dynamic Programming:
Dynamic programming is a technique used to solve problems by
breaking them down into smaller sub-problems and caching
their solutions to avoid redundant computation.

```go
// Example of Fibonacci using dynamic programming
func fibonacci(n int) int {
// Implementation
}
```
