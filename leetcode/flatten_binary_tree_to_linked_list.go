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
func flatten(root *TreeNode) {
	// inner function 를 사용하는 이유
	// preNode 를 global 변수로 사용하면 초기화 되지 않아 다른 testcase 이전 testcase 값이 남아 있어 실패하게 된다.
	var preNode *TreeNode = nil
	var dfs func(root *TreeNode)
	// DFS preorder 로 탐색
	// return 이 없어 root 를 주소를 수정(노드 추가,삭제)할 수 없고 val,left,right 값만 변경해야 한다.
	// space complexity: in-place O(1)
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
