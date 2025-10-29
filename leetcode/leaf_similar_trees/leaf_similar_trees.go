/*
https://leetcode.com/problems/leaf-similar-trees/
872. Leaf-Similar Trees
Easy

Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
Two binary trees are considered leaf-similar if their leaf value sequence is the same.
Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

Example 1:
Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true

Example 2:
Input: root1 = [1,2,3], root2 = [1,3,2]
Output: false

Constraints:
The number of nodes in each tree will be in the range [1, 200].
Both of the given trees will have values in the range [0, 200].
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

func findLeafNode(root *TreeNode, leafVal *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leafVal = append(*leafVal, root.Val)
	}
	if root.Left != nil {
		findLeafNode(root.Left, leafVal)
	}
	if root.Right != nil {
		findLeafNode(root.Right, leafVal)
	}
	return
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1 := []int{}
	findLeafNode(root1, &r1)
	r2 := []int{}
	findLeafNode(root2, &r2)
	if len(r1) != len(r2) {
		return false
	}
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

func main() {
	root1 := makeArrayToBinaryTreeNode([]string{"3", "5", "1", "6", "2", "9", "8", "null", "null", "7", "4"})
	root2 := makeArrayToBinaryTreeNode([]string{"3", "5", "1", "6", "7", "4", "2", "null", "null", "null", "null", "null", "null", "9", "8"})
	fmt.Println(leafSimilar(root1, root2))
	root1 = makeArrayToBinaryTreeNode([]string{"1", "2", "3"})
	root2 = makeArrayToBinaryTreeNode([]string{"1", "3", "2"})
	fmt.Println(leafSimilar(root1, root2))
}
