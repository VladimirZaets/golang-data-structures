package main

import (
	"fmt"
	"reflect"
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
}

func TestBinarySearchTreeHeight(t *testing.T) {
	bst := NewBinarySearchTree()
	for i, index := range []int{30, 55, 15, 14, 20, 25, 60, 50, 70, 61, 81, 99} {
		bst.Add(&Comparable{
			data:  fmt.Sprintf("val-%d", i),
			index: index,
		})
	}
	height := bst.Height()
	if height != 6 {
		t.Errorf("The height should be %d, got %d", 6, height)
	}
}

func TestBinarySearchTreeTraverse(t *testing.T) {
	bst := NewBinarySearchTree()
	for i, index := range []int{30, 55, 15, 14, 20, 25, 60, 50, 70, 61, 81, 99} {
		bst.Add(&Comparable{
			data:  fmt.Sprintf("val-%d", i),
			index: index,
		})
	}

	sl, _ := bst.Traverse(PreOrder)
	traverseExpectedIndexes := []int{30, 15, 14, 20, 25, 55, 50, 60, 70, 61, 81, 99}
	traverseResultIndexes := injectIndexes(sl)
	if !reflect.DeepEqual(traverseExpectedIndexes, traverseResultIndexes) {
		t.Errorf("Expected traverse not equal result. Expected %v, got %v", traverseExpectedIndexes, traverseResultIndexes)
	}
	sl, _ = bst.Traverse(PostOrder)
	traverseExpectedIndexes = []int{14, 25, 20, 15, 50, 61, 99, 81, 70, 60, 55, 30}
	traverseResultIndexes = injectIndexes(sl)
	if !reflect.DeepEqual(traverseExpectedIndexes, traverseResultIndexes) {
		t.Errorf("Expected traverse not equal result. Expected %v, got %v", traverseExpectedIndexes, traverseResultIndexes)
	}
	sl, _ = bst.Traverse(InOrder)
	traverseExpectedIndexes = []int{14, 15, 20, 25, 30, 50, 55, 60, 61, 70, 81, 99}
	traverseResultIndexes = injectIndexes(sl)
	if !reflect.DeepEqual(traverseExpectedIndexes, traverseResultIndexes) {
		t.Errorf("Expected traverse not equal result. Expected %v, got %v", traverseExpectedIndexes, traverseResultIndexes)
	}
	sl, _ = bst.Traverse(LevelOrder)
	traverseExpectedIndexes = []int{30, 15, 55, 14, 20, 50, 60, 25, 70, 61, 81, 99}
	traverseResultIndexes = injectIndexes(sl)
	if !reflect.DeepEqual(traverseExpectedIndexes, traverseResultIndexes) {
		t.Errorf("Expected traverse not equal result. Expected %v, got %v", traverseExpectedIndexes, traverseResultIndexes)
	}
}

func injectIndexes(sl []*Comparable) []int {
	result := make([]int, len(sl))

	for i, val := range sl {
		result[i] = val.index
	}

	return result
}
