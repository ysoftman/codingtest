/*
https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
114. Flatten Binary Tree to Linked List
Medium
Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.

Example 1:
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [0]
Output: [0]

Constraints:
The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100

Follow up: Can you flatten the tree in-place (with O(1) extra space)?
*/
package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// DFS preorder 로 탐색, easy approach(별도 결과 저공간을 사용하는 경우)
// space complexity: O(N)
func dfs(root, r *TreeNode) *TreeNode {
	if root == nil {
		return r
	}
	// fmt.Println(r.Val)
	r.Right = &TreeNode{}
	r = r.Right
	r.Val = root.Val
	r = dfs(root.Left, r)
	r = dfs(root.Right, r)
	return r
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	r := TreeNode{}
	dfs(root, &r)
	*root = *r.Right
}

// DFS preorder 로 탐색
// space complexity: in-place O(1)
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// inner function 를 사용하는 이유
	// preNode 를 global 변수로 사용하면 초기화 되지 않아 다른 testcase 이전 testcase 값이 남아 있어 실패하게 된다.
	var preNode *TreeNode = nil
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// fmt.Println(r.Val)
		dfs(root.Right)
		dfs(root.Left)

		root.Left = nil
		root.Right = preNode
		preNode = root
	}
	dfs(root)
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"1", "2", "5", "3", "4", "null", "6"})
	flatten(root)
	printTreeNodeByDFS(root)
	fmt.Println()
	root = makeArrayToBinaryTreeNode([]string{})
	flatten(root)
	printTreeNodeByDFS(root)
	fmt.Println()
	root = makeArrayToBinaryTreeNode([]string{"0"})
	flatten(root)
	printTreeNodeByDFS(root)
	fmt.Println()
}
