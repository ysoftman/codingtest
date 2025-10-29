/*
https://leetcode.com/problems/remove-linked-list-elements/
203. Remove Linked List Elements
Easy
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

Example 1:
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

Example 2:
Input: head = [], val = 1
Output: []

Example 3:
Input: head = [7,7,7,7], val = 7
Output: []
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
func removeElements(head *ysoftmancommon.ListNode, val int) *ysoftmancommon.ListNode {
	root := head
	var pre *ysoftmancommon.ListNode = nil
	for head != nil {
		if head.Val == val {
			if pre != nil {
				pre.Next = head.Next
			} else {
				root = head.Next
			}
		} else {
			pre = head
		}
		head = head.Next
	}
	return root
}

func main() {
	ysoftmancommon.PrintLinkedList(removeElements(ysoftmancommon.MakeLinkedList([]int{1, 2, 6, 3, 4, 5, 6}), 6))
	ysoftmancommon.PrintLinkedList(removeElements(ysoftmancommon.MakeLinkedList([]int{}), 1))
	ysoftmancommon.PrintLinkedList(removeElements(ysoftmancommon.MakeLinkedList([]int{7, 7, 7, 7}), 7))
}
