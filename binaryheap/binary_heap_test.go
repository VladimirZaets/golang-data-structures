package binaryheap

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/VladimirZaets/godatastructures/common/comparable"
)

func TestAddToMaxHeap(t *testing.T) {
	heap := createMaxHeap()
	isValid := isMaxHeap(heap, 0)
	if isValid == false {
		t.Error("Error, max heap validation failed")
	}
	for i := 0; i < 10; i++ {
		heap.Add(comparable.NewComparable(randInt(0, 100), fmt.Sprintf("Added value - %d", i)))
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
		t.Errorf("Peek must return top value from Max Heap, got %d", heap.Peek().GetIndex())
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
	if node.GetIndex() != res.GetIndex() {
		t.Errorf("Poll must return top value (%d) from Max Heap, got %d", node.GetIndex(), res.GetIndex())
	}
	if length-1 != len(heap.list) {
		t.Errorf("Poll must delete top value from Max Heap, length should %d, got %d", length-1, len(heap.list))
	}
}

func isMaxHeap(heap *BinaryHeap, i int) bool {
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	size := len(heap.list) - 1

	if leftChild <= size && heap.list[leftChild].GetIndex() > heap.list[i].GetIndex() {
		return false
	}

	if rightChild <= size && heap.list[rightChild].GetIndex() > heap.list[i].GetIndex() {
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
	list := make([]*comparable.Comparable, 20)
	for i := 0; i < 20; i++ {
		list[i] = comparable.NewComparable(randInt(0, 100), fmt.Sprintf("val-%d", i))
	}

	return NewBinaryHeap(list)
}
