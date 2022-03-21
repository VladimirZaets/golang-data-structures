package doublelinkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func (n *Node) Get() interface{} {
	return n.data
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dll *DoubleLinkedList) AddByIndex(data interface{}, index int) error {
	if index >= dll.size {
		dll.AddToTail(data)
	}

	if index == 0 {
		dll.AddToHead(data)
	}

	node, err := findNodeByIndex(dll, index)
	if err != nil {
		return err
	}

	n := &Node{
		prev: node,
		next: node.next,
		data: data,
	}

	node.next.prev = n
	node.next = n
	dll.size++
	return nil
}

func (dll *DoubleLinkedList) AddToTail(data interface{}) error {
	if dll.head == nil {
		return dll.AddToHead(data)
	}

	if dll.tail == nil {
		dll.tail = &Node{
			data: data,
			prev: dll.head,
		}
		dll.head.next = dll.tail
		dll.size++
		return nil
	}

	node := &Node{
		prev: dll.tail,
		data: data,
	}

	dll.tail.next = node
	dll.tail = node
	dll.size++
	return nil
}

func (dll *DoubleLinkedList) AddToHead(data interface{}) error {
	if dll.head == nil {
		dll.head = &Node{
			data: data,
		}
		dll.size++
		return nil
	}

	node := &Node{
		next: dll.head,
		data: data,
	}

	if dll.tail == nil {
		dll.tail = dll.head
	}

	dll.head.prev = node
	dll.head = node

	dll.size++
	return nil
}

func (dll *DoubleLinkedList) GetFromTail() *Node {
	if dll.tail == nil {
		return dll.GetFromHead()
	}
	return dll.tail
}

func (dll *DoubleLinkedList) GetFromHead() *Node {
	return dll.head
}

func (dll *DoubleLinkedList) RemoveLast() error {
	if dll.tail == nil {
		return dll.RemoveFirst()
	}

	if dll.tail.prev == dll.head {
		dll.tail = nil
		dll.head.next = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}

	dll.size--
	return nil
}

func (dll *DoubleLinkedList) RemoveFirst() error {
	if dll.head == nil {
		return errors.New("double linked list is empty")
	}

	if dll.tail == nil {
		dll.head = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.size--
	return nil
}

func (dll *DoubleLinkedList) RemoveByIndex(index int) error {
	if index == dll.size {
		dll.RemoveLast()
	}

	if index == 0 {
		dll.RemoveFirst()
	}

	node, err := findNodeByIndex(dll, index)
	if err != nil {
		return err
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	node = nil
	dll.size--
	return nil
}

func (dll *DoubleLinkedList) Remove(n *Node) error {
	if n == dll.head {
		return dll.RemoveFirst()
	}
	if n == dll.tail {
		return dll.RemoveLast()
	}

	node := dll.head.next
	for node != nil {
		if node == n {
			node.prev.next = node.next
			node.next.prev = node.prev
			node = nil
			dll.size--
			return nil
		}
		node = node.next
	}

	return nil
}

func findNodeByIndex(dll *DoubleLinkedList, index int) (*Node, error) {
	var node *Node

	if index < dll.size/2 {
		i := dll.size - 1
		node = dll.tail.prev
		for index != i {
			i--
			node = node.prev
		}
	} else {
		i := 1
		node = dll.head.next
		for index == i {
			i++
			node = node.next
		}
	}

	if node == nil {
		return nil, fmt.Errorf("cannot find node with index %d", index)
	}

	return node, nil
}
