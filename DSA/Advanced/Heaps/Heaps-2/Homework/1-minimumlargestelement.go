package Homework

import "container/heap"

//Given an array A of N numbers, you have to perform B operations. In each operation, you have to pick any one of the N elements and add the original value(value stored at the index before we did any operations) to its current value. You can choose any of the N elements in each operation.
//
//Perform B operations in such a way that the largest element of the modified array(after B operations) is minimized.
//Find the minimum possible largest element after B operations.

type Heap []State

type State struct {
	curr   int
	init   int
	future int
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].future < h[j].future
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(State))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumLargestElement(A []int, B int) int {

	// use a min heap that stores 3 params (curr, initial, future)
	// min heap works on future

	// in every iteration, do following updates
	// 	curr = future
	// 	future = curr + init

	hp := &Heap{}
	heap.Init(hp)

	for i := range A {
		curr := A[i]
		init := A[i]
		future := curr + init
		heap.Push(hp, State{curr, init, future})
	}

	for i := 0; i < B; i++ {
		state := heap.Pop(hp).(State)
		heap.Push(hp, State{state.future, state.init, state.future + state.init})
	}

	max := -10000000
	for len(*hp) > 0 {
		x := heap.Pop(hp).(State).curr
		if x > max {
			max = x
		}
	}

	return max
}
