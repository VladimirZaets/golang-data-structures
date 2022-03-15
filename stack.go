package main

import (
	"errors"
)

type Stack struct {
	list *DoubleLinkedList
	size int
}

func NewStack() *Stack {
	return &Stack{
		list: NewDoubleLinkedList(),
	}
}

func (stk *Stack) Push(data interface{}) error {
	err := stk.list.AddToTail(data)
	if err != nil {
		return err
	}
	stk.size++
	return nil
}

func (stk *Stack) Pop() error {
	if stk.size == 0 {
		return errors.New("trying pop element from empty stack")
	}

	err := stk.list.RemoveLast()
	if err != nil {
		return err
	}
	stk.size--
	return nil
}

func (stk *Stack) Peek() interface{} {
	if stk.size == 0 {
		return nil
	}

	if stk.size == 1 {
		return stk.list.head.data
	}

	return stk.list.tail.data
}
