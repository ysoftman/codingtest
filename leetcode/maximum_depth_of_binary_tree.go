/*
https://leetcode.com/problems/maximum-depth-of-binary-tree/
104. Maximum Depth of Binary Tree
Easye
Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2

Constraints:
The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dfs(recursive approach)
func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth - 1
	}
	return max(dfs(root.Left, depth+1), dfs(root.Right, depth+1))
}
func maxDepth2(root *TreeNode) int {
	return dfs(root, 1)
}

// using BFS approach
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{}
	q = append(q, root)

	result := 1
	depthList := []int{}
	depthList = append(depthList, 1)
	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		depthFront := depthList[0]
		depthList = depthList[1:]
		if front.Left != nil {
			q = append(q, front.Left)
			depthList = append(depthList, depthFront+1)
			if result < depthFront+1 {
				result = depthFront + 1
			}
		}
		if front.Right != nil {
			q = append(q, front.Right)
			depthList = append(depthList, depthFront+1)
			if result < depthFront+1 {
				result = depthFront + 1
			}
		}
	}

	return result
}

func main() {
	// [3,9,20,null,null,15,7]
	t1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	// [1,null,2]
	t2 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(maxDepth(t1))
	fmt.Println(maxDepth(t2))
}
