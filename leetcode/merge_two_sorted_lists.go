/*
https://leetcode.com/problems/merge-two-sorted-lists/
21. Merge Two Sorted Lists
Easy
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]
*/
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	var result, head *ListNode
	result = &ListNode{}
	head = result
	for true {
		if list1 == nil {
			result.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			result.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			result.Val = list1.Val
			list1 = list1.Next
		} else {
			result.Val = list2.Val
			list2 = list2.Next
		}
		if list1 != nil || list2 != nil {
			result.Next = &ListNode{}
			result = result.Next
		} else {
			break
		}
	}
	return head
}
func makeLinkedList(nums []int) (root *ListNode) {
	if len(nums) == 0 {
		return nil
	}
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
func main() {
	list1 := makeLinkedList([]int{1, 2, 4})
	list2 := makeLinkedList([]int{1, 3, 4})
	printLinkedList(mergeTwoLists(list1, list2))
	list1 = makeLinkedList([]int{})
	list2 = makeLinkedList([]int{})
	printLinkedList(mergeTwoLists(list1, list2))
	list1 = makeLinkedList([]int{})
	list2 = makeLinkedList([]int{0})
	printLinkedList(mergeTwoLists(list1, list2))
}
