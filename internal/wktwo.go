package internal

import (
	"reflect"
)

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

// Backspace String Compare:
// Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
// 1 <= S.length <= 200
// 1 <= T.length <= 200
// S and T only contain lowercase letters and '#' characters.
func backspaceCompare(S, T string) bool {
	// Length zero case
	if len(S) == 0 && len(T) == 0 {
		return true
	}

	// Make char(rune) arrays
	memS := make([]rune, 0)
	memT := make([]rune, 0)
	iS := 0
	iT := 0

	// Add each char of S to array (or backspace)
	if len(S) > 0 {
		for _, char := range S {
			if char == '#' {
				if iS > 0 {
					iS--
					memS = memS[:iS]
				}
			} else {
				memS = append(memS, char)
				iS++
			}
		}
	}

	// Add each char of T to array (or backspace)
	if len(T) > 0 {
		for _, char := range T {
			if char == '#' {
				if iT > 0 {
					iT--
					memT = memT[:iT]
				}
			} else {
				memT = append(memT, char)
				iT++
			}
		}
	}

	// Compare arrays
	return reflect.DeepEqual(memS, memT)
}
