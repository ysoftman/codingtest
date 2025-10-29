/*
https://leetcode.com/problems/reorder-list/
143. Reorder List
Medium
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]

Constraints:

The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000
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
// using stack
func reorderList2(head *ysoftmancommon.ListNode) {
	stack := []*ysoftmancommon.ListNode{}
	node := head
	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}
	node = head
	for node != nil && len(stack) > 0 {
		tail := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// fmt.Println(tail)
		if tail == node {
			node.Next = nil
			break
		}
		if tail == node.Next {
			node.Next.Next = nil
			break
		}

		next := node.Next
		node.Next = tail
		node = node.Next
		node.Next = next
		node = node.Next
	}
}

// 중간 이전과 이후(reverse 순서로 변경)해서 합치기
// 1,2,3,4,5,6
// 중간노드 찾기 3
// 중간노드 이후 반대방향으로 구성 1,2,3,6,5,4
// head-mid, mid.next~tail 두분으로 나누어 합치기 1,6,2,5,3,4
func reorderList(head *ysoftmancommon.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	node := head
	mid := head
	// 중간 노드 찾기
	for node.Next != nil && node.Next.Next != nil {
		node = node.Next.Next
		mid = mid.Next
	}

	// 중간 노드 이후 반대 방향으로 재구성
	pre := mid
	next := mid.Next
	for next.Next != nil {
		node = next.Next
		next.Next = node.Next
		node.Next = pre.Next
		pre.Next = node
	}

	// 중간을 기준으로 이전,이후 리스트합치기
	l1 := head
	l2 := mid.Next
	for l1 != mid {
		mid.Next = l2.Next
		l2.Next = l1.Next
		l1.Next = l2
		l1 = l2.Next
		l2 = mid.Next
	}
}

func main() {
	head := ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 4})
	reorderList(head)
	ysoftmancommon.PrintLinkedList(head)
	head = ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 4, 5})
	reorderList(head)
	ysoftmancommon.PrintLinkedList(head)
	head = ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 4, 5, 6})
	reorderList(head)
	ysoftmancommon.PrintLinkedList(head)
	head = ysoftmancommon.MakeLinkedList([]int{1})
	reorderList(head)
	ysoftmancommon.PrintLinkedList(head)
}
