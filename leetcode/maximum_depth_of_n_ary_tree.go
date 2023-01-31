/*
https://leetcode.com/problems/maximum-depth-of-n-ary-tree/
559. Maximum Depth of N-ary Tree

Given a n-ary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

Example 1:
Input: root = [1,null,3,2,4,null,5,6]
Output: 3

Example 2:
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: 5

Constraints:
The total number of nodes is in the range [0, 104].
The depth of the n-ary tree is less than or equal to 1000.
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func traversal(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for i := 0; i < len(root.Children); i++ {
		// 자식 노드 탐색 후 리턴되었으니 +1 해줘야 한다.
		depth = max(depth, 1+traversal(root.Children[i]))
	}
	return depth
}
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + traversal(root)
}

func main() {
	fmt.Println(maxDepth(makeArrayToNode([]string{"1", "null", "3", "2", "4", "null", "5", "6"})))
	fmt.Println(maxDepth(makeArrayToNode([]string{"1", "null", "2", "3", "4", "5", "null", "null", "6", "7", "null", "8", "null", "9", "10", "null", "null", "11", "null", "12", "null", "13", "null", "null", "14"})))
}
