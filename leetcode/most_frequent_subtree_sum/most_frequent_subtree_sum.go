/*
https://leetcode.com/problems/most-frequent-subtree-sum/
508. Most Frequent Subtree Sum
Medium

Given the root of a binary tree, return the most frequent subtree sum. If there is a tie, return all the values with the highest frequency in any order.

The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself).

Example 1:
Input: root = [5,2,-3]
Output: [2,-3,4]

Example 2:
Input: root = [5,2,-5]
Output: [2]

Constraints:
The number of nodes in the tree is in the range [1, 104].
-105 <= Node.val <= 105
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
var (
	m            map[int]int
	maxFrequency int
)

func findFrequentTreeSum(root *ysoftmancommon.TreeNode) []int {
	m = make(map[int]int, 0)
	maxFrequency = 0
	findTreeSum(root)
	r := []int{}
	for k, v := range m {
		if v == maxFrequency {
			r = append(r, k)
		}
	}
	return r
}

// postorder 탐색하며 각 노드 위치에서의 합을 구한다
func findTreeSum(root *ysoftmancommon.TreeNode) int {
	if root == nil {
		return 0
	}
	sum := findTreeSum(root.Left) + findTreeSum(root.Right) + root.Val
	m[sum]++
	if m[sum] > maxFrequency {
		maxFrequency = m[sum]
	}
	return sum
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "2", "-3"})
	fmt.Println(findFrequentTreeSum(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "2", "-5"})
	fmt.Println(findFrequentTreeSum(root))
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"1"})
	fmt.Println(findFrequentTreeSum(root))
}
