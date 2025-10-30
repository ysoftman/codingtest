/*
https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/
671. Second Minimum Node In a Binary Tree
Easy

Given a non-empty special binary tree consisting of nodes with the non-negative value, where each node in this tree has exactly two or zero sub-node. If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes. More formally, the property root.val = min(root.left.val, root.right.val) always holds.

Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.

If no such second minimum value exists, output -1 instead.

Example 1:
Input: root = [2,2,5,null,null,5,7]
Output: 5
Explanation: The smallest value is 2, the second smallest value is 5.

Example 2:
Input: root = [2,2,2]
Output: -1
Explanation: The smallest value is 2, but there isn't any second smallest value.

Constraints:
The number of nodes in the tree is in the range [1, 25].
1 <= Node.val <= 231 - 1
root.val == min(root.left.val, root.right.val) for each internal node of the tree.
*/
package main

import (
	"fmt"
	"sort"

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
func dfs(root *ysoftmancommon.TreeNode, r *[]int) {
	if root == nil {
		return
	}
	*r = append(*r, root.Val)
	dfs(root.Left, r)
	dfs(root.Right, r)
}

func findSecondMinimumValue(root *ysoftmancommon.TreeNode) int {
	r := []int{}
	dfs(root, &r)
	sort.Ints(r)
	// fmt.Println("r:", r)
	max1 := r[0]
	for i := 0; i < len(r); i++ {
		if max1 < r[i] {
			return r[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(findSecondMinimumValue(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"2", "2", "5", "null", "null", "5", "7"})))
	fmt.Println(findSecondMinimumValue(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"2", "2", "2"})))
	fmt.Println(findSecondMinimumValue(ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "1", "3", "1", "1", "3", "4", "3", "1", "1", "1", "3", "8", "4", "8", "3", "3", "1", "6", "2", "1"})))
}
