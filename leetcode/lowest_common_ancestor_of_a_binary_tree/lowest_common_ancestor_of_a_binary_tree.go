/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
236. Lowest Common Ancestor of a Binary Tree
Medium
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.

Example 2:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

Example 3:
Input: root = [1,2], p = 1, q = 2
Output: 1

Constraints:
The number of nodes in the tree is in the range [2, 105].
-109 <= Node.val <= 109
All Node.val are unique.
p != q
p and q will exist in the tree.
*/

// go run ./lowest_common_ancestor_of_a_binary_tree.go
package main

import (
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
func lowestCommonAncestor(root, p, q *ysoftmancommon.TreeNode) *ysoftmancommon.TreeNode {
	if root == nil {
		return nil
	}
	// p 나 q 를 찾은 경우
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// left, right 둘다 찾으면 현재 root 가 LCA
	// left, right 둘중 하나만 찾으면 상위 올리고, 마지막까지 둘중 하나만 리턴하면
	// p 의 자식노드가 q 또는 q 의 자식노드가 p 인 경우다.
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

func main() {
	r := lowestCommonAncestor(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"}),
		ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5"}), ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1"}))
	ysoftmancommon.PrintTreeNodeByBFS(r)
	r = lowestCommonAncestor(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"}),
		ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5"}), ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"4"}))
	ysoftmancommon.PrintTreeNodeByBFS(r)
	r = lowestCommonAncestor(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2"}),
		ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1"}), ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"2"}))
	ysoftmancommon.PrintTreeNodeByBFS(r)
}
