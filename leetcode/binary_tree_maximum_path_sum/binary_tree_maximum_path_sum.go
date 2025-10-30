/*
https://leetcode.com/problems/binary-tree-maximum-path-sum/
124. Binary Tree Maximum Path Sum
Hard

A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.

Example 1:
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

Example 2:
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

Constraints:
The number of nodes in the tree is in the range [1, 3 * 104].
-1000 <= Node.val <= 1000
*/
package main

import (
	"fmt"

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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pathSum(root *ysoftmancommon.TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	leftSum := pathSum(root.Left, maxSum)
	rightSum := pathSum(root.Right, maxSum)
	// path 는 노드를 한번만 거쳐가야 하기 때문에
	// left, right 하나만 지나 가야 한다.
	// 따라서 현재노드에서 최대 합은 다음 3가지 중 1개
	// 1. 현재 노드 값만 (left,right 가 음수도 있기 때문에)
	// 2. left,right 중 큰것만 합한것
	// 3. 상위 노드 path 없이 left->cur->right 만 거쳐가는 path 의 경우조 체크
	// 그런데 리턴 값을 상위 노드에서 포함여부를 체크하기 때문에
	// 3번은 리턴값 조건에서 빠진다.
	curSum := max(root.Val, root.Val+max(leftSum, rightSum))

	// 현재 노드의 최대 합이 트리 전체에서 최대 합인지 체크
	// 위 3번 조건 포함해서 체크
	*maxSum = max(*maxSum, max(curSum, root.Val+leftSum+rightSum))
	return curSum
}

func maxPathSum(root *ysoftmancommon.TreeNode) int {
	maxSum := -1 << 31
	pathSum(root, &maxSum)
	return maxSum
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3"})
	fmt.Println(maxPathSum(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"-10", "9", "20", "null", "null", "15", "7"})
	fmt.Println(maxPathSum(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "null", "1"})
	fmt.Println(maxPathSum(root))
}
