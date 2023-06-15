/*
https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
1161. Maximum Level Sum of a Binary Tree
Medium
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

Example 1:
Input: root = [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.

Example 2:
Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2

Constraints:
The number of nodes in the tree is in the range [1, 104].
-105 <= Node.val <= 105
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
// using BFS
func maxLevelSum(root *TreeNode) int {
	maxSum := (1<<31 - 1) * -1
	maxLevel := 1
	level := 1
	levelNodes := []*TreeNode{}
	levelNodes = append(levelNodes, root)
	for len(levelNodes) > 0 {
		sum := 0
		curLevelNodes := []*TreeNode{}
		for len(levelNodes) > 0 {
			curNode := levelNodes[len(levelNodes)-1]
			sum += curNode.Val
			// pop
			levelNodes = levelNodes[:len(levelNodes)-1]

			// push
			if curNode.Left != nil {
				curLevelNodes = append(curLevelNodes, curNode.Left)
			}
			if curNode.Right != nil {
				curLevelNodes = append(curLevelNodes, curNode.Right)
			}
		}
		if maxSum < sum {
			maxSum = sum
			maxLevel = level
		}
		levelNodes = curLevelNodes
		level++
	}

	return maxLevel
}

func main() {
	fmt.Println(maxLevelSum(makeArrayToBinaryTreeNode([]string{"1", "7", "0", "7", "-8", "null", "null"})))
	fmt.Println(maxLevelSum(makeArrayToBinaryTreeNode([]string{"989", "null", "10250", "98693", "-89388", "null", "null", "null", "-32127"})))
	fmt.Println(maxLevelSum(makeArrayToBinaryTreeNode([]string{"-100", "-200", "-300", "-20", "-5", "-10", "null"})))
}
