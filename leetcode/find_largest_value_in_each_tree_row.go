/*
https://leetcode.com/problems/find-largest-value-in-each-tree-row/
515. Find Largest Value in Each Tree Row
Medium
Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).

Example 1:
Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]

Example 2:
Input: root = [1,2,3]
Output: [1,3]

Constraints:
The number of nodes in the tree will be in the range [0, 104].
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

// using BFS
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := []int{root.Val}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelQ := []*TreeNode{}
		for len(q) > 0 {
			front := q[0]
			q = q[1:]
			if front.Left != nil {
				levelQ = append(levelQ, front.Left)
			}
			if front.Right != nil {
				levelQ = append(levelQ, front.Right)
			}
		}
		// get max value from current level vallues
		if len(levelQ) > 0 {
			maxval := levelQ[0].Val
			for _, v := range levelQ {
				if v.Val > maxval {
					maxval = v.Val
				}
			}
			r = append(r, maxval)
		}
		q = levelQ
	}
	return r
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"1", "3", "2", "5", "3", "null", "9"})
	fmt.Println(largestValues(root))
	root = makeArrayToBinaryTreeNode([]string{"1", "2", "3"})
	fmt.Println(largestValues(root))
	root = makeArrayToBinaryTreeNode([]string{"1", "2", "-3"})
	fmt.Println(largestValues(root))
	root = makeArrayToBinaryTreeNode([]string{"1"})
	fmt.Println(largestValues(root))

}
