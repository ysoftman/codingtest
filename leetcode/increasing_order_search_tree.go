/*
https://leetcode.com/problems/increasing-order-search-tree/
897. Increasing Order Search Tree
Easy
Given the root of a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only one right child.

Example 1:
Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

Example 2:
Input: root = [5,1,7]
Output: [1,null,5,null,7]

Constraints:
The number of nodes in the given tree will be in the range [1, 100].
0 <= Node.val <= 1000
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
func recursiveIncreasingBST(root, newRoot *TreeNode) {
	if root == nil {
		return
	}

	recursiveIncreasingBST(root.Left, newRoot)
	// fmt.Print(root.Val, " ")
	for newRoot.Right != nil {
		newRoot = newRoot.Right
	}
	newRoot.Right = &TreeNode{root.Val, nil, nil}
	recursiveIncreasingBST(root.Right, newRoot)
}
func increasingBST(root *TreeNode) *TreeNode {
	newRoot := TreeNode{}
	recursiveIncreasingBST(root, &newRoot)
	return newRoot.Right
}
func main() {
	printTreeNodeByDFS(increasingBST(makeArrayToBinaryTreeNode([]string{"5", "3", "6", "2", "4", "null", "8", "1", "null", "null", "null", "7", "9"})))
	fmt.Println()
	printTreeNodeByDFS(increasingBST(makeArrayToBinaryTreeNode([]string{"5", "1", "7"})))
	fmt.Println()
}
