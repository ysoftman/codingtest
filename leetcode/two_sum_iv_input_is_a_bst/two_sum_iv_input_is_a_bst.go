/*
https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
653. Two Sum IV - Input is a BST
Easy
Given the root of a Binary Search Tree and a target number k, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:
Input: root = [5,3,6,2,4,null,7], k = 9
Output: true

Example 2:
Input: root = [5,3,6,2,4,null,7], k = 28
Output: false
*/

// go run ./two_sum_iv_input_is_a_bst.go
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

func dfsTwoSum(node *ysoftmancommon.TreeNode, hashmap map[int]int, k int) bool {
	if node == nil {
		return false
	}
	if _, exist := hashmap[k-node.Val]; exist {
		return true
	} else {
		hashmap[node.Val] = 1
	}

	if dfsTwoSum(node.Left, hashmap, k) {
		return true
	}
	if dfsTwoSum(node.Right, hashmap, k) {
		return true
	}
	return false
}

func findTarget(root *ysoftmancommon.TreeNode, k int) bool {
	hashmap := make(map[int]int)
	return dfsTwoSum(root, hashmap, k)
}

func main() {
	node := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "3", "6", "2", "4", "null", "7"})
	ysoftmancommon.PrintTreeNodeByBFS(node)
	fmt.Println(findTarget(node, 9))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "3", "6", "2", "4", "null", "7"})
	ysoftmancommon.PrintTreeNodeByBFS(node)
	fmt.Println(findTarget(node, 28))
}
