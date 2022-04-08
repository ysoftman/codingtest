/*
https://leetcode.com/problems/subsets/
78. Subsets
Medium
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]

Constraints:
1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/
package main

import "fmt"

func recursiveSubsets(nums []int, idx int, subset []int, r *[][]int) {
	n := len(nums)
	if idx == n {
		sb := make([]int, len(subset))
		copy(sb, subset)
		*r = append(*r, sb)
		return
	}
	sb := make([]int, len(subset))
	copy(sb, subset)
	*r = append(*r, sb)

	for i := idx; i < n; i++ {
		recursiveSubsets(nums, i+1, append(subset, nums[i]), r)
	}
}
func subsets(nums []int) [][]int {
	r := [][]int{}
	subset := []int{}
	recursiveSubsets(nums, 0, subset, &r)
	return r
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
	fmt.Println(subsets([]int{1, 2, 3, 4, 8, 5, 7, 9}))
}
