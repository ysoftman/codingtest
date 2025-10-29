/*
https://leetcode.com/problems/move-zeroes/
283. Move Zeroes
Easy

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]
*/
package main

import "fmt"

func moveZeroes(nums []int) {
	emptyIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[emptyIdx] = nums[i]
			emptyIdx++
		}
	}
	for emptyIdx < len(nums) {
		nums[emptyIdx] = 0
		emptyIdx++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
