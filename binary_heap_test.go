package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAddToMaxHeap(t *testing.T) {
	heap := createMaxHeap()
	isValid := isMaxHeap(heap, 0)
	if isValid == false {
		t.Error("Error, max heap validation failed")
	}
	for i := 0; i < 10; i++ {
		heap.Add(&Comparable{
			index: randInt(0, 100),
			data:  fmt.Sprintf("Added value - %d", i),
		})
	}
	isValid = isMaxHeap(heap, 0)
	if isValid == false {
		t.Error("Error, max heap validation failed after items were added")
	}
}

func TestPeekFromMaxHeap(t *testing.T) {
	heap := createMaxHeap()
	isValid := isMaxHeap(heap, 0)
	if isValid == false {
		t.Error("Error, max heap validation failed")
	}
	if heap.list[0] != heap.Peek() {
		t.Errorf("Peek must return top value from Max Heap, got %d", heap.Peek().index)
	}

}

func TestPollFromMaxHeap(t *testing.T) {
	heap := createMaxHeap()
	isValid := isMaxHeap(heap, 0)
	length := len(heap.list)
	if isValid == false {
		t.Error("Error, max heap validation failed")
	}
	node := heap.Peek()
	res := heap.Poll()
	isValid = isMaxHeap(heap, 0)
	if isValid == false {
		t.Error("Error, max heap validation failed after Poll from heap")
	}
	if node.index != res.index {
		t.Errorf("Poll must return top value (%d) from Max Heap, got %d", node.index, res.index)
	}
	if length-1 != len(heap.list) {
		t.Errorf("Poll must delete top value from Max Heap, length should %d, got %d", length-1, len(heap.list))
	}
}

func isMaxHeap(heap *BinaryHeap, i int) bool {
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	size := len(heap.list) - 1

	if leftChild <= size && heap.list[leftChild].index > heap.list[i].index {
		return false
	}

	if rightChild <= size && heap.list[rightChild].index > heap.list[i].index {
		return false
	}

	if rightChild > size || leftChild > size {
		return true
	}

	return isMaxHeap(heap, leftChild) && isMaxHeap(heap, rightChild)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func createMaxHeap() *BinaryHeap {
	list := make([]*Comparable, 20)
	for i := 0; i < 20; i++ {
		list[i] = &Comparable{
			index: randInt(0, 100),
			data:  fmt.Sprintf("val-%d", i),
		}
	}

	return NewBinaryHeap(list)
}
