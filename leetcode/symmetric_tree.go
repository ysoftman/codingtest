/*
https://leetcode.com/problems/symmetric-tree

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:
   1
  / \
 2   2
/ \ / \
3 4 4 3

Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
   1
  / \
 2   2
/ \ / \
  3    3

Input: root = [1,2,2,null,3,null,3]
Output: false
*/
// Definition for a binary tree node.
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isMirror(root, root)
}

func main() {

	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isSymmetric(root1))
}
