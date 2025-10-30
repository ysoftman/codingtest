/*
https://leetcode.com/problems/minimum-absolute-difference-in-bst/
530. Minimum Absolute Difference in BST
Easy
Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

Example 1:
Input: root = [4,2,6,1,3]
Output: 1

Example 2:
Input: root = [1,0,48,null,null,12,49]
Output: 1

Constraints:
The number of nodes in the tree is in the range [2, 104].
0 <= Node.val <= 105

Note: This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/
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
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
		    236

	104             701

null   227        null  911

BinarySearchTree(BST) 는 inorder 탐색 순서로 정렬되어 있다.
104 < 227 < 236 < 701 < 911

227 노드에선 104(pre) 와 비교
236 노드에선 227(pre) 와 비교
701 노드에선 236(pre) 와 비교
911 노드에선 701(pre) 와 비교 하면 된다.
*/
var pre *TreeNode = nil

func inOrder(root *TreeNode, mindiff *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, mindiff)
	if pre != nil {
		// fmt.Println(root.Val, pre.Val)
		*mindiff = min(*mindiff, abs(root.Val, pre.Val))
	}
	pre = root
	inOrder(root.Right, mindiff)
}
func getMinimumDifference(root *TreeNode) int {
	r := 1<<31 - 1
	pre = nil
	inOrder(root, &r)
	return r
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"4", "2", "6", "1", "3"})
	fmt.Println(getMinimumDifference(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "0", "48", "null", "null", "12", "49"})
	fmt.Println(getMinimumDifference(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"236", "104", "701", "null", "227", "null", "911"})
	fmt.Println(getMinimumDifference(root))
}
