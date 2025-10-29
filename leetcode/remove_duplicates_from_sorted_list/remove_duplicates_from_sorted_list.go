/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/
83. Remove Duplicates from Sorted List
Easy
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]
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

// func deleteDuplicates(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
//     if head == nil {
//         return nil
//     }
//     result := &ysoftmancommon.ListNode{}
//     result.Val = head.Val
//     resultHead := result

//     head = head.Next
//     for head != nil {
//         if head.Val != result.Val {
//             result.Next = &ysoftmancommon.ListNode {
//                 Val: head.Val,
//                 Next: nil,
//             }
//             result = result.Next
//         }
//         head = head.Next
//     }
//     return resultHead
// }

// 입력 노드 자체를 수정하는 방법
func deleteDuplicates(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	if head == nil {
		return nil
	}
	headOrg := head

	for head != nil {
		curNode := head
		for curNode.Next != nil && curNode.Next.Val == head.Val {
			curNode = curNode.Next
		}
		head.Next = curNode.Next
		head = head.Next
	}
	return headOrg
}

func main() {
	head := ysoftmancommon.MakeLinkedList([]int{1, 1, 2})
	ysoftmancommon.PrintLinkedList(deleteDuplicates(head))

	head = ysoftmancommon.MakeLinkedList([]int{1, 1, 2, 3, 3})
	ysoftmancommon.PrintLinkedList(deleteDuplicates(head))
}
