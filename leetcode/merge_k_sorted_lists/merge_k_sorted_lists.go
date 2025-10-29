/*
https://leetcode.com/problems/merge-k-sorted-lists/
23. Merge k Sorted Lists
Hard
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.

Example 1:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Example 2:
Input: lists = []
Output: []

Example 3:
Input: lists = [[]]
Output: []

Constraints:
k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 104.
*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	node := head

	for {
		min := 1<<31 - 1
		minIdx := 0
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if lists[i].Val < min {
				min = lists[i].Val
				minIdx = i
			}
		}
		if min == 1<<31-1 {
			break
		}
		lists[minIdx] = lists[minIdx].Next
		node.Next = &ListNode{Val: min}
		node = node.Next
	}

	return head.Next
}

func main() {
	lists := []*ListNode{}
	lists = append(lists, makeLinkedList([]int{1, 4, 5}))
	lists = append(lists, makeLinkedList([]int{1, 3, 4}))
	lists = append(lists, makeLinkedList([]int{2, 6}))
	printLinkedList(mergeKLists(lists))

	lists = []*ListNode{}
	printLinkedList(mergeKLists(lists))

	lists = []*ListNode{}
	lists = append(lists, makeLinkedList([]int{}))
	printLinkedList(mergeKLists(lists))

	lists = append(lists, makeLinkedList([]int{1, 3, 5, 98}))
	lists = append(lists, makeLinkedList([]int{2, 4, 7, 7, 10, 12, 50}))
	lists = append(lists, makeLinkedList([]int{12, 67}))
	printLinkedList(mergeKLists(lists))
}
