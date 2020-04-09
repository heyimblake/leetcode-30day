package internal

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	inputs := []*ListNode{createLinkedList(1, 1), createLinkedList(1, 5), createLinkedList(1, 6)}
	expected := []*ListNode{createLinkedList(1, 1), createLinkedList(3, 3), createLinkedList(4, 3)}

	for i, input := range inputs {
		actual := middleNode(input)
		if !reflect.DeepEqual(actual, expected[i]) {
			t.Fatalf("Expected %d but got %d.\n", expected[i].Val, actual.Val)
		}
	}
}

func TestBackspaceCompare(t *testing.T) {
	inputs := [][]string{{"ab#c", "ad#c"}, {"ab##", "c#d#"}, {"a##c", "#a#c"}, {"a#c", "b"}, {"###", ""}, {"", ""}, {"ab##", "##ab#c##"}, {"a#a", "a"}}
	expected := []bool{true, true, true, false, true, true, true, true}

	for i, input := range inputs {
		actual := backspaceCompare(input[0], input[1])
		if actual != expected[i] {
			t.Fatalf("Expected %t but got %t.\n", expected[i], actual)
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
