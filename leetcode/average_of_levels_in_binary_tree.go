/*
https://leetcode.com/problems/average-of-levels-in-binary-tree/
637. Average of Levels in Binary Tree
Easy
Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.


Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]
Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
Hence return [3, 14.5, 11].

Example 2:
Input: root = [3,9,20,15,7]
Output: [3.00000,14.50000,11.00000]


Constraints:
The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1
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
// dfs, time complexity : O(n)
func recursiveAerageOfLevels(root *TreeNode, level int, mv, mcnt map[int]int) {
	if root == nil {
		return
	}
	mv[level] += root.Val
	mcnt[level]++
	recursiveAerageOfLevels(root.Left, level+1, mv, mcnt)
	recursiveAerageOfLevels(root.Right, level+1, mv, mcnt)
}
func averageOfLevels(root *TreeNode) []float64 {
	mv := make(map[int]int)
	mcnt := make(map[int]int)
	recursiveAerageOfLevels(root, 0, mv, mcnt)
	r := []float64{}
	for i := 0; i < len(mv); i++ {
		r = append(r, float64(mv[i])/float64(mcnt[i]))
	}
	return r
}

func main() {
	fmt.Println(averageOfLevels(makeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})))
	fmt.Println(averageOfLevels(makeArrayToBinaryTreeNode([]string{"3", "9", "20", "15", "7"})))
}
