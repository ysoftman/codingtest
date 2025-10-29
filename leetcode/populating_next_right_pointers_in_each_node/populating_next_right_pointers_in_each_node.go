/*
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
116. Populating Next Right Pointers in Each Node
Medium
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
Initially, all next pointers are set to NULL.

Example 1:
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Example 2:
Input: root = []
Output: []
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
	q := []*Node{}
	q = append(q, root)
	cnt, exp := 0, 0
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top == nil {
			continue
		}
		cnt++
		// 1(1<<0), 2(1<<1), 4(1<<2), 8(1<<3), 16(1<<4) ... 번째 노드 next 는 nil
		if cnt == 1<<exp {
			top.Next = nil
			exp++
			cnt = 0
		} else if len(q) > 0 && q[0] != nil {
			top.Next = q[0]
		}

		q = append(q, top.Left)
		q = append(q, top.Right)

	}
	return root
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
			Val: 3,
			Left: &Node{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	printTreeNextValByBFS(connect(root))
}
