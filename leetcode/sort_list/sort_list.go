/*
https://leetcode.com/problems/sort-list/
148. Sort List
Medium
Given the head of a linked list, return the list after sorting it in ascending order.

Example 1:
Input: head = [4,2,1,3]
Output: [1,2,3,4]

Example 2:
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]

Example 3:
Input: head = []
Output: []

Constraints:
The number of nodes in the list is in the range [0, 5 * 104].
-105 <= Node.val <= 105

Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
*/
package main

import (
	"sort"

	"github.com/ysoftman/ysoftmancommon"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 정렬된 새로운 ListNode 를 만들기
// time complexity(nlogn)
// space complexity(n)
func sortList(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	l := []int{}
	for head != nil {
		l = append(l, head.Val)
		head = head.Next
	}
	sort.Ints(l)
	r := &ysoftmancommon.ListNode{}
	newhead := r
	for _, v := range l {
		r.Next = &ysoftmancommon.ListNode{
			Val: v,
		}
		r = r.Next
	}
	return newhead.Next
}

func main() {
	ysoftmancommon.PrintLinkedList(sortList(ysoftmancommon.MakeLinkedList([]int{4, 2, 1, 3})))
	ysoftmancommon.PrintLinkedList(sortList(ysoftmancommon.MakeLinkedList([]int{-1, 5, 3, 4, 0})))
	ysoftmancommon.PrintLinkedList(sortList(ysoftmancommon.MakeLinkedList([]int{})))
}
