/*
https://leetcode.com/problems/diameter-of-binary-tree/
543. Diameter of Binary Tree
Easy

Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.

Example 1:
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:
Input: root = [1,2]
Output: 1

Constraints:
The number of nodes in the tree is in the range [1, 104].
-100 <= Node.val <= 100
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
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	maxDiameter(root, &diameter)
	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxDiameter(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	ld := maxDiameter(root.Left, diameter)
	rd := maxDiameter(root.Right, diameter)
	// 현재 위치에서 최대 직경은 left, right 자식 노드의 깊이(path)를 합한것과 비교
	if *diameter < ld+rd {
		*diameter = ld + rd
	}
	// 현재 위치에서 직경은 left, right 자식 노드의 깊이(path) 중 큰것 + 1(현재 노드)
	return max(ld, rd) + 1
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "5"})
	fmt.Println(diameterOfBinaryTree(root))
	root = makeArrayToBinaryTreeNode([]string{"1", "2"})
	fmt.Println(diameterOfBinaryTree(root))
}
