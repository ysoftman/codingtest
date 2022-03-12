/*
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
105. Construct Binary Tree from Preorder and Inorder Traversal
Medium
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]
*/

//  go run ./construct_binary_tree_from_preorder_and_traversal.go ./ysoftman_common.go
package main

import "fmt"

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var preorder_index int = 0

func searchinorder(inorder []int, start, end, v int) int {
	for i := start; i <= end; i++ {
		if inorder[i] == v {
			return i
		}
	}
	return -1
}

func makeTree(preorder, inorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// preorder 의 원소값을 inorder 에서 찾아 그 index 를 기준으로
	// 왼쪽부분에서는 left 오른쪽부분에서 right 노드를 재귀로 찾아간다.
	node := &TreeNode{
		Val: preorder[preorder_index],
	}
	preorder_index++
	if left == right {
		return node
	}
	i := searchinorder(inorder, left, right, node.Val)
	node.Left = makeTree(preorder, inorder, left, i-1)
	node.Right = makeTree(preorder, inorder, i+1, right)

	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preorder_index = 0
	return makeTree(preorder, inorder, 0, len(inorder)-1)
}

func main() {
	printTreeNodeByDFS(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	fmt.Println()
	printTreeNodeByDFS(buildTree([]int{-1}, []int{-1}))
	fmt.Println()
}
