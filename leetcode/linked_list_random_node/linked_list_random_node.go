/*
https://leetcode.com/problems/linked-list-random-node/
382. Linked List Random Node
Medium
Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.

Implement the Solution class:

Solution(ListNode head) Initializes the object with the head of the singly-linked list head.
int getRandom() Chooses a node randomly from the list and returns its value. All the nodes of the list should be equally likely to be chosen.

Example 1:
Input
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 3, 2, 2, 3]

Explanation
Solution solution = new Solution([1, 2, 3]);
solution.getRandom(); // return 1
solution.getRandom(); // return 3
solution.getRandom(); // return 2
solution.getRandom(); // return 2
solution.getRandom(); // return 3
// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.

Constraints:
The number of nodes in the linked list will be in the range [1, 104].
-104 <= Node.val <= 104
At most 104 calls will be made to getRandom.

Follow up:
What if the linked list is extremely large and its length is unknown to you?
Could you solve this efficiently without using extra space?
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	nodePointer []*ListNode
}

// time complexity : O(n)
// space complexity: O(n)
func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano())
	sol := Solution{}
	for head != nil {
		sol.nodePointer = append(sol.nodePointer, head)
		head = head.Next
	}
	return sol
}

func (this *Solution) GetRandom() int {
	return this.nodePointer[rand.Intn(len(this.nodePointer))].Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
	head := makeLinkedList([]int{1, 2, 3})
	obj := Constructor(head)
	fmt.Println(obj.GetRandom())
}
