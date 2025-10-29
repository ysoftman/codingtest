/*
https://leetcode.com/problems/deepest-leaves-sum/
1302. Deepest Leaves Sum
Medium
Given the root of a binary tree, return the sum of values of its deepest leaves.

Example 1:
Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15

Example 2:
Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 19

Constraints:
The number of nodes in the tree is in the range [1, 104].
1 <= Node.val <= 100
*/

// go run ./deepest_leaves_sum.go

package main

import (
	"fmt"

	ysoftmancommon "github.com/ysoftman/ysoftmancommon"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BFS
func deepestLeavesSum(root *ysoftmancommon.TreeNode) int {
	q := make([]*ysoftmancommon.TreeNode, 0)
	q = append(q, root)
	deepestSum := 0
	for len(q) > 0 {
		currentQ := make([]*ysoftmancommon.TreeNode, 0)
		deepestSum = 0
		for len(q) > 0 {
			head := q[0]
			q = q[1:]
			if head == nil {
				continue
			}
			if head.Left != nil {
				currentQ = append(currentQ, head.Left)
			}
			if head.Right != nil {
				currentQ = append(currentQ, head.Right)
			}

			deepestSum += head.Val
		}
		// fmt.Println("deepestSum:", deepestSum)
		q = append(q, currentQ...)
	}
	return deepestSum
}

func main() {
	node := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "5", "null", "6", "7", "null", "null", "null", "null", "8"})
	fmt.Println(deepestLeavesSum(node))
	node = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"6", "7", "8", "2", "7", "1", "3", "9", "null", "1", "4", "null", "null", "null", "5"})
	fmt.Println(deepestLeavesSum(node))
}
