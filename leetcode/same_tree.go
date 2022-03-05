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

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pNodes := []TreeNode{}
	qNodes := []TreeNode{}
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
func main() {
	left := &TreeNode{
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
	right := &TreeNode{
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
	fmt.Println(isSameTree(left, right))

	left = &TreeNode{
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
	}
	right = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(left, right))

	left = &TreeNode{
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
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	right = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(left, right))
}
