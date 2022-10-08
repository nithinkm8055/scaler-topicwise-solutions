package Assignment

import (
	"container/heap"
	"sort"
)

type pair struct {
	time   int
	profit int
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FreeCars(A []int, B []int) int {
	var pairs []pair
	for i := range A {
		pairs = append(pairs, pair{A[i], B[i]})
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].time < pairs[j].time
	})

	day := 0
	hp := &Heap{}
	heap.Init(hp)
	sum := 0
	for i := 0; i < len(pairs); i++ {

		if pairs[i].time > day {
			day++
			heap.Push(hp, pairs[i].profit)
			sum += pairs[i].profit

		} else {
			min_ele := heap.Pop(hp).(int)
			if pairs[i].profit > min_ele {
				heap.Push(hp, pairs[i].profit)
				sum -= min_ele
				sum += pairs[i].profit
			} else {
				heap.Push(hp, min_ele)
			}
		}
	}

	return sum % (1e9 + 7)
}
