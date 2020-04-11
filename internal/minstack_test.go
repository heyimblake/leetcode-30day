package internal

import "testing"

// LeetCode provided example
func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if stack.GetMin() != -3 {
		t.Fatalf("Expected min of %d but got %d.\n", -3, stack.GetMin())
	}

	// Extra Check
	if stack.Top() != -3 {
		t.Fatalf("Expected top of %d but got %d.\n", -3, stack.Top())
	}

	stack.Pop()

	if stack.Top() != 0 {
		t.Fatalf("Expected top of %d but got %d.\n", 0, stack.Top())
	}

	if stack.GetMin() != -2 {
		t.Fatalf("Expected min of %d but got %d.\n", -2, stack.GetMin())
	}

	stack.Pop()

	// Extra Check
	if stack.Top() != -2 {
		t.Fatalf("Expected top of %d but got %d.\n", -2, stack.Top())
	}

	// Extra Check
	if stack.GetMin() != -2 {
		t.Fatalf("Expected min of %d but got %d.\n", -2, stack.GetMin())
	}
}
