/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
235. Lowest Common Ancestor of a Binary Search Tree
Medium
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

// go run ./lowest_common_ancestor_of_a_binary_search_tree.go ./ysoftman_common.go
package main

import (
	"fmt"
)

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// LCA : 두 노드(자신 포함)의 첫번째  공통 부모
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		// if root == p {
		//     fmt.Println("found p")
		// }
		// if root == q {
		//     fmt.Println("found q")
		// }
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
	root := makeArrayToBinaryTreeNode([]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	p := findNode(root, 2)
	q := findNode(root, 8)
	printTreeNodeByBFS(root)
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	root = makeArrayToBinaryTreeNode([]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	p = findNode(root, 2)
	q = findNode(root, 4)
	printTreeNodeByBFS(root)
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	root = makeArrayToBinaryTreeNode([]string{"2", "1"})
	p = findNode(root, 2)
	q = findNode(root, 1)
	printTreeNodeByBFS(root)
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}
