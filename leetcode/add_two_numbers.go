/*
// https: //leetcode.com/problems/add-two-numbers/
2. Add Two Numbers
Medium
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node

	carry := 0
	temp := 0
	for l1 != nil || l2 != nil {
		temp = 0
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		temp += carry
		node.Val = temp % 10
		carry = temp / 10

		if carry == 0 && l1 == nil && l2 == nil {
			break
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	if node != nil && carry > 0 {
		node.Val = carry
	}

	return head
}

func main() {
	// Input: l1 = [2,4,3], l2 = [5,6,4]
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	l1Head := l1
	l2Head := l2
	// printLinkedList(l1Head)
	// printLinkedList(l2Head)
	printLinkedList(addTwoNumbers(l1, l2))

	// Input: l1 = [0], l2 = [0]
	l1 = &ListNode{Val: 0}
	l2 = &ListNode{Val: 0}
	printLinkedList(addTwoNumbers(l1, l2))

	// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	l1 = &ListNode{Val: 9}
	l2 = &ListNode{Val: 9}
	l1Head = l1
	l2Head = l2
	for i := 0; i < 7; i++ {
		newNode := &ListNode{
			Val: 9,
		}
		l1.Next = newNode
		l1 = l1.Next
	}
	for i := 0; i < 4; i++ {
		newNode := &ListNode{
			Val: 9,
		}
		l2.Next = newNode
		l2 = l2.Next
	}
	// printLinkedList(l1Head)
	// printLinkedList(l2Head)

	printLinkedList(addTwoNumbers(l1Head, l2Head))
}
