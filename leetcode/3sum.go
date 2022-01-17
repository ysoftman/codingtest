/*
https://leetcode.com/problems/3sum/
15. 3Sum
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Example 2:
Input: nums = []
Output: []

Example 3:
Input: nums = [0]
Output: []

*/

package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) [][]int {
	// fmt.Println(">", target, ">>", nums)
	result := [][]int{}
	m := make(map[int]int, 0)
	for i := range nums {
		if val, exist := m[target-nums[i]]; exist {
			r := []int{}
			r = append(r, target*(-1), nums[i], val)
			result = append(result, r)
		}
		m[nums[i]] = nums[i]
	}
	return result
}

func compare(v, r []int) bool {
	for i := 0; i < 3; i++ {
		if v[i] != r[i] {
			return false
		}
	}
	return true
}
func skipDuplicate(results [][]int, r []int) bool {
	for _, v := range results {
		if compare(v, r) {
			return true
		}
	}
	return false
}
func threeSum(nums []int) [][]int {
	results := [][]int{}
	if len(nums) < 3 {
		return results
	}

	for i, v := range nums {
		// 현재 i번째 값을 만들 수 있는 twoSum 문제로 풀이
		temp := []int{}
		temp = append(temp, nums[:i]...)
		if i+1 < len(nums) {
			temp = append(temp, nums[i+1:]...)
		}

		ts := [][]int{}
		ts = twoSum(temp, v*(-1))

		for i := 0; i < len(ts); i++ {
			sort.Ints(ts[i])
			if skipDuplicate(results, ts[i]) {
				continue
			}

			results = append(results, ts[i])
		}
	}
	return results
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
