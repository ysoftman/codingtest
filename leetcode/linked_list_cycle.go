/*
https://leetcode.com/problems/linked-list-cycle/
141. Linked List Cycle
Easy
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
*/
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	for head != nil {
		if head.Val == 1<<31-1 {
			return true
		}
		// mark traversa;
		head.Val = 1<<31 - 1
		head = head.Next
	}
	return false
}

func makeCycleLinkedList(slice []int, pos int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	root := &ListNode{}
	head := root
	var temp *ListNode = nil
	var pre *ListNode = nil
	for i, v := range slice {
		head.Val = v
		if i < len(slice)-1 {
			head.Next = &ListNode{}
		}
		if i == pos {
			temp = head
			// fmt.Println("i:", i, ", pos:", pos, ", temp:", temp)
		}
		pre = head
		head = head.Next
	}
	// make Cycle
	if pre != nil && temp != nil {
		pre.Next = temp
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
	if !hasCycle(makeCycleLinkedList([]int{3, 2, 0, -4}, 1)) {
		printLinkedList(makeCycleLinkedList([]int{3, 2, 0, -4}, 1))
	}
	fmt.Println(hasCycle(makeCycleLinkedList([]int{3, 2, 0, -4}, 1)))
	fmt.Println(hasCycle(makeCycleLinkedList([]int{1, 2}, 0)))
	fmt.Println(hasCycle(makeCycleLinkedList([]int{1}, -1)))

}
