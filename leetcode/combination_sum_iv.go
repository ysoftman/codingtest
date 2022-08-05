/*
https://leetcode.com/problems/combination-sum-iv/
377. Combination Sum IV
Medium
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.

The test cases are generated so that the answer can fit in a 32-bit integer.

Example 1:
Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.

Example 2:
Input: nums = [9], target = 3
Output: 0

Constraints:
1 <= nums.length <= 200
1 <= nums[i] <= 1000
All the elements of nums are unique.
1 <= target <= 1000

Follow up: What if negative numbers are allowed in the given array? How does it change the problem? What limitation we need to add to the question to allow negative numbers?
*/
package main

import (
	"fmt"
)

// dynamic programming 을 사용하지 않으면 time limit exceeded
func combsum4(nums []int, target int, dp *[]int, debug []int) int {
	// dp 에 cnt 값이 기록된적이 있으면 이를 사용하면 된다.
	if (*dp)[target] != -1 {
		return (*dp)[target]
	}
	if target == 0 {
		// fmt.Println(debug)
		return 1
	}
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if target-nums[i] >= 0 {
			cnt += combsum4(nums, target-nums[i], dp, append(debug, nums[i]))
		}
	}
	(*dp)[target] = cnt
	return cnt
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	// 초기값 설정, -1 이면 값이 저장된것이 아님
	for i := 0; i < target+1; i++ {
		dp[i] = -1
	}
	dp[0] = 1
	debug := []int{}
	return combsum4(nums, target, &dp, debug)
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4([]int{9}, 3))
	fmt.Println(combinationSum4([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 20))
	fmt.Println(combinationSum4([]int{3, 1, 2, 4}, 4))
	fmt.Println(combinationSum4([]int{1, 2, 3}, 32))
	fmt.Println(combinationSum4([]int{3, 1, 2, 4}, 4))
	fmt.Println(combinationSum4([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700, 710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980, 990, 111}, 999))
}
