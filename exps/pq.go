package exps

import (
	"sort"
	"fmt"
)

type PQ interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
	Get(int) interface{}
}

func Init(h PQ) {
	// heapify

	n := h.Len()

	for i := n/2 -1 ; i >= 0; i-- {
		down(h,i,n)
	}
}

func Push(h PQ, x interface{}){
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h PQ) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove(h PQ, i int) interface{} {
	n := h.Len() -1

	if n != i {
		h.Swap(i, n)
		if !down(h, i, n){
			up(h, i)
		}
	}

	return h.Pop()
}

func up(h PQ, j int){
	for {
		i := (j-1)/2
		if i ==j || !h.Less(j, i){
			break
		}
		h.Swap(i,j)
		j = i
	}
}

func down(h PQ, i0, n int) bool {
	i := i0

	for {
		j1 := 2 * i + 1
		if j1 >= n || j1 < 0 {
			break
		}

		j := j1

		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2
		}

		if !h.Less(j, i) {
			break
		}

		fmt.Printf("swap [%v]: %v, [%v]: %v\n", j, h.Get(j), i, h.Get(i))
		h.Swap(j, i)
		i = j
	}
	return i > i0
}