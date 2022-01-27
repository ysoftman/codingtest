/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list

19. Remove Nth Node From End of List
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]

*/

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	arr := []*ListNode{}
	for node != nil {
		arr = append(arr, node)
		node = node.Next
	}
	if n > len(arr) || n < 1 {
		return head
	}

	targetIdx := len(arr) - n
	if targetIdx == 0 {
		head = nil
		if targetIdx+1 < len(arr) {
			head = arr[targetIdx+1]
		}
		return head
	}
	if targetIdx == len(arr)-1 {
		arr[targetIdx-1].Next = nil
		return head
	}
	arr[targetIdx-1].Next = arr[targetIdx+1]
	return head
}

func printNodeList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v", head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}

func array2LinkNode(arr *[]int) *ListNode {
	node := &ListNode{}
	head := node
	if len(*arr) > 0 {
		node.Val = (*arr)[0]
	}
	for i := 1; i < len(*arr); i++ {
		node.Next = &ListNode{
			Val: (*arr)[i],
		}
		node = node.Next
	}

	return head
}
func main() {
	head := array2LinkNode(&[]int{1, 2, 3, 4, 5})
	printNodeList(head)
	printNodeList(removeNthFromEnd(head, 2))

	head = array2LinkNode(&[]int{1})
	printNodeList(head)
	printNodeList(removeNthFromEnd(head, 1))

	head = array2LinkNode(&[]int{1, 2})
	printNodeList(head)
	printNodeList(removeNthFromEnd(head, 1))

	head = array2LinkNode(&[]int{1, 2})
	printNodeList(head)
	printNodeList(removeNthFromEnd(head, 2))

}
