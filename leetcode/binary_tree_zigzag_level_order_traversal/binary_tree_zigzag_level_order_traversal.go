/*
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
103. Binary Tree Zigzag Level Order Traversal
Medium
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100
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

// BFS
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		levelq := []*TreeNode{}
		levelValues := []int{}
		for i := 0; i < len(q); i++ {
			f := q[i]
			if f == nil {
				continue
			}

			// 큐를 이용한 BFS 로 탐색하고, 값을 넣을때만 level 에 따라 right->left 순으로 한다.
			index := i
			if level%2 == 1 {
				index = len(q) - 1 - i
			}
			levelValues = append(levelValues, q[index].Val)

			if f.Left != nil {
				levelq = append(levelq, f.Left)
			}
			if f.Right != nil {
				levelq = append(levelq, f.Right)
			}
		}
		if len(levelValues) == 0 {
			break
		}
		result = append(result, levelValues)
		q = levelq
		level++
	}
	return result
}

func main() {
	fmt.Println(zigzagLevelOrder(makeArrayToBinaryTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})))
	fmt.Println(zigzagLevelOrder(makeArrayToBinaryTreeNode([]string{"3", "9", "20", "5", "3", "15", "7"})))
	fmt.Println(zigzagLevelOrder(makeArrayToBinaryTreeNode([]string{"3"})))
	fmt.Println(zigzagLevelOrder(makeArrayToBinaryTreeNode([]string{})))
	fmt.Println(zigzagLevelOrder(makeArrayToBinaryTreeNode([]string{"0", "2", "4", "1", "null", "3", "-1", "5", "1", "null", "6", "null", "8"})))
}
