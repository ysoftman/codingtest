/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
82. Remove Duplicates from Sorted List II
Medium
Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

Example 1:
Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:
Input: head = [1,1,1,2,3]
Output: [2,3]

Constraints:
The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
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
func deleteDuplicates2(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	var cur *ysoftmancommon.ListNode = nil
	var root *ysoftmancommon.ListNode = nil
	var pre *ysoftmancommon.ListNode = nil
	var next *ysoftmancommon.ListNode = nil
	for head != nil {
		next = head.Next
		// first node
		if pre == nil {
			if next == nil {
				cur = &ysoftmancommon.ListNode{
					Val: head.Val,
				}
				root = cur
			} else if head.Val != next.Val {
				cur = &ysoftmancommon.ListNode{
					Val: head.Val,
				}
				root = cur
			}
		}
		// last node
		if next == nil {
			if pre == nil {
				cur = &ysoftmancommon.ListNode{
					Val: head.Val,
				}
				root = cur
			} else if head.Val != pre.Val {
				if cur == nil {
					cur = &ysoftmancommon.ListNode{
						Val: head.Val,
					}
					root = cur
				} else {
					cur.Next = &ysoftmancommon.ListNode{
						Val: head.Val,
					}
					cur = cur.Next
				}
			}
		}

		if pre != nil && next != nil {
			if pre.Val != head.Val && head.Val != next.Val {
				if cur == nil {
					cur = &ysoftmancommon.ListNode{
						Val: head.Val,
					}
					root = cur
				} else {
					cur.Next = &ysoftmancommon.ListNode{
						Val: head.Val,
					}
					cur = cur.Next
				}
			}
		}

		pre = head
		head = head.Next
	}
	return root
}

func deleteDuplicates(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	// head 이전에 빈값의 root 노드를 둔다.
	root := &ysoftmancommon.ListNode{
		Val:  0,
		Next: head,
	}
	pre := root

	for head != nil {
		// skip duplicates
		if head.Next != nil && head.Next.Val == head.Val {
			for head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			pre.Next = head.Next
		} else { // 중복되지 않았다면 다음 노드로 이동해서, 체크 계속
			pre = pre.Next
		}
		head = head.Next
	}
	return root.Next
}

func main() {
	ysoftmancommon.PrintLinkedList(deleteDuplicates(ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 3, 4, 4, 5})))
	ysoftmancommon.PrintLinkedList(deleteDuplicates(ysoftmancommon.MakeLinkedList([]int{1, 1, 1, 2, 3})))
	ysoftmancommon.PrintLinkedList(deleteDuplicates(ysoftmancommon.MakeLinkedList([]int{1})))
	ysoftmancommon.PrintLinkedList(deleteDuplicates(ysoftmancommon.MakeLinkedList([]int{})))
	ysoftmancommon.PrintLinkedList(deleteDuplicates(ysoftmancommon.MakeLinkedList([]int{1, 1, 1, 1, 1})))
	ysoftmancommon.PrintLinkedList(deleteDuplicates(ysoftmancommon.MakeLinkedList([]int{1, 2, 2})))
	ysoftmancommon.PrintLinkedList(deleteDuplicates(ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 3})))
}
