package queue

import "testing"

func TestQSize(t *testing.T) {
	queue := NewQueue(10)
	if queue.Size() != 0 {
		t.Error("The size should be 0")
	}
	queue.Offer("val-1")
	queue.Offer("val-1")
	if queue.Size() != 2 {
		t.Error("The size should be 2")
	}
}

func TestQIsEmpty(t *testing.T) {
	queue := NewQueue(10)
	if !queue.IsEmpty() {
		t.Error("The queue should be empty")
	}
	queue.Offer("val-1")
	queue.Offer("val-2")
	if queue.IsEmpty() {
		t.Error("The queue should contain 2 element")
	}
}

func TestQPeek(t *testing.T) {
	queue := NewQueue(10)
	data, err := queue.Peek()
	if err == nil {
		t.Errorf("The Peek operation on empty queue must return error, got %s", data)
	}
	queue.Offer("val-1")
	queue.Offer("val-2")
	if val, _ := queue.Peek(); val != "val-1" {
		t.Errorf("The value should be 'val-1', got %s", val)
	}
}

func TestQPoll(t *testing.T) {
	queue := NewQueue(10)
	err := queue.Poll()
	if err == nil {
		t.Error("The Pool operation on empty queue must return error")
	}
	queue.Offer("val-1")
	queue.Offer("val-2")
	queue.Offer("val-3")
	queue.Offer("val-4")

	queue.Poll()
	if val, _ := queue.Peek(); val != "val-2" {
		t.Errorf("The value should be 'val-2', got %s", val)
	}
	if queue.Size() != 3 {
		t.Errorf("The size of queue should be 3, got %d", queue.Size())
	}
	if queue.frontI != 1 {
		t.Errorf("The front pointer of queue should be 1, got %d", queue.frontI)
	}

	if queue.backI != 4 {
		t.Errorf("The back pointer of queue should be 4, got %d", queue.frontI)
	}
}

func TestQffer(t *testing.T) {
	queue := NewQueue(5)
	queue.Offer("val-1")
	queue.Offer("val-2")
	queue.Offer("val-3")
	queue.Offer("val-4")
	queue.Offer("val-5")
	if queue.Size() != 5 {
		t.Errorf("The size of queue should be 5, got %d", queue.Size())
	}
	if queue.backI != 5 {
		t.Errorf("The back pointer of queue should be 5, got %d", queue.frontI)
	}
	if queue.frontI != 0 {
		t.Errorf("The front pointer of queue should be 0, got %d", queue.frontI)
	}
	err := queue.Offer("val-6")
	if err == nil {
		t.Errorf("Trying to add more elements than queue capacity. Should be an error")
	}

	queue.Poll()
	queue.Offer("val-6")
	if queue.backI != 1 {
		t.Errorf("The back pointer of queue should be 1, got %d", queue.frontI)
	}
	queue.Poll()
	queue.Offer("val-7")
	if queue.backI != 2 {
		t.Errorf("The back pointer of queue should be 1, got %d", queue.frontI)
	}
}
