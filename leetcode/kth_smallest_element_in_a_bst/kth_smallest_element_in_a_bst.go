/*
https://leetcode.com/problems/kth-smallest-element-in-a-bst/
230. Kth Smallest Element in a BST
Medium
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

Constraints:
The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104

Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
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

// inorder traversal by iteration
func kthSmallest2(root *TreeNode, k int) int {
	cnt := 0
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		for root != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop != nil {
			// fmt.Println(pop.Val)
			cnt++
			if cnt == k {
				return pop.Val
			}
			stack = append(stack, pop.Right)
			root = pop.Right
		}
	}
	return 0
}

// inorder traversal by recursion
func recursiveInorderTree(root *TreeNode, k int, cnt, r *int) {
	if root == nil {
		return
	}

	recursiveInorderTree(root.Left, k, cnt, r)
	// fmt.Println(root.Val)
	*cnt++
	if *cnt == k {
		*r = root.Val
	}
	recursiveInorderTree(root.Right, k, cnt, r)
}
func kthSmallest(root *TreeNode, k int) int {
	cnt := 0
	r := 0
	recursiveInorderTree(root, k, &cnt, &r)
	return r
}
func main() {

	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"3", "1", "4", "null", "2"})
	fmt.Println(kthSmallest(root, 1))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "3", "6", "2", "4", "null", "null", "1"})
	fmt.Println(kthSmallest(root, 3))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"41", "37", "44", "24", "39", "42", "48", "1", "35", "38", "40", "null", "43", "46", "49", "0", "2", "30", "36", "null", "null", "null", "null", "null", "null", "45", "47", "null", "null", "null", "null", "null", "4", "29", "32", "null", "null", "null", "null", "null", "null", "3", "9", "26", "null", "31", "34", "null", "null", "7", "11", "25", "27", "null", "null", "33", "null", "6", "8", "10", "16", "null", "null", "null", "28", "null", "null", "5", "null", "null", "null", "null", "null", "15", "19", "null", "null", "null", "null", "12", "null", "18", "20", "null", "13", "17", "null", "null", "22", "null", "14", "null", "null", "21", "23"})
	fmt.Println(kthSmallest(root, 25))
}
