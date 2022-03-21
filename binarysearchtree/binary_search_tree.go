package binarysearchtree

import (
	"fmt"

	"github.com/VladimirZaets/godatastructures/common/comparable"
	"github.com/VladimirZaets/godatastructures/queue"
	"github.com/VladimirZaets/godatastructures/stack"
)

const (
	driftLeft = iota
	driftRight
)

const (
	PreOrder = iota
	InOrder
	PostOrder
	LevelOrder
)

type BinarySearchTreeNode struct {
	data  *comparable.Comparable
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
	size int
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		size: 0,
	}
}

func NewBinarySearchTreeNode(data *comparable.Comparable, left *BinarySearchTreeNode, right *BinarySearchTreeNode) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{
		data:  data,
		left:  left,
		right: right,
	}
}

func (bst *BinarySearchTree) Add(element *comparable.Comparable) error {
	if bst.Contains(element) {
		return fmt.Errorf("element with index %d already exists in tree", element.GetIndex())
	}
	node := NewBinarySearchTreeNode(element, nil, nil)
	bst.size++
	bst.root = bst.add(bst.root, node)
	return nil
}

func (bst *BinarySearchTree) add(root *BinarySearchTreeNode, node *BinarySearchTreeNode) *BinarySearchTreeNode {
	if root == nil {
		return node
	}

	if root.data.GetIndex() > node.data.GetIndex() {
		root.left = bst.add(root.left, node)
	}

	if root.data.GetIndex() < node.data.GetIndex() {
		root.right = bst.add(root.right, node)
	}
	return root
}

func (bst *BinarySearchTree) Contains(element *comparable.Comparable) bool {
	_, err := bst.Find(element.GetIndex())
	return err == nil
}

func (bst *BinarySearchTree) Find(index int) (*comparable.Comparable, error) {
	node := bst.find(bst.root, index)
	if node == nil {
		return nil, fmt.Errorf("cannot find element with index %d", index)
	}
	return node.data, nil
}

func (bst *BinarySearchTree) find(node *BinarySearchTreeNode, index int) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}
	if node.data.GetIndex() == index {
		return node
	}
	if node.data.GetIndex() > index {
		return bst.find(node.left, index)
	}
	if node.data.GetIndex() < index {
		return bst.find(node.right, index)
	}
	return nil
}

func (bst *BinarySearchTree) Remove(element *comparable.Comparable) error {
	if !bst.Contains(element) {
		return fmt.Errorf("element with index %d doesn't exists in tree", element.GetIndex())
	}
	bst.root = bst.remove(bst.root, element.GetIndex())
	bst.size--
	return nil
}

func (bst *BinarySearchTree) remove(node *BinarySearchTreeNode, index int) *BinarySearchTreeNode {
	if node.data.GetIndex() > index {
		node.left = bst.remove(node.left, index)
	}
	if node.data.GetIndex() < index {
		node.right = bst.remove(node.right, index)
	}
	if node.data.GetIndex() == index {
		if node.left == nil {
			right := node.right
			node = nil
			return right
		} else if node.right == nil {
			left := node.left
			node = nil
			return left
		} else {
			temporaryNode := bst.drift(driftRight, node.left)
			temporaryNode.left = bst.remove(node.left, temporaryNode.data.GetIndex())
			temporaryNode.right = node.right
			return temporaryNode
		}
	}
	return node
}

func (bst *BinarySearchTree) drift(side int, node *BinarySearchTreeNode) *BinarySearchTreeNode {
	temporaryNode := node
	if side == driftRight {
		for temporaryNode.right != nil {
			temporaryNode = temporaryNode.right
		}
	}

	if side == driftLeft {
		for temporaryNode.left != nil {
			temporaryNode = temporaryNode.left
		}
	}
	return temporaryNode
}

func (bst *BinarySearchTree) Height() int {
	return bst.height(bst.root)
}

func (bst *BinarySearchTree) height(root *BinarySearchTreeNode) int {
	if root == nil {
		return 0
	}

	return max(bst.height(root.left), bst.height(root.right)) + 1
}

func (bst *BinarySearchTree) Traverse(order int) ([]*comparable.Comparable, error) {
	switch order {
	case PreOrder:
		return bst.preOrderTraverse(bst.root), nil
	case InOrder:
		return bst.inOrderTraverse(bst.root), nil
	case PostOrder:
		return bst.postOrderTraverse(bst.root), nil
	case LevelOrder:
		return bst.levelOrderTraverse(bst.root), nil
	}

	return nil, fmt.Errorf("order type %d does not exists", order)
}

func (bst *BinarySearchTree) preOrderTraverse(node *BinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	stack := stack.NewStack()
	stack.Push(node)
	i := 0

	for stack.Peek() != nil {
		node := stack.Peek().(*BinarySearchTreeNode)
		sl[i] = node.data
		stack.Pop()
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
		i++
	}
	return sl
}

func (bst *BinarySearchTree) inOrderTraverse(node *BinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	i := 0
	stack := stack.NewStack()
	current := node
	for current != nil || stack.Peek() != nil {
		for current != nil {
			stack.Push(current)
			current = current.left
		}
		node := stack.Peek().(*BinarySearchTreeNode)
		stack.Pop()
		sl[i] = node.data
		i++
		if node.right != nil {
			current = node.right
		}
	}

	return sl
}

func (bst *BinarySearchTree) postOrderTraverse(node *BinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	stackFinal := stack.NewStack()
	stack := stack.NewStack()
	stack.Push(node)
	i := 0

	for stack.Peek() != nil {
		node := stack.Peek().(*BinarySearchTreeNode)
		stack.Pop()
		stackFinal.Push(node)
		if node.left != nil {
			stack.Push(node.left)
		}
		if node.right != nil {
			stack.Push(node.right)
		}
	}

	for stackFinal.Peek() != nil {
		node := stackFinal.Peek().(*BinarySearchTreeNode)
		sl[i] = node.data
		stackFinal.Pop()
		i++
	}

	return sl
}

func (bst *BinarySearchTree) levelOrderTraverse(node *BinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	queue := queue.NewQueue(bst.size)
	i := 0
	queue.Offer(node)
	for !queue.IsEmpty() {
		elem, _ := queue.Peek()
		node = elem.(*BinarySearchTreeNode)
		sl[i] = node.data
		i++
		queue.Poll()
		if node.left != nil {
			queue.Offer(node.left)
		}

		if node.right != nil {
			queue.Offer(node.right)
		}
	}
	return sl
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
