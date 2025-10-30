/*
https://leetcode.com/problems/same-tree/
100. Same Tree
Easy
Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
Input: p = [1,2], q = [1,null,2]
Output: false

Example 3:
Input: p = [1,2,1], q = [1,1,2]
Output: false
*/
package main

import (
	"fmt"

	"github.com/ysoftman/ysoftmancommon"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree2(p *ysoftmancommon.TreeNode, q *ysoftmancommon.TreeNode) bool {
	pNodes := []ysoftmancommon.TreeNode{}
	qNodes := []ysoftmancommon.TreeNode{}
	if p != nil {
		pNodes = append(pNodes, *p)
	}
	if q != nil {
		qNodes = append(qNodes, *q)
	}
	if len(pNodes) != len(qNodes) {
		return false
	}

	for len(pNodes) > 0 || len(qNodes) > 0 {
		if len(pNodes) != len(qNodes) {
			return false
		}
		p := pNodes[0]
		q := qNodes[0]
		pNodes = pNodes[1:]
		qNodes = qNodes[1:]

		// fmt.Println("p:",p)
		// fmt.Println("q:",q)
		if p.Val != q.Val {
			return false
		}
		if p.Left != nil {
			pNodes = append(pNodes, *p.Left)
			if q.Left == nil {
				return false
			}
		}
		if q.Left != nil {
			if p.Left == nil {
				return false
			}
			qNodes = append(qNodes, *q.Left)
		}
		if p.Right != nil {
			if q.Right == nil {
				return false
			}
			pNodes = append(pNodes, *p.Right)
		}
		if q.Right != nil {
			if p.Right == nil {
				return false
			}
			qNodes = append(qNodes, *q.Right)
		}

	}
	return true
}

func isSameTree(p *ysoftmancommon.TreeNode, q *ysoftmancommon.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if isSameTree(p.Left, q.Left) == false {
		return false
	}
	if isSameTree(p.Right, q.Right) == false {
		return false
	}
	return true
}

func main() {
	left := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3"})
	right := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3"})
	fmt.Println(isSameTree(left, right))

	left = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2"})
	right = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "null", "2"})
	fmt.Println(isSameTree(left, right))

	left = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "1"})
	right = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "1", "2"})
	fmt.Println(isSameTree(left, right))
}
