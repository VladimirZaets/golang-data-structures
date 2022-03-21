package stack

import (
	"errors"

	"github.com/VladimirZaets/godatastructures/doublelinkedlist"
)

type Stack struct {
	list *doublelinkedlist.DoubleLinkedList
	size int
}

func NewStack() *Stack {
	return &Stack{
		list: doublelinkedlist.NewDoubleLinkedList(),
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
	node := stk.list.GetFromTail()
	if node == nil {
		return nil
	}
	return node.Get()
}
