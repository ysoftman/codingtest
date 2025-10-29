/*
https://leetcode.com/problems/partition-list/
86. Partition List
Medium
Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.

Example 1:
Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]

Example 2:
Input: head = [2,1], x = 2
Output: [1,2]

Constraints:
The number of nodes in the list is in the range [0, 200].
-100 <= Node.val <= 100
-200 <= x <= 200
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
func partition(head *ysoftmancommon.ListNode, x int) *ysoftmancommon.ListNode {
	root := &ysoftmancommon.ListNode{
		0,
		head,
	}
	less := root
	var pre *ysoftmancommon.ListNode
	var geStart *ysoftmancommon.ListNode
	for head != nil {
		if geStart == nil {
			if head.Val >= x {
				geStart = head
			} else {
				less = head
			}
		}
		if geStart != nil && head.Val < x {
			headNext := head.Next

			// insert lessNode
			less.Next = head
			less = less.Next
			less.Next = geStart

			// remove lessNode
			if pre != nil {
				pre.Next = headNext
			}

			head = headNext
			continue
		}
		pre = head
		head = head.Next
	}
	return root.Next
}

func main() {
	head := ysoftmancommon.MakeLinkedList([]int{1, 4, 3, 2, 5, 2})
	ysoftmancommon.PrintLinkedList(partition(head, 3))

	head = ysoftmancommon.MakeLinkedList([]int{2, 1})
	ysoftmancommon.PrintLinkedList(partition(head, 2))

	head = ysoftmancommon.MakeLinkedList([]int{6, 4, 2, 6, 3, 4, 5})
	ysoftmancommon.PrintLinkedList(partition(head, 3))
}
