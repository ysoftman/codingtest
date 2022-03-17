/*
https://leetcode.com/problems/4sum/
18. 4Sum
Medium
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]

Constraints:
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	r := [][]int{}
	sort.Ints(nums)
	// fmt.Println("nums:", nums)
	for i := 0; i < len(nums); i++ {
		// skip duplicate num, ex) [2,2,2,2,2]
		if i-1 >= 0 && nums[i-1] == nums[i] {
			continue
		}
		// skip duplicate num, ex) [2,2,2,2,2]
		for j := len(nums) - 1; j >= 0; j-- {
			// skip duplicate num, ex) [2,2,2,2,2]
			if j+1 <= len(nums)-1 && nums[j+1] == nums[j] {
				continue
			}
			subTarget := target - (nums[i] + nums[j])
			// twosum problem
			left, right := i+1, j-1
			for left < right {
				// skip duplicate num, ex) [2,2,2,2,2]
				if left-1 >= i+1 && nums[left-1] == nums[left] {
					left++
					continue
				}
				// skip duplicate num, ex) [2,2,2,2,2]
				if right+1 <= j-1 && nums[right+1] == nums[right] {
					right--
					continue
				}
				twosum := nums[left] + nums[right]
				if twosum == subTarget {
					r = append(r, []int{nums[i], nums[j], nums[left], nums[right]})
				}
				if nums[left]+nums[right] > subTarget {
					right--
				} else {
					left++
				}
			}
		}
	}
	return r
}
func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
