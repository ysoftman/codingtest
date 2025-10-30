/*
https://leetcode.com/problems/balanced-binary-tree/
110. Balanced Binary Tree
Easy
Given a binary tree, determine if it is height-balanced.
For this problem, a height-balanced binary tree is defined as:
a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Example 3:
Input: root = []
Output: true

Constraints:
The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104
*/
package main

import (
	"fmt"

	"github.com/ysoftman/ysoftmancommon"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
	    1
	  2  3
	 4 5 6 x
	8

==> balanced

	    1
	  2  3
	 4 5 x x
	8

==> ubalanced

	1

x   2

	x  3

==> ubalanced
*/
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS 로 각 레벨에서의 높이 파악
// left, right 서브트리의 노드의 높이 차이가 1보다 크면 안된다.
func checkBalance(root *ysoftmancommon.TreeNode) int {
	if root == nil {
		return 0
	}
	left := checkBalance(root.Left)
	if left == -1 {
		return -1
	}
	right := checkBalance(root.Right)
	if right == -1 {
		return -1
	}

	// left, right 차이가 1 이상이면 unbalanced 트리다.
	if abs(left-right) > 1 {
		return -1
	}
	// balanced 라면 큰쪽 노드를 기준으로
	return max(left, right) + 1
}

func isBalanced(root *ysoftmancommon.TreeNode) bool {
	if root == nil {
		return true
	}
	if checkBalance(root) >= 0 {
		return true
	}
	return false
}

func main() {
	node := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "2", "3", "3", "null", "null", "4", "4"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "null", "2"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "null", "2", "null", "3"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "5", "6", "null", "8"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "5", "null", "null", "8"})
	fmt.Println(isBalanced(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "2", "3", "null", "null", "3", "4", "null", "null", "4"})
	fmt.Println(isBalanced(node))
}
