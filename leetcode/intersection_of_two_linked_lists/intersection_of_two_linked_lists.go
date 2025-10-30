/*
https://leetcode.com/problems/intersection-of-two-linked-lists/
160. Intersection of Two Linked Lists
Easy
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

For example, the following two linked lists begin to intersect at node c1:
The test cases are generated such that there are no cycles anywhere in the entire linked structure.
Note that the linked lists must retain their original structure after the function returns.

Custom Judge:
The inputs to the judge are given as follows (your program is not given these inputs):

intersectVal - The value of the node where the intersection occurs. This is 0 if there is no intersected node.
listA - The first linked list.
listB - The second linked list.
skipA - The number of nodes to skip ahead in listA (starting from the head) to get to the intersected node.
skipB - The number of nodes to skip ahead in listB (starting from the head) to get to the intersected node.
The judge will then create the linked structure based on these inputs and pass the two heads, headA and headB to your program. If you correctly return the intersected node, then your solution will be accepted.

Example 1:
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.

Example 2:
Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'
Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.

Example 3:
Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection
Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.

Constraints:
The number of nodes of listA is in the m.
The number of nodes of listB is in the n.
1 <= m, n <= 3 \* 104
1 <= Node.val <= 105
0 <= skipA < m
0 <= skipB < n
intersectVal is 0 if listA and listB do not intersect.
intersectVal == listA[skipA] == listB[skipB] if listA and listB intersect.

Follow up: Could you write a solution that runs in O(m + n) time and use only O(1) memory?
*/
package main

import "github.com/ysoftman/ysoftmancommon"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity: O(m+n), space complesxity: O(m)
func getIntersectionNode2(headA, headB *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	a := headA
	b := headB
	m := make(map[*ysoftmancommon.ListNode]bool)
	for a != nil {
		m[a] = true
		a = a.Next
	}
	for b != nil {
		if _, ok := m[b]; ok {
			return b
		}
		b = b.Next
	}
	return nil
}

// time complexity: O(m+n), space complexity: O(1)
func getIntersectionNode(headA, headB *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	a := headA
	b := headB
	// a, b 전체 길이가 다르면 nil 이후에 a,b 를 크로스해서 intersection 전까지 길이를 맞춰지도록 한다.
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

func main() {
	commonList := ysoftmancommon.MakeLinkedList([]int{8, 4, 5})
	listHeadA := ysoftmancommon.MakeLinkedList([]int{4, 1})
	listHeadB := ysoftmancommon.MakeLinkedList([]int{5, 6, 1})
	listA := listHeadA
	for listA.Next != nil {
		listA = listA.Next
	}
	listA.Next = commonList
	listB := listHeadB
	for listB.Next != nil {
		listB = listB.Next
	}
	listB.Next = commonList
	ysoftmancommon.PrintLinkedList(listHeadA)
	ysoftmancommon.PrintLinkedList(listHeadB)
	ysoftmancommon.PrintLinkedList(getIntersectionNode2(listHeadA, listHeadB))
	ysoftmancommon.PrintLinkedList(getIntersectionNode(listHeadA, listHeadB))
}
