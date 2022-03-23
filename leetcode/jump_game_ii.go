/*
https://leetcode.com/problems/jump-game-ii/
45. Jump Game II
Medium
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
You can assume that you can always reach the last index.

Example 1:
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [2,3,0,1,4]
Output: 2

Constraints:
1 <= nums.length <= 104
0 <= nums[i] <= 1000
*/

package main

import "fmt"

// time limit exceeded
func recursiveJump(nums []int, idx int, stack []int, minJump *int) {
	if len(nums) <= idx {
		return
	}
	if len(nums)-1 == idx {
		if *minJump > len(stack) {
			// fmt.Println("stack:", stack)
			*minJump = len(stack)
			return
		}
	}
	for i := 1; i <= nums[idx]; i++ {
		recursiveJump(nums, idx+i, append(stack, i), minJump)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func jump(nums []int) int {
	// stack := []int{}
	// minJump := 1<<31-1
	// recursiveJump(nums, 0, stack, &minJump)
	// return minJump

	dp := make([]int, len(nums))
	// dynamic progrmming 방법으로 풀자.
	// 뒤쪽인덱스 부터 마지막 인덱스로 갈 최소의 점프수를 memoization 한다.
	for i := len(nums) - 2; i >= 0; i-- {
		minJump := 1<<31 - 1
		for j := nums[i]; j > 0; j-- {
			if i+j >= len(nums) {
				continue
			}
			if minJump > dp[i+j] {
				minJump = dp[i+j]
			}
		}
		dp[i] = minJump + 1
	}
	// fmt.Println("dp:",dp)
	return dp[0]
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{2, 1, 5, 8, 7, 5, 2, 50, 7, 5, 2, 2, 2, 3, 3, 3, 3, 4, 1, 1, 1, 5, 1, 1, 8, 2, 3, 6, 200, 5, 2, 1, 1, 1, 1, 1, 1, 1, 3, 0, 1, 4}))
}
