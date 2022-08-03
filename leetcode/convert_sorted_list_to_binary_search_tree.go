/*
https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
109. Convert Sorted List to Binary Search Tree
Medium
Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.


Example 1:
Input: head = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.

Example 2:
Input: head = []
Output: []

Constraints:
The number of nodes in head is in the range [0, 2 * 104].
-105 <= Node.val <= 105
*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recursiveSortedListToBST(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	// array 가 아니기 때문에 다음과 같이 중간 값을 바로 파악할 수 없다.
	// mid := (right-left)/2 + left
	temp := left
	mid := left
	// temp 로 두단계식 움직일 수 있다면 mid 는 그것의 반만 웅직이면 된다.
	for temp != right && temp.Next != right {
		temp = temp.Next.Next
		mid = mid.Next
	}

	root := &TreeNode{Val: mid.Val}
	root.Left = recursiveSortedListToBST(left, mid)
	root.Right = recursiveSortedListToBST(mid.Next, right)
	return root
}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return recursiveSortedListToBST(head, nil)
}

func main() {
	printTreeNodeByBFS(sortedListToBST(makeLinkedList([]int{-10, -3, 0, 5, 9})))
	printTreeNodeByBFS(sortedListToBST(makeLinkedList([]int{})))
}
