package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	list   []interface{}
	size   int
	frontI int
	backI  int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		list:   make([]interface{}, cap),
		size:   0,
		frontI: 0,
		backI:  0,
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peek() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("call peek on empty queue")
	}
	return q.list[q.frontI], nil
}

func (q *Queue) Poll() error {
	if q.IsEmpty() {
		return errors.New("Queue is empty, cannot poll")
	}

	q.list[q.frontI] = nil
	q.frontI++
	q.size--
	return nil
}

func (q *Queue) Offer(data interface{}) error {
	if q.size+1 > cap(q.list) {
		return fmt.Errorf("trying to add more elements that Queue capacity. The capacity of Queue is %d", cap(q.list))
	}

	if q.backI+1 > cap(q.list) {
		q.backI = 0
	}

	q.list[q.backI] = data
	q.backI++
	q.size++
	return nil
}
