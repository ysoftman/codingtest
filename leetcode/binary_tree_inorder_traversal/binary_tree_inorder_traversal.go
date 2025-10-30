/*
https://leetcode.com/problems/binary-tree-inorder-traversal/
94. Binary Tree Inorder Traversal
Easy
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]
*/
package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, result)
	*result = append(*result, node.Val)
	inOrder(node.Right, result)
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inOrder(root, &result)
	fmt.Println("result:", result)
	return result
}

func main() {
	node := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(node))
}
