/*
https://leetcode.com/problems/subsets-ii/
90. Subsets II
Medium
Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Example 2:
Input: nums = [0]
Output: [[],[0]]

Constraints:
1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/
package main

import (
	"fmt"
	"sort"
)

func recursiveSubsetsWithDuplicate(nums []int, startIdx int, subset []int, result *[][]int) {
	if startIdx > len(nums) {
		return
	}
	s := make([]int, len(subset))
	copy(s, subset)
	*result = append(*result, s)

	// 중복 제거한 power set(멱집합, 주어진 집합의 모든 부분 집합들로 구성된 집합이다.)
	for i := startIdx; i < len(nums); i++ {
		// skip duplicate subset
		// 하지만 1,2,2 처럼 처음 subset 길이가 커질때는 스킵하면 안된다.
		if i > startIdx && nums[i-1] == nums[i] {
			continue
		}
		recursiveSubsetsWithDuplicate(nums, i+1, append(subset, nums[i]), result)
	}
}
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	r := [][]int{}
	subset := []int{}
	recursiveSubsetsWithDuplicate(nums, 0, subset, &r)
	return r
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 5, 1, 2}))
}
