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

func makeArrayToBinaryTree(arr []string) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	temp, _ := strconv.Atoi(arr[0])
	root := &TreeNode{
		Val: temp,
	}
	node := root

	queue := make([]*TreeNode, 0)
	queue = append(queue, node)
	for i := 1; i < len(arr); i++ {
		if len(queue) == 0 {
			continue
		}
		front := queue[0]
		queue = queue[1:]

		if arr[i] != "null" {
			left, _ := strconv.Atoi(arr[i])
			front.Left = &TreeNode{
				Val: left,
			}
			queue = append(queue, front.Left)
		}
		if i+1 >= len(arr) {
			continue
		}
		i++
		if arr[i] != "null" {
			right, _ := strconv.Atoi(arr[i])
			front.Right = &TreeNode{
				Val: right,
			}
			queue = append(queue, front.Right)
		}
	}
	return root
}
func printTreeByBFS(root *TreeNode) {
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top == nil {
			fmt.Printf("nil ")
			continue
		}
		fmt.Printf("%v ", top.Val)
		q = append(q, top.Left)
		q = append(q, top.Right)
	}
	fmt.Println()
}

func main() {
	root := makeArrayToBinaryTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	printTreeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTree([]string{"1", "2", "null", "3", "null", "4", "null", "5"})
	printTreeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTree([]string{"1", "null", "2"})
	printTreeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	printTreeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTree([]string{"1"})
	printTreeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))

	root = makeArrayToBinaryTree([]string{})
	printTreeByBFS(root)
	fmt.Println(levelOrderBFS(root))
	fmt.Println(levelOrder(root))
}
