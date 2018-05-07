package priority_queue

type Ele struct {
	Name   string
	Weight int
}

type PQ []*Ele

func (pq *PQ) Less(i, j int) bool {
	return (*pq)[i].Weight < (*pq)[j].Weight
}

func (pq *PQ) Len() int { return len(*pq) }

func (pq *PQ) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

func (pq *PQ) Push(x interface{}) {
	ele := x.(*Ele)
	*pq = append(*pq, ele)
}

func (pq *PQ) Pop() interface{} {
	ele := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return ele
}
