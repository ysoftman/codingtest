/*
https://leetcode.com/problems/invert-binary-tree
226. Invert Binary Tree
Easy
Given the root of a binary tree, invert the tree, and return its root.

Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []
*/
package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// traverse left child
	invertTree(root.Left)
	// traverse right child
	invertTree(root.Right)
	// current(parent)
	// fmt.Println(root)

	// PostOrder (left->right->root order)
	// we swap left, right
	root.Left, root.Right = root.Right, root.Left
	return root
}

func printBFS(root *TreeNode) {
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		n := q[0]
		fmt.Printf("%v ", n.Val)
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
		q = q[1:]
	}
	fmt.Println()
}
func main() {
	// 4,2,7,1,3,6,9
	root := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	printBFS(&root)
	invertTree(&root)
	printBFS(&root)
}
