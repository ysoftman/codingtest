/*
https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
117. Populating Next Right Pointers in Each Node II
Medium
Given a binary tree

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
Initially, all next pointers are set to NULL.

Example 1:
Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Example 2:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 6000].
-100 <= Node.val <= 100


Follow-up:
You may only use constant extra space.
The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.
*/

package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	start := root
	q := make([]*Node, 0)
	q = append(q, root)
	// BFS
	for len(q) > 0 {
		var pre *Node = nil
		siblingQ := []*Node{}
		// sibling node
		for len(q) > 0 {
			// head
			head := q[0]
			if head == nil {
				q = q[1:]
				continue
			}
			// add current node's right/left node to sibling-queue
			siblingQ = append(siblingQ, head.Right)
			siblingQ = append(siblingQ, head.Left)
			head.Next = pre

			pre = head
			// remove head
			q = q[1:]
		}
		// for next level sibling
		q = append(q, siblingQ...)
	}
	return start
}

func printTreeNextValByBFS(root *Node) {
	q := []*Node{}
	q = append(q, root)
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top == nil {
			// fmt.Printf("nil ")
			continue
		}
		fmt.Printf("%v ", top.Val)
		if top.Next != nil {
			fmt.Printf("(%v) ", top.Next.Val)
		} else {
			fmt.Printf("nil ")
		}
		q = append(q, top.Left)
		q = append(q, top.Right)
	}
	fmt.Println()
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Node{
			Val:  3,
			Left: nil,
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	printTreeNextValByBFS(connect(root))
}
