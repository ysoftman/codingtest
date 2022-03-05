/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/
83. Remove Duplicates from Sorted List
Easy
Share
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func deleteDuplicates(head *ListNode) *ListNode {
//     if head == nil {
//         return nil
//     }
//     result := &ListNode{}
//     result.Val = head.Val
//     resultHead := result

//     head = head.Next
//     for head != nil {
//         if head.Val != result.Val {
//             result.Next = &ListNode {
//                 Val: head.Val,
//                 Next: nil,
//             }
//             result = result.Next
//         }
//         head = head.Next
//     }
//     return resultHead
// }

// 입력 노드 자체를 수정하는 방법
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	headOrg := head

	for head != nil {
		curNode := head
		for curNode.Next != nil && curNode.Next.Val == head.Val {
			curNode = curNode.Next
		}
		head.Next = curNode.Next
		head = head.Next
	}
	return headOrg
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
func main() {
	head := makeLinkedList([]int{1, 1, 2})
	printLinkedList(deleteDuplicates(head))

	head = makeLinkedList([]int{1, 1, 2, 3, 3})
	printLinkedList(deleteDuplicates(head))
}
