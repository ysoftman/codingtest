/*
https://leetcode.com/problems/house-robber-iii/
337. House Robber III
Medium
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.

Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if two directly-linked houses were broken into on the same night.

Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.

Example 1:
Input: root = [3,2,3,null,3,null,1]
Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

Example 2:
Input: root = [3,4,5,1,3,null,1]
Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

Constraints:
The number of nodes in the tree is in the range [1, 104].
0 <= Node.val <= 104
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
// dfs + dynamic programming
// time complexity: O(n)
func dfs_rob(root *ysoftmancommon.TreeNode, dp map[*ysoftmancommon.TreeNode]int) int {
	if root == nil {
		return 0
	}

	if val, exist := dp[root]; exist {
		return val
	}

	// case1 : root+child.child
	case1money := root.Val
	if root.Left != nil {
		case1money += dfs_rob(root.Left.Left, dp)
		case1money += dfs_rob(root.Left.Right, dp)
	}
	if root.Right != nil {
		case1money += dfs_rob(root.Right.Left, dp)
		case1money += dfs_rob(root.Right.Right, dp)
	}
	// case2 : child
	case2money := dfs_rob(root.Left, dp)
	case2money += dfs_rob(root.Right, dp)

	// directly-linked houses(node, parent-child) 를 털면 경찰에 연락되니
	// case1 또는 case2 경우 중 큰 값을 선택한다.
	// 이런식 방식이 모든 경우의 수를 다 따지진 않지만 가장 큰값을 찾을 수 있다.(greedy 접근방식)
	// dp 를 사용하지 않으면 time limit exceed
	dp[root] = max(case1money, case2money)
	return dp[root]
}

func rob(root *ysoftmancommon.TreeNode) int {
	dp := make(map[*ysoftmancommon.TreeNode]int)
	return dfs_rob(root, dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"3", "2", "3", "null", "3", "null", "1"})
	fmt.Println(rob(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"3", "4", "5", "1", "3", "null", "1"})
	fmt.Println(rob(root))
}
