/*
https://leetcode.com/problems/path-sum/
112. Path Sum
Easy
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Explanation: The root-to-leaf path with the target sum is shown.

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: false
Explanation: There two root-to-leaf paths in the tree:
(1 --> 2): The sum is 3.
(1 --> 3): The sum is 4.
There is no root-to-leaf path with sum = 5.

Example 3:
Input: root = [], targetSum = 0
Output: false
Explanation: Since the tree is empty, there are no root-to-leaf paths.
*/
package main

import (
	"fmt"
	"strconv"
)

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder dfs
func dfsSum(root *TreeNode, parentSum, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && parentSum+root.Val == targetSum {
		return true
	}

	if dfsSum(root.Left, parentSum+root.Val, targetSum) {
		return true
	}
	if dfsSum(root.Right, parentSum+root.Val, targetSum) {
		return true
	}
	return false
}
func hasPathSum(root *TreeNode, targetSum int) bool {
	parentSum := 0
	return dfsSum(root, parentSum, targetSum)
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
	root := makeArrayToBinaryTree([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "null", "1"})
	printTreeByBFS(root)
	fmt.Println(hasPathSum(root, 22))
	root = makeArrayToBinaryTree([]string{"1", "2", "3"})
	printTreeByBFS(root)
	fmt.Println(hasPathSum(root, 5))
	root = makeArrayToBinaryTree([]string{})
	printTreeByBFS(root)
	fmt.Println(hasPathSum(root, 0))
	root = makeArrayToBinaryTree([]string{"1", "2"})
	printTreeByBFS(root)
	fmt.Println(hasPathSum(root, 1))
	root = makeArrayToBinaryTree([]string{"1", "2"})
	printTreeByBFS(root)
	fmt.Println(hasPathSum(root, 3))
}
