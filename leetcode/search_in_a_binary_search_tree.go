/*
https://leetcode.com/problems/search-in-a-binary-search-tree/
700. Search in a Binary Search Tree
Easy
You are given the root of a binary search tree (BST) and an integer val.
Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

Example 1:
Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]

Example 2:
Input: root = [4,2,7,1,3], val = 5
Output: []
*/

// go run ./search_in_a_binary_search_tree.go ./ysoftman_common.go
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	node := searchBST(root.Left, val)
	if node != nil {
		return node
	}
	node = searchBST(root.Right, val)
	if node != nil {
		return node
	}
	return nil
}

func main() {
	root := makeArrayToBinaryTree([]string{"4", "2", "7", "1", "3"})
	printTreeByBFS(searchBST(root, 2))
	printTreeByBFS(searchBST(root, 5))
}
