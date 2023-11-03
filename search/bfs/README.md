# Breadth-First Search
A graph traversal algorithm that explores all the vertices of a graph level by level. It can be used for searching in graphs or trees.

## Description
Breadth-First Search (BFS) is a graph traversal algorithm that explores a graph or tree in a level-by-level fashion. It starts at a given source node and systematically explores all its neighboring nodes before moving on to their neighbors. BFS is commonly used to find the shortest path between two nodes in an unweighted graph and to explore the structure of a graph.

## How it works
1. **Begin at the source node:** You start the BFS algorithm at a specified source node.

2. **Initialize a queue:** Create a queue (FIFO - First-In-First-Out) data structure to keep track of the nodes to be visited. Place the source node in the queue.

3. **Mark the source node as visited:** To avoid revisiting the same node, mark the source node as visited.

4. **Explore neighbors:** While the queue is not empty, repeat the following steps:

   4.1 Dequeue a node: Remove the node at the front of the queue. This node will be the current node for exploration. 
   
   4.2 Visit the node: Process the current node, which may involve recording or displaying information about it. 
   
   4.3 Enqueue unvisited neighbors: Visit all the neighbors of the current node that have not been visited yet. Mark each of them as visited, and enqueue them in the queue. This ensures that you explore all neighbors at the current level before moving on to the next level.

5. **Repeat until the queue is empty:** Continue the process until the queue is empty. This guarantees that you visit all nodes in the graph reachable from the source node, and you do so in a breadth-first order.

6. **Terminate:** The BFS algorithm terminates when all nodes have been visited, or when you have found the target node (if you are searching for a specific node).

---
BFS is an effective way to explore a graph systematically while finding the shortest path in an unweighted graph. It's a valuable algorithm in various applications, such as web crawling, social network analysis, and solving puzzles like the 8-puzzle or maze solving.

---
## Problem
You are given an undirected graph represented as an adjacency list, and you need to find the shortest path between two specified nodes in the graph. The graph is unweighted, which means all edges have the same weight (distance). Implement a BFS algorithm to solve this problem.

Input:

The graph as an adjacency list.
Two nodes, a source node, and a target node.
Output:

The shortest path from the source node to the target node, represented as a list of nodes in the order of the path.
For example, consider the following graph:

```lua
Graph:
1 -- 2 -- 3
|         |
4 -- 5 -- 6
```

If you want to find the shortest path from node 1 to node 6, the expected output would be [1, 2, 3, 6].

To solve this problem, you can implement the BFS algorithm as described earlier. Start from the source node, perform a BFS traversal of the graph, and keep track of the path as you go. When you reach the target node, you'll have the shortest path from the source to the target.

This problem is a common application of BFS and is used in various scenarios, such as finding the shortest route in a maze, navigating a network of interconnected locations, or determining the fewest steps required to solve a puzzle.
