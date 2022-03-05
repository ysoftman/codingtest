/*
https://leetcode.com/problems/middle-of-the-linked-list/

876. Middle of the Linked List
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.

Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	cnt := 0
	middleHead := head
	for head.Next != nil {
		head = head.Next
		if cnt%2 == 0 {
			middleHead = middleHead.Next
		}
		cnt++
	}
	return middleHead
}

func main() {
	printLinkedList(middleNode(makeLinkedList([]int{1, 2, 3, 4, 5})))
	printLinkedList(middleNode(makeLinkedList([]int{1, 2, 3, 4, 5, 6})))
}

func makeLinkedList(nums []int) (root *ListNode) {
	head := &ListNode{}
	root = head
	for i, v := range nums {
		head.Val = v
		if i < len(nums)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return root
}
func printLinkedList(head *ListNode) {
	fmt.Printf("[")
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Printf(",")
		}
		head = head.Next
	}
	fmt.Printf("]")
	fmt.Println()
}
