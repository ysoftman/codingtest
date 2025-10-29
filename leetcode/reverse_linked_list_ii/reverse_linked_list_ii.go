/*
https://leetcode.com/problems/reverse-linked-list-ii/
92. Reverse Linked List II
Medium
Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]

Example 2:
Input: head = [5], left = 1, right = 1
Output: [5]

Constraints:
The number of nodes in the list is n.
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
*/
package main

import "github.com/ysoftman/ysoftmancommon"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ysoftmancommon.ListNode, left int, right int) *ysoftmancommon.ListNode {
	base := &ysoftmancommon.ListNode{
		0,
		head,
	}
	cnt := 0
	var leftNodePre *ysoftmancommon.ListNode
	var leftNode *ysoftmancommon.ListNode
	var ppre *ysoftmancommon.ListNode = base
	var pre *ysoftmancommon.ListNode = base
	for head != nil {
		cnt++
		if cnt == left {
			leftNode = head
			leftNodePre = pre
		}
		if cnt == right {
			leftNodePre.Next = head
			leftNode.Next = head.Next
		}
		ppre = pre
		pre = head
		head = head.Next
		if cnt > left && cnt <= right {
			pre.Next = ppre
		}
	}
	return base.Next
}

func main() {
	head := []int{1, 2, 3, 4, 5}
	ysoftmancommon.PrintLinkedList(reverseBetween(ysoftmancommon.MakeLinkedList(head), 2, 4))

	head = []int{5}
	ysoftmancommon.PrintLinkedList(reverseBetween(ysoftmancommon.MakeLinkedList(head), 1, 1))

	head = []int{0, 1, 3, 2, 5, 2, 3, 1, 6}
	ysoftmancommon.PrintLinkedList(reverseBetween(ysoftmancommon.MakeLinkedList(head), 1, 5))
}
