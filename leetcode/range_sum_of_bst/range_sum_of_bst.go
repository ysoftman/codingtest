/*
https://leetcode.com/problems/range-sum-of-bst/
938. Range Sum of BST
Easy

Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high].

Example 1:
Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.

Example 2:
Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.

Constraints:
The number of nodes in the tree is in the range [1, 2 * 104].
1 <= Node.val <= 105
1 <= low <= high <= 105
All Node.val are unique.
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
// search all nodes
func rangeSumBST2(root *ysoftmancommon.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if low <= root.Val && root.Val <= high {
		sum = root.Val
	}
	return sum + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func rangeSumBST(root *ysoftmancommon.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	// low 보다 작으면 오른쪽 노드를 탐색
	if low > root.Val {
		return rangeSumBST(root.Right, low, high)
	}
	// high 보다 크면 왼쪽 노드를 탐색
	if high < root.Val {
		return rangeSumBST(root.Left, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"10", "5", "15", "3", "7", "null", "18"})
	low := 7
	high := 15
	fmt.Println(rangeSumBST(root, low, high))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"10", "5", "15", "3", "7", "13", "18", "1", "null", "6"})
	low = 6
	high = 10
	fmt.Println(rangeSumBST(root, low, high))
}
