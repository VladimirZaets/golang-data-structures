package binarysearchtree

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/VladimirZaets/godatastructures/common/comparable"
)

func TestBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	element01 := comparable.NewComparable(50, "val-0")
	element02 := comparable.NewComparable(33, "val-1")
	element03 := comparable.NewComparable(11, "val-2")
	element04 := comparable.NewComparable(22, "val-3")
	element05 := comparable.NewComparable(66, "val-4")
	element06 := comparable.NewComparable(12, "val-5")
	notExistsElement := comparable.NewComparable(9, "none")
	bst.Add(element01)
	bst.Add(element02)
	bst.Add(element03)
	bst.Add(element04)
	bst.Add(element05)
	bst.Add(element06)

	contains := bst.Contains(notExistsElement)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", notExistsElement.GetIndex())
	}
	contains = bst.Contains(element05)
	if contains == false {
		t.Errorf("Element with index %d exists in tree, got does not exists", element05.GetIndex())
	}
	bst.Remove(element05)
	contains = bst.Contains(element05)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", element05.GetIndex())
	}
	bst.Remove(element01)
	bst.Remove(element03)

	contains = bst.Contains(element01)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", element01.GetIndex())
	}
	contains = bst.Contains(element03)
	if contains == true {
		t.Errorf("Element with index %d does not exists in tree, got exists", element03.GetIndex())
	}
	if bst.size != 3 {
		t.Errorf("The size of tree is %d, should be %d", bst.size, 3)
	}
}

func TestBinarySearchTreeHeight(t *testing.T) {
	bst := NewBinarySearchTree()
	for i, index := range []int{30, 55, 15, 14, 20, 25, 60, 50, 70, 61, 81, 99} {
		bst.Add(comparable.NewComparable(index, fmt.Sprintf("val-%d", i)))
	}
	height := bst.Height()
	if height != 6 {
		t.Errorf("The height should be %d, got %d", 6, height)
	}
}

func TestBinarySearchTreeTraverse(t *testing.T) {
	bst := NewBinarySearchTree()
	for i, index := range []int{30, 55, 15, 14, 20, 25, 60, 50, 70, 61, 81, 99} {
		bst.Add(comparable.NewComparable(index, fmt.Sprintf("val-%d", i)))
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

func injectIndexes(sl []*comparable.Comparable) []int {
	result := make([]int, len(sl))

	for i, val := range sl {
		result[i] = val.GetIndex()
	}

	return result
}
