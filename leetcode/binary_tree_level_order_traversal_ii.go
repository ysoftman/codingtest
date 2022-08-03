/*
https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
107. Binary Tree Level Order Traversal II
Medium
Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
*/
package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// using BFS
func levelOrderBottom(root *TreeNode) [][]int {
	r := [][]int{}
	if root == nil {
		return r
	}
	temp := [][]int{}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		curQ := make([]*TreeNode, 0)
		t := []int{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			t = append(t, node.Val)

			if node.Left != nil {
				curQ = append(curQ, node.Left)
			}
			if node.Right != nil {
				curQ = append(curQ, node.Right)
			}
		}
		q = curQ
		temp = append(temp, t)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		r = append(r, temp[i])
	}
	return r
}

func main() {
	fmt.Println(levelOrderBottom(makeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})))
	fmt.Println(levelOrderBottom(makeArrayToBinaryTreeNode([]string{"1"})))
	fmt.Println(levelOrderBottom(makeArrayToBinaryTreeNode([]string{})))
}
