/*
https://leetcode.com/problems/rotate-list/
61. Rotate List
Medium
Given the head of a linked list, rotate the list to the right by k places.

Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]

Example 2:
Input: head = [0,1,2], k = 4
Output: [2,0,1]

Constraints:
The number of nodes in the list is in the range [0, 500].
-100 <= Node.val <= 100
0 <= k <= 2 * 109
*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	length := 0
	root := head
	var tail *ListNode = nil
	for head != nil {
		length++
		head = head.Next
		if head != nil {
			tail = head
		}
	}
	head = root
	if length == 0 {
		return head
	}

	k = k % length
	if k == 0 {
		return head
	}

	k = length - k
	// fmt.Println("k:", k)
	// fmt.Println("tail.Val:", tail.Val)
	for k > 0 {
		if k == 1 {
			tail.Next = root
			root = head.Next
			head.Next = nil
			break
		}
		head = head.Next
		k--
	}
	return root
}

func main() {
	printLinkedList(rotateRight(makeLinkedList([]int{1, 2, 3, 4, 5}), 2))
	printLinkedList(rotateRight(makeLinkedList([]int{1, 2, 3, 4, 5}), 12))
	printLinkedList(rotateRight(makeLinkedList([]int{1, 2, 3, 4, 5}), 5))
	printLinkedList(rotateRight(makeLinkedList([]int{0, 1, 2}), 4))
	printLinkedList(rotateRight(makeLinkedList([]int{}), 0))
	printLinkedList(rotateRight(makeLinkedList([]int{}), 1))
}
