/*
https://leetcode.com/problems/arithmetic-slices/
413. Arithmetic Slices
Medium

An integer array is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.

For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.
Given an integer array nums, return the number of arithmetic subarrays of nums.

A subarray is a contiguous subsequence of the array.

Example 1:
Input: nums = [1,2,3,4]
Output: 3
Explanation: We have 3 arithmetic slices in nums: [1, 2, 3], [2, 3, 4] and [1,2,3,4] itself.

Example 2:
Input: nums = [1]
Output: 0

Constraints:
1 <= nums.length <= 5000
-1000 <= nums[i] <= 1000
*/
package main

import "fmt"

// arithmetic slice 조건
// 최소 3개의 원소가 있어야 한다.
// 두개의 연속된 원소간의 차이가 같아야 한다.
// --> 1,2 의 차이(1) == 2,3 의 차이(1) => 1,2,3
// --> 1,2 의 차이(1) == 2,3 의 차이(1) == 3,4 의 차이(1) 같음 => 1,2,3,4
// --> 2,3 의 차이(1) == 3,45의 차이(1) 이 같음 => 2,3,4
// time complexity: O(m*n)
// space complexity: O(1)
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	cnt := 0
	// 3 개의 연속된 원소 차이 체크
	for i := 0; i < len(nums)-2; i++ {
		diff := nums[i] - nums[i+1]
		for j := i + 2; j < len(nums); j++ {
			if nums[j-1]-nums[j] == diff {
				cnt++
				// fmt.Println(nums[i : j+1])
			} else {
				break
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println(numberOfArithmeticSlices([]int{1}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 3, 5, 7, 9}))
	fmt.Println(numberOfArithmeticSlices([]int{7, 7, 7, 7}))
	fmt.Println(numberOfArithmeticSlices([]int{3, -1, -5, -9}))
}
