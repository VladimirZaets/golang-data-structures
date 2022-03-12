package main

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	element01 := &Comparable{
		data:  "val-0",
		index: 50,
	}
	element02 := &Comparable{
		data:  "val-1",
		index: 33,
	}
	element03 := &Comparable{
		data:  "val-2",
		index: 11,
	}
	element04 := &Comparable{
		data:  "val-3",
		index: 22,
	}
	element05 := &Comparable{
		data:  "val-4",
		index: 66,
	}
	element06 := &Comparable{
		data:  "val-5",
		index: 12,
	}
	notExistsElement := &Comparable{
		data:  "none",
		index: 9,
	}
	bst.Add(element01)
	bst.Add(element02)
	bst.Add(element03)
	bst.Add(element04)
	bst.Add(element05)
	bst.Add(element06)

	contains := bst.Contains(notExistsElement)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", notExistsElement.index)
	}
	contains = bst.Contains(element05)
	if contains == false {
		t.Errorf("Element with index %d exists in tree, got does not exists", element05.index)
	}
	bst.Remove(element05)
	contains = bst.Contains(element05)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", element05.index)
	}
	bst.Remove(element01)
	bst.Remove(element03)

	contains = bst.Contains(element01)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", element01.index)
	}
	contains = bst.Contains(element03)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", element03.index)
	}
	if bst.size != 3 {
		t.Errorf("The size of tree is %d, should be %d", bst.size, 3)
	}
	fmt.Println("Hello")
}
