/*
https://leetcode.com/problems/maximum-subarray/
53. Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23
*/
package main

import "fmt"

func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if max > nums[i] {
			max = nums[i]
		}
	}
	// time limit exceeded
	// temp := make(map[string]int)
	// for i:=0; i<len(nums); i++ {
	//     pre := 0
	//     for j:=i; j<len(nums); j++ {
	//         key := fmt.Sprintf("%v-v", i, j)
	//         temp[key] += pre + nums[j]
	//         if max < temp[key] {
	//             max = temp[key]
	//         }
	//         pre = temp[key]
	//     }
	// }

	temp := 0
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		if max < temp {
			max = temp
		}
		if temp < 0 {
			temp = 0
		}
	}
	return max
}

func max1(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// kadane's(카데인) algorithm
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	// max := 1<<31*-1
	max := dp[0]
	// O(N)
	for i := 1; i < len(nums); i++ {
		// 현재까 합과 현재값 중 큰 값을 선택하는 dynamic programming
		dp[i] = max1(nums[i], dp[i-1]+nums[i])
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
