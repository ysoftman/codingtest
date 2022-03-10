/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
Easy
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.
According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

Example 2:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

Example 3:
Input: root = [2,1], p = 2, q = 1
Output: 2
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

// LCA : 두 노드(자신 포함)의 첫번째  공통 부모
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
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

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	node := findNode(root.Right, val)
	if node != nil {
		return node
	}
	node = findNode(root.Left, val)
	if node != nil {
		return node
	}
	return nil

}
func main() {
	root := makeArrayToBinaryTree([]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	p := findNode(root, 2)
	q := findNode(root, 8)
	printTreeByBFS(root)
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	root = makeArrayToBinaryTree([]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	p = findNode(root, 2)
	q = findNode(root, 4)
	printTreeByBFS(root)
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	root = makeArrayToBinaryTree([]string{"2", "1"})
	p = findNode(root, 2)
	q = findNode(root, 1)
	printTreeByBFS(root)
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}
