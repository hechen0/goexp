package leetcode

import (
	"container/heap"
	"bytes"
)

type node struct {
	value byte
	count int
	index int
}

type sortQueue []*node

func (s sortQueue) Len() int {return len(s)}
func (s sortQueue) Less(i, j int) bool { return s[i].count > s[j].count }
func (s sortQueue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].index = i
	s[j].index = j
}

func (s *sortQueue) Push(x interface{}){
	n := len(*s)
	item := x.(*node)
	item.index = n
	*s = append(*s, item)
}

func (s *sortQueue) Pop() interface{}{
	old := *s
	n := len(old)
	item := old[n-1]
	item.index = -1
	*s = old[0 : n-1]
	return item
}

func (s *sortQueue) update(item *node){
	item.count = item.count + 1
	heap.Fix(s, item.index)
}

func FrequencySort(s string) string {

	t := make(map[byte]int)
	for i := 0; i < len(s); i++{
		t[s[i]] = 1
	}

	queue := make(sortQueue, len(t))

	table := make(map[byte]*node)

	index := 0

	for i := 0; i < len(s); i++{
		v := s[i]

		if table[v] == nil {
			table[v] = &node{
				value: v,
				count: 0,
				index: index,
			}

			queue[index] = table[v]

			index++
		}
	}

	heap.Init(&queue)

	for i := 0; i < len(s); i++{
		v := s[i]

		if table[v] != nil {
			queue.update(table[v])
		}
	}

	ns := bytes.NewBuffer(nil)


	for queue.Len() > 0 {
		i := heap.Pop(&queue).(*node)

		for j := 0; j < i.count; j++{
			ns.WriteByte(i.value)
		}
	}

	return ns.String()
}
