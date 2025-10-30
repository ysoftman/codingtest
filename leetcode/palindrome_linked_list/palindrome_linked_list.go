/*
https://leetcode.com/problems/palindrome-linked-list/
234. Palindrome Linked List
Easy
Given the head of a singly linked list, return true if it is a palindrome.

Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false

Constraints:
The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9

Follow up: Could you do it in O(n) time and O(1) space?
*/
package main

import (
	"fmt"

	"github.com/ysoftman/ysoftmancommon"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Palindrome, 똑바로 읽으나 뒤에서부터 읽으나 같은 문자열
func isPalindrome2(head *ysoftmancommon.ListNode) bool {
	if head == nil {
		return false
	}
	s := []int{}
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	j := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func reverseList(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	pre := (*ysoftmancommon.ListNode)(nil)
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// time complexity: O(N)
func isPalindrome(head *ysoftmancommon.ListNode) bool {
	// 중간 노드 찾기
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// fast 리스트 처음 -> 중간
	fast = head
	// slow 리스트 중간 -> 끝 순서를 반대로 변경
	slow = reverseList(slow)
	// for debugging...
	// t := slow
	// for t != nil {
	// 	fmt.Println("slow:", t.Val)
	// 	t = t.Next
	// }
	// t = fast
	// for t != nil {
	// 	fmt.Println("fast:", t.Val)
	// 	t = t.Next
	// }

	// fast(처음 -> 중간), slow(중간 -> 끝)을 비교해서 같아야 palindrome
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(ysoftmancommon.MakeLinkedList([]int{1, 2, 2, 1})))
	fmt.Println(isPalindrome(ysoftmancommon.MakeLinkedList([]int{1, 2})))
	fmt.Println(isPalindrome(ysoftmancommon.MakeLinkedList([]int{1})))
	fmt.Println(isPalindrome(ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 5, 5, 3})))
	fmt.Println(isPalindrome(ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 2, 1})))
	fmt.Println(isPalindrome(ysoftmancommon.MakeLinkedList([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})))
}
