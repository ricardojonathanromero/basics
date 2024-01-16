package bfs

import (
	"container/list"
	"fmt"
	"strings"
)

type Node struct {
	Value     int
	Visited   bool
	Neighbors []int
}

type queue struct{ items []int }

func (q *queue) Dequeue()        { q.items = q.items[1:] }
func (q *queue) Enqueue(val int) { q.items = append(q.items, val) }
func (q *queue) IsEmpty() bool   { return len(q.items) == 0 }

func newQueue() *queue { return &queue{items: make([]int, 0)} }

func breadthFirstSearch(input map[int]Node, source, target int) string {
	var answer string
	path := make(map[int]int)

	_, ok := input[source]
	if !ok {
		return "not_found"
	}

	q := newQueue()
	q.Enqueue(source)

	bfs(input, q, path, target)

	if _, exists := path[target]; !exists {
		return "not_found"
	}

	parent := target
	answer = fmt.Sprintf("%d", target)
	for parent > 0 {
		val, parentExists := path[parent]
		parent = val
		if parentExists {
			answer = fmt.Sprintf("%d->%s", val, answer)
		}
	}

	return answer
}

func bfs(input map[int]Node, q *queue, path map[int]int, target int) {
	if q.IsEmpty() {
		return
	}

	idx := q.items[0] // get first element
	q.Dequeue()       // Remove first element

	node, exist := input[idx]
	if !exist || node.Visited || node.Value == target {
		return
	}

	node.Visited = true // mark node as visited
	input[idx] = node   // save node info

	// enqueue new items
	for _, neighbor := range node.Neighbors {
		path[neighbor] = node.Value
		q.Enqueue(neighbor)
	}

	bfs(input, q, path, target)
}

// optimized
func breadthFirstSearchOptimized(input map[int]Node, source, target int) string {
	// Check if the source and target nodes exist in the input
	sourceNode, sourceExists := input[source]
	_, targetExists := input[target]
	if !sourceExists || !targetExists {
		return "not_found"
	}

	// Create a queue for BFS and initialize it with the source node
	q := list.New()
	q.PushBack(sourceNode)

	// Create a map to store the parent of each node in the BFS traversal
	parent := make(map[int]int)
	parent[source] = -1

	// Perform BFS
	for q.Len() > 0 {
		current := q.Front().Value.(Node)
		q.Remove(q.Front())

		if current.Value == target {
			// Found the target node, reconstruct the path
			var path []int
			for current.Value != -1 {
				path = append([]int{current.Value}, path...)
				current.Value = parent[current.Value]
			}
			return formatPath(path)
		}

		for _, neighbor := range current.Neighbors {
			if _, visited := parent[neighbor]; !visited {
				parent[neighbor] = current.Value
				q.PushBack(input[neighbor])
			}
		}
	}

	return "not_found"
}

func formatPath(path []int) string {
	result := fmt.Sprintf("%d", path[0])
	for i := 1; i < len(path); i++ {
		result = fmt.Sprintf("%s->%d", result, path[i])
	}
	return result
}

type NodeStr struct {
	Value     string
	Neighbors []string
}

func breadthFirstSearchStr(tree map[string]NodeStr, root, target string) string {
	const notFound = "not_found" // default value

	// check if root and target exist in the tree
	rootNode, rootExists := tree[root]
	_, targetExists := tree[target]
	if !rootExists || !targetExists {
		return notFound
	}

	// initialize the queue and push the root node
	q := list.New()
	q.PushBack(rootNode)

	// create a parent map to save the interactions and recreate the path
	parents := make(map[string]string) // initialize queue
	parents[root] = ""                 // initialize root without any parents

	// while queue has elements, keep iterating
	for q.Len() > 0 {
		currentNode := q.Front().Value.(NodeStr) // get first element
		q.Remove(q.Front())                      // remove first element from queue

		// compare if node is equals to target
		if strings.EqualFold(currentNode.Value, target) {
			// the target has been looked
			// reconstructing the path
			var route []string
			for len(currentNode.Value) > 0 {
				// recreating route
				route = append([]string{currentNode.Value}, route...)
				// changing pointer
				currentNode.Value = parents[currentNode.Value]
			}

			// returning path result
			return strings.Join(route, "->")
		}

		// iterate neighbors
		for _, neighbor := range currentNode.Neighbors {
			// check if the neighbor has not already been visited
			if _, visited := parents[neighbor]; !visited {
				parents[neighbor] = currentNode.Value // add neighbor to parents map associated to current node value
				q.PushBack(tree[neighbor])            // add neighbor to the end of the queue
			}
		}
	}

	return notFound
}
