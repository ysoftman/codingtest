/*
https://leetcode.com/problems/unique-binary-search-trees-ii/
95. Unique Binary Search Trees II
Medium
Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.

Example 1:
Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

Example 2:
Input: n = 1
Output: [[1]]


Constraints:
1 <= n <= 8
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
func recursiveGenerateTrees(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	r := []*TreeNode{}

	for i := left; i <= right; i++ {
		leftNode := recursiveGenerateTrees(left, i-1)
		rightNode := recursiveGenerateTrees(i+1, right)
		// fmt.Println(i, leftNode, rightNode)
		for _, ln := range leftNode {
			for _, rn := range rightNode {
				r = append(r, &TreeNode{i, ln, rn})
			}
		}
	}
	return r
}
func generateTrees(n int) []*TreeNode {
	return recursiveGenerateTrees(1, n)
}
func main() {
	root := generateTrees(1)
	fmt.Println("-----")
	for _, n := range root {
		printTreeNodeByDFS(n)
		fmt.Println()
	}
	root = generateTrees(2)
	fmt.Println("-----")
	for _, n := range root {
		printTreeNodeByDFS(n)
		fmt.Println()
	}
	root = generateTrees(3)
	fmt.Println("-----")
	for _, n := range root {
		printTreeNodeByDFS(n)
		fmt.Println()
	}
	root = generateTrees(4)
	fmt.Println("-----")
	for _, n := range root {
		printTreeNodeByDFS(n)
		fmt.Println()
	}
	root = generateTrees(5)
	fmt.Println("-----")
	for _, n := range root {
		printTreeNodeByDFS(n)
		fmt.Println()
	}
}
