package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
https://leetcode.com/problems/add-two-numbers-ii/description/?envType=study-plan-v2&envId=programming-skills
*/

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	rev1, rev2 := reverse(l1), reverse(l2)
	carry := 0
	newHead := &ListNode{}
	current := newHead
	for rev1 != nil || rev2 != nil || carry != 0 {
		sum := carry
		if rev1 != nil {
			sum += rev1.Val
			rev1 = rev1.Next
		}
		if rev2 != nil {
			sum += rev2.Val
			rev2 = rev2.Next
		}
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}
	//l1 = reverse(l1)
	//l2 = reverse(l2)
	return reverse(newHead.Next)
}

func reverse(l *ListNode) *ListNode {
	var prev, next *ListNode
	for l != nil {
		next = l.Next
		l.Next = prev
		prev = l
		l = next
	}
	return prev
}
