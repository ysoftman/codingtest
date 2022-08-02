/*
https://leetcode.com/problems/binary-tree-paths
257. Binary Tree Paths
Easy
Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

Example 1:
    1
   / \
  2   3
   \
    5
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Example 2:
Input: root = [1]
Output: ["1"]
*/

package main

import (
	"fmt"
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, path string, results *[]string) {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*results = append(*results, path)
		return
	}
	path += "->"
	if root.Left != nil {
		dfs(root.Left, path, results)
	}
	if root.Right != nil {
		dfs(root.Right, path, results)
	}
}
func binaryTreePaths(root *TreeNode) []string {
	path := ""
	results := []string{}
	dfs(root, path, &results)
	// fmt.Println(results)
	return results

}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(binaryTreePaths(root))
}
