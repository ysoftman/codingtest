/*
https://leetcode.com/problems/odd-even-linked-list/
328. Odd Even Linked List
Medium

Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.
Note that the relative order inside both the even and odd groups should remain as it was in the input.
You must solve the problem in O(1) extra space complexity and O(n) time complexity.

Example 1:
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]

Example 2:
Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]

Constraints:
The number of nodes in the linked list is in the range [0, 104].
-106 <= Node.val <= 106
*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time complexity: O(n)
// space complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	odd := &ListNode{}
	even := &ListNode{}
	oddHead := odd
	evenHead := even

	// 1,3,5,7 번째 노드 홀수 그룹핑
	// 2,4,6,8 번째 노드 짝수 그룹핑
	cnt := 1
	for head != nil {
		if cnt%2 == 0 {
			even.Next = head
			even = even.Next
		} else {
			odd.Next = head
			odd = odd.Next
		}
		head = head.Next
		cnt++
	}
	even.Next = nil
	odd.Next = nil
	odd.Next = evenHead.Next
	return oddHead.Next
}

func main() {
	head := makeLinkedList([]int{1, 2, 3, 4, 5})
	printLinkedList(oddEvenList(head))
	head = makeLinkedList([]int{2, 1, 3, 5, 6, 4, 7})
	printLinkedList(oddEvenList(head))
}
