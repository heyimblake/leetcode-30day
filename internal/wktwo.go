package internal

type ListNode struct {
	Val  int
	Next *ListNode
}

// Middle of the Linked List:
// Given a non-empty, singly linked list with head node, return a middle node of linked list.
// If there are two middle nodes, return the second middle node.
// The number of nodes in the given list will be between 1 and 100.
func middleNode(head *ListNode) *ListNode {
	// Handle null
	if head == nil {
		panic("Provided node cannot be nil.")
	}

	// Handle single element case
	if head.Next == nil {
		return head
	}

	size := 0
	curr := head

	// Count how many elements are in the list
	for curr != nil {
		size++
		curr = curr.Next
	}

	curr = head
	middle := size / 2 //Calculate middle index

	// Find the middle element
	for i := 0; i < middle; i++ {
		curr = curr.Next
	}

	return curr
}
