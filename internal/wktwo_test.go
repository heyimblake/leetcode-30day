package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	expected := []*ListNode{createLinkedList(1, 1), createLinkedList(1, 5), createLinkedList(1, 6)}
	actual := []*ListNode{createLinkedList(1, 1), createLinkedList(3, 3), createLinkedList(4, 3)}

	for i, b := range expected {
		if !reflect.DeepEqual(middleNode(b), actual[i]) {
			fmt.Printf("Expected %d but got %d.\n", b.Val, actual[i].Val)
			t.Fail()
		}
	}
}

func createLinkedList(starting, n int) *ListNode {
	head := &ListNode{Val: starting, Next: nil}

	if n == 1 {
		return head
	}

	curr := head

	for i := 1; i < n; i++ {
		curr.Next = &ListNode{Val: starting + i, Next: nil}
		curr = curr.Next
	}

	return head
}
