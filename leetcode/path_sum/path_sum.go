/*
https://leetcode.com/problems/path-sum/
112. Path Sum
Easy
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Explanation: The root-to-leaf path with the target sum is shown.

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: false
Explanation: There two root-to-leaf paths in the tree:
(1 --> 2): The sum is 3.
(1 --> 3): The sum is 4.
There is no root-to-leaf path with sum = 5.

Example 3:
Input: root = [], targetSum = 0
Output: false
Explanation: Since the tree is empty, there are no root-to-leaf paths.
*/

// go run ./path_sum.go ./ysoftman_common.go
package main

import (
	"fmt"
)

//* Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// preorder dfs
func dfsSum(root *TreeNode, parentSum, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && parentSum+root.Val == targetSum {
		return true
	}

	if dfsSum(root.Left, parentSum+root.Val, targetSum) {
		return true
	}
	if dfsSum(root.Right, parentSum+root.Val, targetSum) {
		return true
	}
	return false
}
func hasPathSum(root *TreeNode, targetSum int) bool {
	parentSum := 0
	return dfsSum(root, parentSum, targetSum)
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "null", "1"})
	printTreeNodeByBFS(root)
	fmt.Println(hasPathSum(root, 22))
	root = makeArrayToBinaryTreeNode([]string{"1", "2", "3"})
	printTreeNodeByBFS(root)
	fmt.Println(hasPathSum(root, 5))
	root = makeArrayToBinaryTreeNode([]string{})
	printTreeNodeByBFS(root)
	fmt.Println(hasPathSum(root, 0))
	root = makeArrayToBinaryTreeNode([]string{"1", "2"})
	printTreeNodeByBFS(root)
	fmt.Println(hasPathSum(root, 1))
	root = makeArrayToBinaryTreeNode([]string{"1", "2"})
	printTreeNodeByBFS(root)
	fmt.Println(hasPathSum(root, 3))
}
