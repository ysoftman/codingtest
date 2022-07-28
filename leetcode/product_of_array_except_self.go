/*
https://leetcode.com/problems/product-of-array-except-self/
238. Product of Array Except Self
Medium
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:
2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/
package main

import "fmt"

/*
1,2,3,4 의 경우

step1
i=i-1까지 누적곱값을 설정
1,1,2,6

step2
마지막 원소부터 0까지
i-1=(마지막에서 i까지 누적곱을 곱한다.
1*24,1*12,2*4,6*1
*/
// without division, time complexity:O(n), space complexity:O(1)
func productExceptSelf(nums []int) []int {
	r := make([]int, len(nums))
	temp := 1
	for i := 0; i < len(nums); i++ {
		r[i] = temp
		temp *= nums[i]
	}
	// fmt.Println("step1:", r)
	temp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		r[i] *= temp
		temp *= nums[i]
	}
	// fmt.Println("step2:", r)
	return r
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
