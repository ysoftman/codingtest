/*
https://leetcode.com/problems/linked-list-cycle-ii/
142. Linked List Cycle II
Medium
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.

Constraints:
The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.

Follow up: Can you solve it using O(1) (i.e. constant) memory?
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
// using hashmap
// time complexity: O(n)
// space complexity: O(n)
func detectCycle2(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	m := make(map[*ysoftmancommon.ListNode]bool)
	for head != nil {
		if _, exist := m[head.Next]; exist {
			return head.Next
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

// rabbit and tortoise 알고리즘으로 순환 여부와 순환 시작노드 찾기
// 1. rabbit 은 한번에 2칸씩, tortoise 은 한번에 1칸씩 움직인다.
// 2. 순환 리스트라면 언젠간 둘은 만나게 된다. 만나지 않으면 둘중 nil 을 만나서 종료(비순환 리스트임)
// 3. 둘이 만나면, tortoise 만 head 보내고, 둘아 한번에 1칸씩 움직여서 만나는 곳이 순환 노드의 시작점이다.
// time complexity: O(n)
// space complexity: O(1)
func detectCycle(head *ysoftmancommon.ListNode) *ysoftmancommon.ListNode {
	tortoise := head
	rabbit := head
	isCycle := false
	for tortoise != nil && rabbit != nil {
		tortoise = tortoise.Next
		if rabbit.Next == nil {
			return nil
		}
		rabbit = rabbit.Next.Next
		// 둘이 만나면 순환 리스트임
		if tortoise == rabbit {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	// 이제 거북이만 head 두고, 둘다 1칸씩 움직여서 처음으로 만나는 노드가 순환의 시작부분
	tortoise = head
	for tortoise != rabbit {
		tortoise = tortoise.Next
		rabbit = rabbit.Next
	}
	return tortoise // 또는 rabbit
}

func main() {
	// CycleLinkedList 만들어서 확인
	fmt.Println(detectCycle(ysoftmancommon.MakeCycleLinkedList([]int{3, 2, 0, -4}, 1)))
	fmt.Println(detectCycle(ysoftmancommon.MakeCycleLinkedList([]int{1, 2}, 0)))
	fmt.Println(detectCycle(ysoftmancommon.MakeCycleLinkedList([]int{1}, -1)))
}
