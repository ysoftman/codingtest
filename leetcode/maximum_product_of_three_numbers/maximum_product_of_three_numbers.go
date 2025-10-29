/*
https://leetcode.com/problems/maximum-product-of-three-numbers/
628. Maximum Product of Three Numbers
Easy

Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

Example 1:
Input: nums = [1,2,3]
Output: 6

Example 2:
Input: nums = [1,2,3,4]
Output: 24

Example 3:
Input: nums = [-1,-2,-3]
Output: -6

Constraints:
3 <= nums.length <= 104
-1000 <= nums[i] <= 1000
*/
package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	// fmt.Println(nums)
	// [-3,-2,-1,1] 처럼 음수*음수=양수 인상태에서 제일은 큰 마지막원소가 양수를 곱하면 가강큰 값이 되는 경우도 고려해야 한다.
	// -3 * -2 * 1 이 가장 큰 값
	temp1 := nums[0] * nums[1] * nums[len(nums)-1]
	temp2 := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	if temp1 > temp2 {
		return temp1
	}
	return temp2
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3}))
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-1, -2, -3}))
	fmt.Println(maximumProduct([]int{-3, -2, -1, 1}))
}
