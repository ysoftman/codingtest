/*
https://leetcode.com/problems/jump-game/
55. Jump Game
Medium
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
Return true if you can reach the last index, or false otherwise.


Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

package main

import "fmt"

/*
greedy 방식으로 뒤에서 부터 마지막인덱스(last index)를 0번째 인덱스로 변경될 수 있는지 확인

2,3,1,1,4
0 1 2 3 4 => index

index 3 -> jump1 -> index 4, last index = 3
index 2 -> jump1 -> index 3, last index = 2
index 1 -> jump1 -> index 2 or jump2 -> index 3 or jump3 -> index 4, last index = 1
index 0 -> jump1 -> index 1 or jump2 -> index 2, last index = 0
*/
func canJump(nums []int) bool {
	lastindex := len(nums) - 1
	// O(N) linear time complexity
	for i := len(nums) - 1; i >= 0; i-- {
		// update lastindex
		if i+nums[i] >= lastindex {
			lastindex = i
		}
	}
	if lastindex == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
