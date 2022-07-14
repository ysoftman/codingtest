/*
https://leetcode.com/problems/insertion-sort-list/
147. Insertion Sort List
Medium
Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's head.

The steps of the insertion sort algorithm:

Insertion sort iterates, consuming one input element each repetition and growing a sorted output list.
At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list and inserts it there.
It repeats until no input elements remain.
The following is a graphical example of the insertion sort algorithm. The partially sorted list (black) initially contains only the first element in the list. One element (red) is removed from the input data and inserted in-place into the sorted list with each iteration.

Example 1:
Input: head = [4,2,1,3]
Output: [1,2,3,4]

Example 2:
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]

Constraints:
The number of nodes in the list is in the range [1, 5000].
-5000 <= Node.val <= 5000
*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	sortedHead := &ListNode{Val: 0}
	for head != nil {
		next := head.Next

		sortedNode := sortedHead.Next
		var insertPos *ListNode = nil
		for sortedNode != nil {
			if sortedNode.Val < head.Val {
				insertPos = sortedNode
			}
			sortedNode = sortedNode.Next
		}
		// 제일 작아 맨 앞에 추가하는 경우
		if insertPos == nil {
			head.Next = sortedHead.Next
			sortedHead.Next = head
		} else {
			head.Next = insertPos.Next
			insertPos.Next = head
		}

		// fmt.Println("head:", head.Val)
		head = next
	}

	return sortedHead.Next
}

func main() {
	head := makeLinkedList([]int{4, 2, 1, 3})
	printLinkedList(insertionSortList(head))
	head = makeLinkedList([]int{-1, 5, 3, 4, 0})
	printLinkedList(insertionSortList(head))
	head = makeLinkedList([]int{1})
	printLinkedList(insertionSortList(head))
	head = makeLinkedList([]int{5, 4, 3, 2, 3, 33, 2, 3, 2, 1})
	printLinkedList(insertionSortList(head))

}
