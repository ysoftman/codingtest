/*
https://leetcode.com/problems/insert-into-a-binary-search-tree/
701. Insert into a Binary Search Tree
Medium
You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

Example 1:
Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:

Example 2:
Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]

Example 3:
Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]
*/
// go run ./insert_into_a_binary_search_tree.go
package main

import "github.com/ysoftman/ysoftmancommon"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBSTbyDFS(root *ysoftmancommon.TreeNode, val int) {
	if root == nil {
		return
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = &ysoftmancommon.TreeNode{
				Val: val,
			}
		} else {
			insertIntoBSTbyDFS(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &ysoftmancommon.TreeNode{
				Val: val,
			}
		} else {
			insertIntoBSTbyDFS(root.Right, val)
		}
	}
}

func insertIntoBST(root *ysoftmancommon.TreeNode, val int) *ysoftmancommon.TreeNode {
	if root == nil {
		return &ysoftmancommon.TreeNode{
			Val: val,
		}
	}
	node := root
	insertIntoBSTbyDFS(node, val)
	return root
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"4", "2", "7", "1", "3"})
	ysoftmancommon.PrintTreeNodeByBFS(insertIntoBST(root, 5))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"40", "20", "60", "10", "30", "50", "70"})
	ysoftmancommon.PrintTreeNodeByBFS(insertIntoBST(root, 25))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"4", "2", "7", "1", "3", "null", "null", "null", "null", "null", "null"})
	ysoftmancommon.PrintTreeNodeByBFS(insertIntoBST(root, 5))
}
