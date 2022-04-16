/*
https://leetcode.com/problems/reverse-linked-list
206. Reverse Linked List
Easy
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []
*/

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	printList(head)
	if head == nil {
		return nil
	}
	previous := head
	head = head.Next
	previous.Next = nil
	for head != nil {
		orgNext := head.Next
		head.Next = previous
		previous = head
		head = orgNext
	}
	printList(previous)
	return previous
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%v", node.Val)
		node = node.Next
		if node != nil {
			fmt.Printf("->")
		}
	}
	fmt.Println()
}

func main() {
	// 1,2,3,4,5
	node := &ListNode{Val: 1}
	head := node
	node.Next = &ListNode{Val: 2}
	node = node.Next
	node.Next = &ListNode{Val: 3}
	node = node.Next
	node.Next = &ListNode{Val: 4}
	node = node.Next
	node.Next = &ListNode{Val: 5}
	reverseList(head)
}
