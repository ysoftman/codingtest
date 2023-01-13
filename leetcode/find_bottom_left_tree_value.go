/*
https://leetcode.com/problems/find-bottom-left-tree-value/
513. Find Bottom Left Tree Value
Medium

Given the root of a binary tree, return the leftmost value in the last row of the tree.

Example 1:
Input: root = [2,1,3]
Output: 1

Example 2:
Input: root = [1,2,3,4,null,5,6,null,null,7]
Output: 7

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
// BFS
func findBottomLeftValue2(root *TreeNode) int {
	r := root.Val
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelq := []*TreeNode{}
		for len(q) > 0 {
			front := q[0]
			q = q[1:]
			if front.Left != nil {
				levelq = append(levelq, front.Left)
			}
			if front.Right != nil {
				levelq = append(levelq, front.Right)
			}
		}
		q = append(q, levelq...)

		// 매 level(depth) 마다 가장 left 값을 기록한다.
		// 맨 마지막에 저장된 값이 가강 깊은 레벨에서의(last row)에서의 가장 왼쪽 값이 된다.
		if len(levelq) > 0 {
			r = levelq[0].Val
		}
	}
	return r
}

// DFS
func findBottomLeftValue(root *TreeNode) int {
	maxdepth := 0
	r := 0
	dfs(root, 1, &maxdepth, &r)
	return r
}

func dfs(root *TreeNode, depth int, maxdepth *int, r *int) {
	if root == nil {
		return
	}

	dfs(root.Left, depth+1, maxdepth, r)
	if *maxdepth < depth {
		*maxdepth = depth
		*r = root.Val
		// fmt.Println(root.Val)
	}
	dfs(root.Right, depth+1, maxdepth, r)
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"2", "1", "3"})
	fmt.Println(findBottomLeftValue(root))
	root = makeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "null", "5", "6", "null", "null", "7"})
	fmt.Println(findBottomLeftValue(root))
	root = makeArrayToBinaryTreeNode([]string{"1"})
	fmt.Println(findBottomLeftValue(root))
}
