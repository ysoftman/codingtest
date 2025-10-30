/*
https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
1026. Maximum Difference Between Node and Ancestor
Medium

Given the root of a binary tree, find the maximum value v for which there exist different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.
A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.

Example 1:
Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
Output: 7
Explanation: We have various ancestor-node differences, some of which are given below :
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.

Example 2:
Input: root = [1,null,2,null,0,3]
Output: 3

Constraints:
The number of nodes in the tree is in the range [2, 5000].
0 <= Node.val <= 105
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
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 현재 노드까지 왔을때 최대, 최소 값을 기준으로 최대 차이를 계산
func findMaxDiff(root *TreeNode, minValue, maxValue int) int {
	if root == nil {
		return maxValue - minValue
	}
	minValue = min(minValue, root.Val)
	maxValue = max(maxValue, root.Val)
	return max(findMaxDiff(root.Left, minValue, maxValue), findMaxDiff(root.Right, minValue, maxValue))
}
func maxAncestorDiff(root *TreeNode) int {
	return findMaxDiff(root, root.Val, root.Val)
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"8", "3", "10", "1", "6", "null", "14", "null", "null", "4", "7", "13"})
	fmt.Println(maxAncestorDiff(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "null", "2", "null", "0", "3"})
	fmt.Println(maxAncestorDiff(root))
}
