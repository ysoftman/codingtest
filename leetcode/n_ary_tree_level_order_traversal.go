/*
https://leetcode.com/problems/n-ary-tree-level-order-traversal/
429. N-ary Tree Level Order Traversal
Medium
Given an n-ary tree, return the level order traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

Example 1:
Input: root = [1,null,3,2,4,null,5,6]
Output: [[1],[3,2,4],[5,6]]

Example 2:
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]


Constraints:
The height of the n-ary tree is less than or equal to 1000
The total number of nodes is between [0, 104]
*/
package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// bfs
func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	r := [][]int{}
	r = append(r, []int{root.Val})

	q := []*Node{}
	q = append(q, root)

	for len(q) != 0 {
		levelq := []*Node{}
		curLevelValues := []int{}

		for len(q) != 0 {
			curNode := q[0]
			// pop q
			q = q[1:]
			if curNode == nil {
				continue
			}

			for i := 0; i < len(curNode.Children); i++ {
				levelq = append(levelq, curNode.Children[i])
				if curNode.Children[i] != nil {
					curLevelValues = append(curLevelValues, curNode.Children[i].Val)
				}
			}

		}

		if len(curLevelValues) > 0 {
			r = append(r, curLevelValues)
		}

		q = levelq
	}
	return r
}

func main() {
	root := []string{"1", "null", "3", "2", "4", "null", "5", "6"}
	fmt.Println(levelOrder(makeArrayToNode(root)))
	root = []string{"1", "null", "2", "3", "4", "5", "null", "null", "6", "7", "null", "8", "null", "9", "10", "null", "null", "11", "null", "12", "null", "13", "null", "null", "14"}
	fmt.Println(levelOrder(makeArrayToNode(root)))
}
