package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func smallestK(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	if k == 0 {
		return nil
	}
	h := IntHeap(arr[:k])
	heap.Init(&h)
	for i := k; i < len(arr); i++ {
		if arr[i] < h[0] {
			heap.Pop(&h)
			heap.Push(&h, arr[i])
		}
	}
	return h
}
