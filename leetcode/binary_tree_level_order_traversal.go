/*
https://leetcode.com/problems/binary-tree-level-order-traversal/
102. Binary Tree Level Order Traversal
Medium
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
*/

// go run ./binary_tree_level_order_traversal.go ./ysoftman_common.go
package main

import (
	"fmt"
)

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// using BFS
func levelOrderBFS(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	// result = append(result, []int{root.Val})
	queue := []*TreeNode{root}
	levelQueue := []*TreeNode{}
	levelVal := []int{}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front != nil {
			levelQueue = append(levelQueue, front.Left)
			levelQueue = append(levelQueue, front.Right)
			levelVal = append(levelVal, front.Val)
		}

		if len(queue) == 0 {
			if len(levelVal) > 0 {
				result = append(result, levelVal)
				levelVal = []int{}
			}
			queue = levelQueue
			levelQueue = []*TreeNode{}
		}

	}
	return result
}

// using DFS
func dfsTreeLevel(root *TreeNode, level int, result map[int][]int) {
	if root == nil {
		return
	}
	result[level] = append(result[level], root.Val)
	dfsTreeLevel(root.Left, level+1, result)
	dfsTreeLevel(root.Right, level+1, result)
}
func levelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int, 0)
	dfsTreeLevel(root, 0, m)
	result := [][]int{}
	for i := 0; i < len(m); i++ {
		result = append(result, m[i])
	}
	return result
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})
	printTreeNodeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTreeNode([]string{"1", "2", "null", "3", "null", "4", "null", "5"})
	printTreeNodeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTreeNode([]string{"1", "null", "2"})
	printTreeNodeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})
	printTreeNodeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTreeNode([]string{"1"})
	printTreeNodeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTreeNode([]string{})
	printTreeNodeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))
}
