/*
https://leetcode.com/problems/path-sum-ii/
113. Path Sum II
Medium
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: []

Example 3:
Input: root = [1,2], targetSum = 0
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
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
func recursiveFindPath(root *TreeNode, target int, values []int, result *[][]int) {
	if root == nil {
		return
	}
	// leaf 노드까지 가는 경우만 결과에 포함
	if target-root.Val == 0 && root.Left == nil && root.Right == nil {
		values = append(values, root.Val)
		temp := make([]int, len(values))
		copy(temp, values)
		*result = append(*result, temp)
		return
	}
	values = append(values, root.Val)
	recursiveFindPath(root.Left, target-root.Val, values, result)
	recursiveFindPath(root.Right, target-root.Val, values, result)
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	recursiveFindPath(root, targetSum, []int{}, &result)
	return result
}

func main() {
	fmt.Println(pathSum(makeArrayToBinaryTreeNode([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"}), 22))
	fmt.Println(pathSum(makeArrayToBinaryTreeNode([]string{"1", "2", "3"}), 5))
	fmt.Println(pathSum(makeArrayToBinaryTreeNode([]string{"1", "2"}), 0))
	fmt.Println(pathSum(makeArrayToBinaryTreeNode([]string{"1", "2"}), 1))
	fmt.Println(pathSum(makeArrayToBinaryTreeNode([]string{"-2", "null", "-3"}), -5))
}
