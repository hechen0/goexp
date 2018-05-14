package fifo

import "errors"

var (
	ErrQueueFull  = errors.New("queue full")
	ErrQueueEmpty = errors.New("queue empty")
)

type Interface interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
}

type Fifo struct {
	a                []interface{}
	head, tail, size int64
}

func New(size int64) *Fifo {
	return &Fifo{
		make([]interface{}, size), 0, 0, size,
	}
}

func (q *Fifo) Enqueue(x interface{}) error {
	// check if queue full
	if q.head == q.tail && q.a[q.tail] != nil {
		return ErrQueueFull
	}

	q.a[q.tail] = x
	q.tail = (q.tail + 1) % q.size
	return nil
}

func (q *Fifo) Dequeue() (interface{}, error) {
	// check if queue empty
	if q.head == q.tail && q.a[q.head] == nil {
		return nil, ErrQueueEmpty
	}

	x := q.a[q.head]
	q.a[q.head] = nil
	q.head = (q.head + 1) % q.size
	return x, nil
}
