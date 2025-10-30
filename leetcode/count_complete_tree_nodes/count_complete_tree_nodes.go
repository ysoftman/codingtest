/*
https://leetcode.com/problems/count-complete-tree-nodes/
222. Count Complete Tree Nodes
Medium
Given the root of a complete binary tree, return the number of the nodes in the tree.

According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Design an algorithm that runs in less than O(n) time complexity.

Example 1:
Input: root = [1,2,3,4,5,6]
Output: 6

Example 2:
Input: root = []
Output: 0

Example 3:
Input: root = [1]
Output: 1

Constraints:
The number of nodes in the tree is in the range [0, 5 * 104].
0 <= Node.val <= 5 * 104
The tree is guaranteed to be complete.
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
// DFS, time complexity: O(n)
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return 1 + left + right
}

// time complexity: O(logN)
func getHeight(root *TreeNode) int {
	if root == nil {
		// 마지막 leaf 노드가 꽉 찼는지 모르기 때문에 -1 로 꽉찬노드까지를 높이로 파악한다.
		return -1
	}
	// left 노드 기준으로 전체 tree height 파악
	return 1 + getHeight(root.Left)
}

// DFS, time complexity: O(logN*logN)
// BinarySearchTree(bst) 노드가 꽉차있는 경우 노드 개수는 2^h(트리높이)-1 이다.
func countNodes(root *TreeNode) int {
	h := getHeight(root)
	if h < 0 {
		return 0
	}
	if getHeight(root.Right) == h-1 {
		return 1<<h + countNodes(root.Right)
	}
	return 1<<(h-1) + countNodes(root.Left)
}

func main() {
	fmt.Println(countNodes(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "5", "6"})))
	fmt.Println(countNodes(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{})))
	fmt.Println(countNodes(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1"})))
	fmt.Println(countNodes(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3", "4"})))
}
