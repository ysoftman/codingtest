/*
https://leetcode.com/problems/increasing-triplet-subsequence/
334. Increasing Triplet Subsequence
Medium

Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

Example 1:
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.

Example 2:
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.

Example 3:
Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.

Constraints:
1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1

Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?
*/
package main

import "fmt"

// time complexity: O(n)
// space complexity: O(1)
func increasingTriplet(nums []int) bool {
	min := 1<<31 - 1
	max := 1<<31 - 1
	for i := 0; i < len(nums); i++ {
		// nums[i] 값이 가장 작은것과 가장 큰것 사이에 있으면 된다.
		if nums[i] <= min {
			min = nums[i] // 현재까지중 가장 작은것
			// fmt.Println("min:", min)
		} else if nums[i] <= max {
			max = nums[i] // 현재까지중 가장 큰것
			// fmt.Println("max:", max)
		} else {
			return true // nums[i] 값이 가장 작은것과 가장 큰것 사이에 있으면 된다.
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	fmt.Println(increasingTriplet([]int{1, 1, 1, 1, 1}))
	fmt.Println(increasingTriplet([]int{1, 1, 1, 1, 1, 2, 3}))
}
