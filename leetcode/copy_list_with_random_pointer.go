/*
https://leetcode.com/problems/copy-list-with-random-pointer/
138. Copy List with Random Pointer
Medium
A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.

Example 1:
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

Example 2:
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

Example 3:
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]

Constraints:
0 <= n <= 1000
-104 <= Node.val <= 104
Node.random is null or is pointing to some node in the linked list.
*/
package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node, 0)
	node := head
	newNode := &Node{}
	newNodeHead := newNode
	// copy node next
	for node != nil {
		newNode.Next = &Node{
			Val: node.Val,
		}
		newNode = newNode.Next
		// 생성된 노드 정보 기록
		m[node] = newNode
		node = node.Next
	}
	// copy node random
	node = head
	newNode = newNodeHead
	for node != nil {
		if node.Random == nil {
			newNode.Next.Random = nil
		} else {
			newNode.Next.Random = m[node.Random]
		}

		node = node.Next
		newNode = newNode.Next
		// debugging
		fmt.Printf("%v Next:%v Random:%v\n", newNode, newNode.Next, newNode.Random)
	}

	return newNodeHead.Next
}

func main() {
	// 노드 순서 Next -> 7->13->11->10->1->nil
	// 첫번째 원소는 Val
	// 두번째 원소는 Random -> 0부터 시작하는 노드 인덱스(순서)
	list := [][]string{{"7", "null"}, {"13", "0"}, {"11", "4"}, {"10", "2"}, {"1", "0"}}
	Head := &Node{}
	curNode := Head
	m := make(map[int]*Node, 0)
	cnt := 0
	for i := 0; i < len(list); i++ {
		num, _ := strconv.Atoi(list[i][0])
		n := &Node{
			Val: num,
		}
		curNode.Next = n
		curNode = curNode.Next
		m[cnt] = curNode
		cnt++
	}
	curNode = Head.Next
	for i := 0; i < len(list); i++ {
		if list[i][1] == "null" {
			curNode.Random = nil
		} else {
			num, _ := strconv.Atoi(list[i][1])
			curNode.Random = m[num]
		}
		curNode = curNode.Next
	}
	copyRandomList(Head.Next)
}
