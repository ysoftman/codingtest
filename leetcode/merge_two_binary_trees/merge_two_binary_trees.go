/*
https://leetcode.com/problems/merge-two-binary-trees/
617. Merge Two Binary Trees
Easy
You are given two binary trees root1 and root2.
Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

Return the merged tree.
Note: The merging process must start from the root nodes of both trees.

Example 1:
Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
Output: [3,4,5,5,4,null,7]

Example 2:
Input: root1 = [1], root2 = [1,2]
Output: [2,2]
*/

// go run ./merge_two_binary_trees.go
package main

import (
	"fmt"

	"github.com/ysoftman/ysoftmancommon"
)

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func recursiveMergeTrees(root1 *ysoftmancommon.TreeNode, root2 *ysoftmancommon.TreeNode) *ysoftmancommon.TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 != nil && root2 == nil {
		return &ysoftmancommon.TreeNode{
			Val:   root1.Val,
			Left:  recursiveMergeTrees(root1.Left, nil),
			Right: recursiveMergeTrees(root1.Right, nil),
		}
	}
	if root1 == nil && root2 != nil {
		return &ysoftmancommon.TreeNode{
			Val:   root2.Val,
			Left:  recursiveMergeTrees(nil, root2.Left),
			Right: recursiveMergeTrees(nil, root2.Right),
		}
	}
	result := &ysoftmancommon.TreeNode{}
	result.Val = root1.Val + root2.Val
	result.Left = recursiveMergeTrees(root1.Left, root2.Left)
	result.Right = recursiveMergeTrees(root1.Right, root2.Right)
	return result
}

func mergeTrees(root1 *ysoftmancommon.TreeNode, root2 *ysoftmancommon.TreeNode) *ysoftmancommon.TreeNode {
	return recursiveMergeTrees(root1, root2)
}

func main() {
	root1 := &ysoftmancommon.TreeNode{
		Val: 1,
		Left: &ysoftmancommon.TreeNode{
			Val: 3,
			Left: &ysoftmancommon.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &ysoftmancommon.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	root2 := &ysoftmancommon.TreeNode{
		Val: 2,
		Left: &ysoftmancommon.TreeNode{
			Val:  1,
			Left: nil,
			Right: &ysoftmancommon.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &ysoftmancommon.TreeNode{
			Val:  3,
			Left: nil,
			Right: &ysoftmancommon.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	result := mergeTrees(root1, root2)
	ysoftmancommon.PrintTreeNodeByDFS(result)
	fmt.Println()
	ysoftmancommon.PrintTreeNodeByBFS(result)
}
