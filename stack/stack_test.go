package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push("val-01")

	if stack.size != 1 {
		t.Error("The one element should be in stack")
	}

	if stack.Peek() != "val-01" {
		t.Errorf("The last value in stack should be 'val-01', got %s", stack.Peek())
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push("val-01")
	stack.Push("val-02")
	stack.Pop()

	if stack.size != 1 {
		t.Error("The one element should be in stack")
	}

	if stack.Peek() != "val-01" {
		t.Errorf("The last value in stack should be 'val-01', got %s", stack.Peek())
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	stack := NewStack()
	err := stack.Pop()
	if err == nil {
		t.Error("Pop on empty stack should return error")
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push("val-01")
	stack.Push("val-02")

	if stack.Peek() != "val-02" {
		t.Errorf("The last value in stack should be 'val-02', got %s", stack.Peek())
	}
}

func TestPeekOnEmptyStack(t *testing.T) {
	stack := NewStack()

	if stack.Peek() != nil {
		t.Errorf("The stack is empty, peek should return nil, got %s", stack.Peek())
	}

	stack.Push("val-01")
	stack.Pop()

	if stack.Peek() != nil {
		t.Errorf("The stack is empty, peek should return nil, got %s", stack.Peek())
	}

	stack.Push("val-01")
	stack.Push("val-02")
	stack.Pop()
	stack.Pop()
	if stack.Peek() != nil {
		t.Errorf("The stack is empty, peek should return nil, got %s", stack.Peek())
	}
}
