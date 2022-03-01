package main

type Comparable struct {
	index int
	data  interface{}
}

type BinaryHeap struct {
	list []*Comparable
}

func NewBinaryHeap(list []*Comparable) *BinaryHeap {
	heap := &BinaryHeap{
		list: list,
	}

	heap.maxHeapify(list)
	heap.list = list
	return heap
}

func (bh *BinaryHeap) maxHeapify(list []*Comparable) {
	for i := len(list)/2 - 1; i >= 0; i-- {
		bh.bublingDown(i)
	}
}

func (bh *BinaryHeap) bublingDown(i int) {
	node := bh.list[i]
	childLeft := 2*i + 1
	childRight := 2*i + 2

	largest := i

	if childLeft < len(bh.list) && bh.list[childLeft].index > bh.list[largest].index {
		largest = childLeft
	}

	if childRight < len(bh.list) && bh.list[childRight].index > bh.list[largest].index {
		largest = childRight
	}

	if largest != i {
		bh.list[i] = bh.list[largest]
		bh.list[largest] = node
		bh.bublingDown(largest)
	}
}

func (bh *BinaryHeap) bublingUp(i int) {
	parent := (i - 1) / 2

	if bh.list[i].index > bh.list[parent].index {
		node := bh.list[parent]
		bh.list[parent] = bh.list[i]
		bh.list[i] = node
		bh.bublingUp(parent)
	}
}

func (bh *BinaryHeap) Add(item *Comparable) {
	bh.list = append(bh.list, item)
	bh.bublingUp(len(bh.list) - 1)
}

func (bh *BinaryHeap) Poll() *Comparable {
	last := len(bh.list) - 1
	node := bh.list[0]
	bh.list[0] = bh.list[last]
	bh.list = bh.list[:last]
	bh.bublingDown(0)
	return node
}

func (bh *BinaryHeap) Peek() *Comparable {
	return bh.list[0]
}
