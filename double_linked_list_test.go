package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddToTail(t *testing.T) {
	iteration := 5
	dll, err := generateBaseDLL(iteration)
	if err != nil {
		t.Error(err)
	}

	if dll.size != iteration {
		t.Errorf("%d nods were add to DLL, got %d", iteration, dll.size)
	}

	node := dll.head
	for i := 0; i < iteration; i++ {
		if node.data != fmt.Sprintf("val-%d", i) {
			t.Errorf("The node data %s must be equal %s", node.data, fmt.Sprintf("val-%d", i))
		}
		node = node.next
	}
}

func TestAddToHead(t *testing.T) {
	dll := NewDoubleLinkedList()
	iteration := 5

	for i := 0; i < iteration; i++ {
		err := dll.AddToHead(fmt.Sprintf("val-%d", i))
		if err != nil {
			t.Error(err)
		}
	}

	if dll.size != iteration {
		t.Errorf("%d nods were add to DLL, got %d", iteration, dll.size)
	}

	node := dll.tail
	for i := 0; i < iteration; i++ {
		if node.data != fmt.Sprintf("val-%d", i) {
			t.Errorf("The node data %s must be equal %s", node.data, fmt.Sprintf("val-%d", i))
		}
		node = node.prev
	}
}

func TestAddByIndex(t *testing.T) {
	dll, err := generateBaseDLL(5)
	if err != nil {
		t.Error(err)
	}
	val := "index-2-inserted"
	err = dll.AddByIndex(val, 2)
	if err != nil {
		t.Error(err)
	}

	if dll.head.next.next.data != val {
		t.Errorf("The item with data '%s' was should be added by index 2, got %s", val, dll.head.next.next.data)
	}
}

func TestAddByIndexOutOfDLL(t *testing.T) {
	dll := NewDoubleLinkedList()
	dll.AddByIndex("value-0", 0)
	dll.AddByIndex("value-5", 5)
	if dll.tail.data != "value-5" {
		t.Error("The out of scope item should be added to the end of linked list")
	}
}

func TestRemoveLast(t *testing.T) {
	dll, err := generateBaseDLL(5)
	if err != nil {
		t.Error(err)
	}

	dll.RemoveLast()
	if dll.size != 4 {
		t.Errorf("The size of DDL should be 4, current size is %s", strconv.Itoa(dll.size))
	}

	if dll.tail.data != "val-3" {
		t.Errorf("The tail node data should be 'val-3', got %s", dll.tail.data)
	}
}

func TestRemoveFirst(t *testing.T) {
	dll, err := generateBaseDLL(5)
	if err != nil {
		t.Error(err)
	}

	dll.RemoveFirst()
	if dll.size != 4 {
		t.Errorf("The size of DDL should be 4, current size is %s", strconv.Itoa(dll.size))
	}

	if dll.head.data != "val-1" {
		t.Errorf("The head node data should be 'val-3', got %s", dll.head.data)
	}
}

func TestRemoveByIndex(t *testing.T) {
	dll, err := generateBaseDLL(5)
	if err != nil {
		t.Error(err)
	}
	dll.RemoveByIndex(2)
	if dll.size != 4 {
		t.Errorf("The size of DDL should be 4, current size is %s", strconv.Itoa(dll.size))
	}

	if dll.head.next.data != "val-2" {
		t.Errorf("The 'dll.head.next.data' should be 'val-2', got %s", dll.head.next.data)
	}

}

func TestRemove(t *testing.T) {
	dll, err := generateBaseDLL(5)
	if err != nil {
		t.Error(err)
	}
	node := dll.head.next
	dll.Remove(node)

	if dll.size != 4 {
		t.Errorf("The size of DDL should be 4, current size is %s", strconv.Itoa(dll.size))
	}

	if node == dll.head.next {
		t.Error("The node should be removed")
	}
}

func generateBaseDLL(size int) (*DoubleLinkedList, error) {
	dll := NewDoubleLinkedList()
	iteration := 5

	for i := 0; i < iteration; i++ {
		err := dll.AddToTail(fmt.Sprintf("val-%d", i))
		if err != nil {
			return nil, err
		}
	}

	return dll, nil
}
