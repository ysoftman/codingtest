/*
https://leetcode.com/problems/minimum-depth-of-binary-tree/
111. Minimum Depth of Binary Tree
Easy
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 2
Example 2:

Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5


Constraints:

The number of nodes in the tree is in the range [0, 105].
-1000 <= Node.val <= 1000
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minDepth(root *TreeNode) int {
	// ignore null node
	if root == nil {
		return 0
	}
	// leaf node
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// if left == 0 {
	//     left = 1<<31-1
	// }
	// if right == 0 {
	//     right = 1<<31-1
	// }
	// or
	if left == 0 || right == 0 {
		return left + right + 1
	}

	return min(left, right) + 1
}
func main() {
	node := makeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})
	fmt.Println(minDepth(node))
	node = makeArrayToBinaryTreeNode([]string{"2", "null", "3", "null", "4", "null", "5", "null", "6"})
	fmt.Println(minDepth(node))
	node = makeArrayToBinaryTreeNode([]string{"2", "2", "3", "null", "4", "null", "5", "null", "6"})
	fmt.Println(minDepth(node))
	node = makeArrayToBinaryTreeNode([]string{})
	fmt.Println(minDepth(node))
}
