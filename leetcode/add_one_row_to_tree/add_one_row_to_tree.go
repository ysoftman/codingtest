/*
https://leetcode.com/problems/add-one-row-to-tree/
623. Add One Row to Tree
Medium

Given the root of a binary tree and two integers val and depth, add a row of nodes with value val at the given depth depth.
Note that the root node is at depth 1.
The adding rule is:

Given the integer depth, for each not null tree node cur at the depth depth - 1, create two tree nodes with value val as cur's left subtree root and right subtree root.
cur's original left subtree should be the left subtree of the new left subtree root.
cur's original right subtree should be the right subtree of the new right subtree root.
If depth == 1 that means there is no depth depth - 1 at all, then create a tree node with value val as the new root of the whole original tree, and the original tree is the new root's left subtree.

Example 1:
Input: root = [4,2,6,3,1,5], val = 1, depth = 2
Output: [4,1,1,2,null,null,6,3,1,5]

Example 2:
Input: root = [4,2,null,3,1], val = 1, depth = 3
Output: [4,2,null,1,1,3,null,null,1]

Constraints:
The number of nodes in the tree is in the range [1, 104].
The depth of the tree is in the range [1, 104].
-100 <= Node.val <= 100
-105 <= val <= 105
1 <= depth <= the depth of tree + 1
*/
package main

import "github.com/ysoftman/ysoftmancommon"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// DFS, time complexity : O(n)
func recursiveAddOneRow(root *ysoftmancommon.TreeNode, val int, depth int, curDepth int) *ysoftmancommon.TreeNode {
	if root == nil || depth < curDepth {
		return nil
	}

	// 1 depth 위치은 root 이전에 insert 해야 하는 예외 상황
	if depth == 1 {
		newNode := &ysoftmancommon.TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
		return newNode
	}

	// left 에 val(node) 추가
	if curDepth+1 == depth {
		newNode := &ysoftmancommon.TreeNode{
			Val:   val,
			Left:  root.Left,
			Right: nil,
		}
		root.Left = newNode
	} else {
		recursiveAddOneRow(root.Left, val, depth, curDepth+1)
	}

	//  right 에 val(node) 추가
	if curDepth+1 == depth {
		newNode := &ysoftmancommon.TreeNode{
			Val:   val,
			Left:  nil,
			Right: root.Right,
		}
		root.Right = newNode
	} else {
		recursiveAddOneRow(root.Right, val, depth, curDepth+1)
	}
	return root
}

func addOneRow(root *ysoftmancommon.TreeNode, val int, depth int) *ysoftmancommon.TreeNode {
	return recursiveAddOneRow(root, val, depth, 1)
}

func main() {
	ysoftmancommon.PrintTreeNodeByBFS(addOneRow(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"4", "2", "6", "3", "1", "5"}), 1, 2))
	ysoftmancommon.PrintTreeNodeByBFS(addOneRow(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"4", "2", "null", "3", "1"}), 1, 3))
	ysoftmancommon.PrintTreeNodeByBFS(addOneRow(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"4", "2", "6", "3", "1", "5"}), 1, 1))
}
