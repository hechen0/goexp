package priority_queue

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Interface){
	n := h.Len()

	for i:= 0;i<n;i++{
		down(h, i)
	}
}

func down(h Interface, i int){
	n := h.Len()
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}

		j := j1

		if j2 := j1+1; j2 < n && h.Less(j2, j1) {

		}
	}
}
