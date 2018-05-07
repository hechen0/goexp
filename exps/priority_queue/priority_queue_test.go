package priority_queue

import (
	"testing"
	"container/heap"
	"fmt"
)

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"A": 3, "B": 2, "C": 1,
	}

	pq := make(PQ, len(items))
	i := 0
	for name, weight := range items {
		pq[i] = &Ele{name, weight}
		i++
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		ele := pq.Pop().(*Ele)
		fmt.Printf("%+v\n", ele)
	}
}
