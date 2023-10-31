package Assignment

import "container/heap"

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Pair struct {
	node   int
	weight int
}

func dijkstra(A int, B [][]int, C int) []int {

	// C is the source node
	// B is the edge representation and weights
	// A is the nodes (0 to A-1)

	// Q: find shortest path from C to each of the other nodes

	// Approach
	// create adjacency list (should store node and weight)

	graph := adjacencyList(A, B)

	// Apply Dijkstra to find the shortest path from any current node
	// push nodes to min heap based on their weight

	hp := &MinHeap{}
	heap.Init(hp)

	// mark initial distances
	distance := make([]int, A)
	for i := 0; i < A; i++ {
		distance[i] = 100000000
	}

	distance[C] = 0
	heap.Push(hp, Pair{
		node:   C,
		weight: 0,
	})

	for hp.Len() > 0 {
		pair := heap.Pop(hp).(Pair)

		for i := 0; i < len(graph[pair.node]); i++ {
			p := graph[pair.node][i]
			//distance[p.node] = min(distance[p.node], distance[pair.node]+p.weight)
			if distance[p.node] > distance[pair.node]+p.weight {
				distance[p.node] = distance[pair.node] + p.weight
				heap.Push(hp, p)
			}
		}
	}

	for i := range distance {

		if distance[i] == 100000000 {
			distance[i] = -1
		}
	}

	return distance
}

func adjacencyList(A int, B [][]int) [][]Pair {
	graph := make([][]Pair, A)
	for i := range B {

		src := B[i][0]
		dest := B[i][1]
		weight := B[i][2]

		graph[src] = append(graph[src], Pair{dest, weight})
		graph[dest] = append(graph[dest], Pair{src, weight})

	}
	return graph
}
