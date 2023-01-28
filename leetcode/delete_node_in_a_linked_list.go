/*
https://leetcode.com/problems/delete-node-in-a-linked-list/
237. Delete Node in a Linked List
Medium
Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list, instead you will be given access to the node to be deleted directly.

It is guaranteed that the node to be deleted is not a tail node in the list.

Example 1:
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.

Example 2:
Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.

Constraints:
The number of the nodes in the given list is in the range [2, 1000].
-1000 <= Node.val <= 1000
The value of each node in the list is unique.
The node to be deleted is in the list and is not a tail node
*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 삭제할 노드에 다음 노드의 값을 가져오고, 다음 노드를 삭제하는 방법(꼼수?)
func deleteNode2(node *ListNode) {
	node.Val = node.Next.Val
	// 삭제할 노드가 tail이 아니니 node.Next 는 nil 이 아님이 보장된다.
	node.Next = node.Next.Next
}

// 좀더 간결하게, 삭제할 노드를 다음 노드로 변경(삭제할 노드의 포인터값을 변경)
func deleteNode(node *ListNode) {
	*node = *(node.Next)
}

func main() {
	node := makeLinkedList([]int{4, 5, 1, 9})
	printLinkedList(node)
	deleteNode(node.Next)
	printLinkedList(node)

	node = makeLinkedList([]int{4, 5, 1, 9})
	printLinkedList(node)
	deleteNode(node.Next.Next)
	printLinkedList(node)
}
