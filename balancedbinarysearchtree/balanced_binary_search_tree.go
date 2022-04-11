package balancedbinarysearchtree

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

type BalancedBinarySearchTreeNode struct {
	data     *comparable.Comparable
	left     *BalancedBinarySearchTreeNode
	right    *BalancedBinarySearchTreeNode
	height   int
	balanced int
}

type BalancedBinarySearchTree struct {
	root *BalancedBinarySearchTreeNode
	size int
}

func NewBalancedBinarySearchTree() *BalancedBinarySearchTree {
	return &BalancedBinarySearchTree{
		size: 0,
	}
}

func NewBalancedBinarySearchTreeNode(data *comparable.Comparable, left *BalancedBinarySearchTreeNode, right *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	return &BalancedBinarySearchTreeNode{
		data:     data,
		left:     left,
		right:    right,
		height:   0,
		balanced: 0,
	}
}

func (bst *BalancedBinarySearchTree) update(n *BalancedBinarySearchTreeNode) {
	nHeightLeft := -1
	nHeightRight := -1
	if n.left != nil {
		nHeightLeft = n.left.height
	}
	if n.right != nil {
		nHeightRight = n.right.height
	}

	n.height = 1 + max(nHeightLeft, nHeightRight)
	n.balanced = nHeightRight - nHeightLeft
}

func (bst *BalancedBinarySearchTree) balance(n *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	if n.balanced == -2 {
		if n.left.balanced <= 0 {
			return bst.leftLeftCase(n)
		} else {
			return bst.leftRightCase(n)
		}
	} else if n.balanced == +2 {
		if n.right.balanced >= 0 {
			return bst.rightRightCase(n)
		} else {
			return bst.rightLeftCase(n)
		}
	}
	return n
}

func (bst *BalancedBinarySearchTree) leftLeftCase(n *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	return bst.rightRotation(n)
}

func (bst *BalancedBinarySearchTree) leftRightCase(n *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	n.left = bst.leftRotation(n.left)
	return bst.leftLeftCase(n)
}

func (bst *BalancedBinarySearchTree) rightRightCase(n *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	return bst.leftRotation(n)
}

func (bst *BalancedBinarySearchTree) rightLeftCase(n *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	n.right = bst.rightRotation(n.right)
	return bst.rightRightCase(n)
}

func (bst *BalancedBinarySearchTree) leftRotation(n *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	newParent := n.right
	n.right = newParent.left
	newParent.left = n
	bst.update(n)
	bst.update(newParent)
	return newParent
}

func (bst *BalancedBinarySearchTree) rightRotation(n *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	newParent := n.left
	n.left = newParent.right
	newParent.right = n
	bst.update(n)
	bst.update(newParent)
	return newParent
}

func (bst *BalancedBinarySearchTree) Add(element *comparable.Comparable) error {
	if bst.Contains(element) {
		return fmt.Errorf("element with index %d already exists in tree", element.GetIndex())
	}
	node := NewBalancedBinarySearchTreeNode(element, nil, nil)
	bst.size++
	bst.root = bst.add(bst.root, node)
	return nil
}

func (bst *BalancedBinarySearchTree) Height() int {
	return bst.root.height + 1
}

func (bst *BalancedBinarySearchTree) add(root *BalancedBinarySearchTreeNode, node *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
	if root == nil {
		return node
	}

	if root.data.GetIndex() > node.data.GetIndex() {
		root.left = bst.add(root.left, node)
	}

	if root.data.GetIndex() < node.data.GetIndex() {
		root.right = bst.add(root.right, node)
	}

	bst.update(root)
	return bst.balance(root)
}

func (bst *BalancedBinarySearchTree) Contains(element *comparable.Comparable) bool {
	_, err := bst.Find(element.GetIndex())
	return err == nil
}

func (bst *BalancedBinarySearchTree) Find(index int) (*comparable.Comparable, error) {
	node := bst.find(bst.root, index)
	if node == nil {
		return nil, fmt.Errorf("cannot find element with index %d", index)
	}
	return node.data, nil
}

func (bst *BalancedBinarySearchTree) find(node *BalancedBinarySearchTreeNode, index int) *BalancedBinarySearchTreeNode {
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

func (bst *BalancedBinarySearchTree) Remove(element *comparable.Comparable) error {
	if !bst.Contains(element) {
		return fmt.Errorf("element with index %d doesn't exists in tree", element.GetIndex())
	}
	bst.root = bst.remove(bst.root, element.GetIndex())
	bst.size--
	return nil
}

func (bst *BalancedBinarySearchTree) remove(node *BalancedBinarySearchTreeNode, index int) *BalancedBinarySearchTreeNode {
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
	bst.update(node)
	return bst.balance(node)
}

func (bst *BalancedBinarySearchTree) drift(side int, node *BalancedBinarySearchTreeNode) *BalancedBinarySearchTreeNode {
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

func (bst *BalancedBinarySearchTree) Traverse(order int) ([]*comparable.Comparable, error) {
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

func (bst *BalancedBinarySearchTree) preOrderTraverse(node *BalancedBinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	stack := stack.NewStack()
	stack.Push(node)
	i := 0

	for stack.Peek() != nil {
		node := stack.Peek().(*BalancedBinarySearchTreeNode)
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

func (bst *BalancedBinarySearchTree) inOrderTraverse(node *BalancedBinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	i := 0
	stack := stack.NewStack()
	current := node
	for current != nil || stack.Peek() != nil {
		for current != nil {
			stack.Push(current)
			current = current.left
		}
		node := stack.Peek().(*BalancedBinarySearchTreeNode)
		stack.Pop()
		sl[i] = node.data
		i++
		if node.right != nil {
			current = node.right
		}
	}

	return sl
}

func (bst *BalancedBinarySearchTree) postOrderTraverse(node *BalancedBinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	stackFinal := stack.NewStack()
	stack := stack.NewStack()
	stack.Push(node)
	i := 0

	for stack.Peek() != nil {
		node := stack.Peek().(*BalancedBinarySearchTreeNode)
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
		node := stackFinal.Peek().(*BalancedBinarySearchTreeNode)
		sl[i] = node.data
		stackFinal.Pop()
		i++
	}

	return sl
}

func (bst *BalancedBinarySearchTree) levelOrderTraverse(node *BalancedBinarySearchTreeNode) []*comparable.Comparable {
	sl := make([]*comparable.Comparable, bst.size)
	queue := queue.NewQueue(bst.size)
	i := 0
	queue.Offer(node)
	for !queue.IsEmpty() {
		elem, _ := queue.Peek()
		node = elem.(*BalancedBinarySearchTreeNode)
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
