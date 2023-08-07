package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	current := &ListNode{}
	newHead := current
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else if list2.Val < list1.Val {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}
	return newHead.Next
}
