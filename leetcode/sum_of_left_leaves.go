/*
https://leetcode.com/problems/sum-of-left-leaves/
404. Sum of Left Leaves
Easy
Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.

Example 2:
Input: root = [1]
Output: 0

Constraints:
The number of nodes in the tree is in the range [1, 1000].
-1000 <= Node.val <= 1000
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
func sumOfLeftLeaves(root *TreeNode) int {
	return dfs_SOLL(root, false)
}

func dfs_SOLL(root *TreeNode, leftchild bool) int {
	if root == nil {
		return 0
	}
	// leaf node
	if root.Left == nil && root.Right == nil {
		// only left chlid returns value
		if leftchild {
			return root.Val
		}
	}
	val := dfs_SOLL(root.Left, true)
	val += dfs_SOLL(root.Right, false)
	return val
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})
	fmt.Println(sumOfLeftLeaves(root))
	root = makeArrayToBinaryTreeNode([]string{"1"})
	fmt.Println(sumOfLeftLeaves(root))
}
