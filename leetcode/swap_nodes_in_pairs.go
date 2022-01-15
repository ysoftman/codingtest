// https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/
package main

import "fmt"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var start *ListNode = head.Next
	for head != nil {
		var NextNode *ListNode = nil
		var NextPos *ListNode = nil
		if head.Next == nil {
			head = head.Next
			continue
		}
		if head.Next.Next != nil {
			NextPos = head.Next.Next
			if NextPos.Next != nil {
				NextNode = NextPos.Next
			} else {
				NextNode = NextPos
			}
		}
		head.Next.Next = head
		head.Next = NextNode

		head = NextPos
	}
	return start

}

func printNode(node *ListNode) {
	for node != nil {
		fmt.Printf("%v ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	node := &ListNode{Val: 1}
	start := node
	node.Next = &ListNode{Val: 2}
	node = node.Next
	node.Next = &ListNode{Val: 3}
	node = node.Next
	node.Next = &ListNode{Val: 4}
	node = node.Next
	printNode(start)
	printNode(swapPairs(start))

	node = &ListNode{Val: 1}
	start = node
	node.Next = &ListNode{Val: 2}
	node = node.Next
	node.Next = &ListNode{Val: 3}
	node = node.Next
	printNode(start)
	printNode(swapPairs(start))
}
