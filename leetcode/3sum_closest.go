/*
https://leetcode.com/problems/3sum-closest/
16. 3Sum Closest
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.

Example 1:
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

Example 2:
Input: nums = [0,0,0], target = 1
Output: 0


Constraints:

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/

package main

import (
	"fmt"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			temp := nums[i] + nums[left] + nums[right]
			if temp == target {
				return temp
			}
			// update closet
			if abs(temp-target) < abs(closest-target) {
				closest = temp
			}

			// find two-sum problem
			if temp < target {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
