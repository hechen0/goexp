package fifo

import "testing"

func TestFifo_Queue(t *testing.T) {
	fifo := New(3)

	if fifo.Enqueue(1) != nil {
		t.Fatal("Enqueue should not return error")
	}
	if fifo.Enqueue(2) != nil {
		t.Fatal("Enqueue should not return error")
	}
	if fifo.Enqueue(3) != nil {
		t.Fatal("Enqueue should not return error")
	}

	err := fifo.Enqueue(4)
	if err == nil {
		t.Fatal("Enqueue should return error")
	}
	if err != ErrQueueFull {
		t.Fatal("Enqueue should return ErrQueueFull")
	}

	x, err := fifo.Dequeue()
	if err != nil {
		t.Fatal("Dequeue should not return error")
	}
	item, _ := x.(int)
	if item != 1 {
		t.Fatal("Dequeue should return 1")
	}
	fifo.Dequeue()
	fifo.Dequeue()

	_, err = fifo.Dequeue()
	if err == nil {
		t.Fatal("Dequeue should return error")
	}
	if err != ErrQueueEmpty {
		t.Fatal("Dequeue should return ErrQueueEmpty")
	}
}
