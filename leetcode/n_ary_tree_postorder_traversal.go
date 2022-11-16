/*
https://leetcode.com/problems/n-ary-tree-postorder-traversal/
590. N-ary Tree Postorder Traversal
Easy
Given the root of an n-ary tree, return the postorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)

Example 1:
Input: root = [1,null,3,2,4,null,5,6]
Output: [5,6,3,2,4,1]

Example 2:
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]

Constraints:
The number of nodes in the tree is in the range [0, 104].
0 <= Node.val <= 104
The height of the n-ary tree is less than or equal to 1000.

Follow up: Recursive solution is trivial, could you do it iteratively?
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
func recursive(root *Node, r *[]int) {
	if root == nil {
		return
	}
	for i := 0; i < len(root.Children); i++ {
		recursive(root.Children[i], r)
	}
	*r = append(*r, root.Val)
}

// recursive iteraion
func postorder2(root *Node) []int {
	r := []int{}
	recursive(root, &r)
	return r
}

// using stack
func postorder(root *Node) []int {
	r := []int{}
	if root == nil {
		return nil
	}
	s := []*Node{root}
	for len(s) > 0 {
		top := s[len(s)-1]
		// pop
		s = s[:len(s)-1]
		r = append(r, top.Val)
		for i := 0; i < len(top.Children); i++ {
			// push
			s = append(s, top.Children[i])
		}
	}
	reverseR := []int{}
	for i := len(r) - 1; i >= 0; i-- {
		reverseR = append(reverseR, r[i])
	}
	return reverseR
}

func main() {
	node := makeArrayToNode([]string{"1", "null", "3", "2", "4", "null", "5", "6"})
	fmt.Println(postorder(node))
}
