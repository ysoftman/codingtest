/*
https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
2095. Delete the Middle Node of a Linked List
Medium

You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.

Example 1:
Input: head = [1,3,4,7,1,2,6]
Output: [1,3,4,1,2,6]
Explanation:
The above figure represents the given linked list. The indices of the nodes are written below.
Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
We return the new list after removing this node.

Example 2:
Input: head = [1,2,3,4]
Output: [1,2,4]
Explanation:
The above figure represents the given linked list.
For n = 4, node 2 with value 3 is the middle node, which is marked in red.

Example 3:
Input: head = [2,1]
Output: [2]
Explanation:
The above figure represents the given linked list.
For n = 2, node 1 with value 1 is the middle node, which is marked in red.
Node 0 with value 2 is the only node remaining after removing node 1.


Constraints:

The number of nodes in the list is in the range [1, 105].
1 <= Node.val <= 105
*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle2(head *ListNode) *ListNode {
	// edge case: node 가 1개일땐느 nil 리턴
	if head.Next == nil {
		return nil
	}

	orgHead := head
	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}
	head = orgHead
	i := 1
	for i < cnt/2 {
		head = head.Next
		i++
	}
	head.Next = head.Next.Next
	return orgHead
}

// slow, fast 2개의 포인터를 사용하는  방법
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	orgHead := head
	slow := head
	// fast 는 2번째 노드부터 시작
	fast := head.Next.Next

	// slow 는 한 노드씩 이동
	// fast 는 2 노드씩 이동
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next
	return orgHead
}

func main() {
	head := makeLinkedList([]int{1, 3, 4, 7, 1, 2, 6})
	head = deleteMiddle(head)
	printLinkedList(head)
	head = makeLinkedList([]int{1, 2, 3, 4})
	head = deleteMiddle(head)
	printLinkedList(head)
	head = makeLinkedList([]int{2, 1})
	head = deleteMiddle(head)
	printLinkedList(head)
	head = makeLinkedList([]int{1})
	head = deleteMiddle(head)
	printLinkedList(head)
}
