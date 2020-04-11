package internal

import (
	"errors"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

var (
	emptyStackError = errors.New("stack is empty")
)

type MinStackNode struct {
	Val int
	Next *MinStackNode
	Prev *MinStackNode
}

type MinStack struct {
	head *MinStackNode
	tail *MinStackNode
	min int
	length int
}

func Constructor() MinStack {
	return MinStack{}
}

// Push element x onto stack.
func (this *MinStack) Push(x int) {
	oldlen := this.Len()

	var toInsert *MinStackNode

	if oldlen == 0 {
		toInsert = &MinStackNode{Val: x, Next: nil, Prev: this.tail}
		this.min = x
	} else if x < this.min {
		toInsert = &MinStackNode{Val: 2 * x - this.min, Next: nil, Prev: this.tail}
		this.min = x
	} else {
		toInsert = &MinStackNode{Val: x, Next: nil, Prev: this.tail}
	}

	this.length++

	// First entry
	if this.tail == nil || this.head == nil {
		toInsert.Next = toInsert
		toInsert.Prev = toInsert
		this.tail = toInsert
		this.head = toInsert
		return
	}

	this.tail.Next = toInsert
	this.tail = toInsert
}

// Removes the element on top of the stack.
func (this *MinStack) Pop()  {
	if this.Len() == 0 {
		panic(emptyStackError)
		return
	}

	oldtop := this.tail.Val
	this.length--

	if this.length == 0 {
		this.head = nil
		this.tail = nil
	} else {
		newtail := this.tail.Prev
		newtail.Next = nil
		this.tail.Prev = nil
		this.tail = newtail
	}

	if oldtop < this.min {
		this.min = 2*this.min - oldtop
	}
}

// Get the top element.
func (this *MinStack) Top() int {
	if this.Len() == 0 {
		panic(emptyStackError)
	}

	val := this.tail.Val

	if val < this.min {
		return this.min
	}

	return val
}

// Retrieve the minimum element in the stack.
func (this *MinStack) GetMin() int {
	if this.Len() == 0 {
		panic(emptyStackError)
	}

	return this.min
}

func (this *MinStack) Len() int {
	return this.length
}