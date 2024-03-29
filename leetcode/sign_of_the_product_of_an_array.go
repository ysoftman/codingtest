/*
https://leetcode.com/problems/sign-of-the-product-of-an-array/
1822. Sign of the Product of an Array
Easy

There is a function signFunc(x) that returns:

1 if x is positive.
-1 if x is negative.
0 if x is equal to 0.
You are given an integer array nums. Let product be the product of all values in the array nums.

Return signFunc(product).

Example 1:
Input: nums = [-1,-2,-3,-4,3,2,1]
Output: 1
Explanation: The product of all values in the array is 144, and signFunc(144) = 1

Example 2:
Input: nums = [1,5,0,2,-3]
Output: 0
Explanation: The product of all values in the array is 0, and signFunc(0) = 0

Example 3:
Input: nums = [-1,1,-1,1,-1]
Output: -1
Explanation: The product of all values in the array is -1, and signFunc(-1) = -1

Constraints:
1 <= nums.length <= 1000
-100 <= nums[i] <= 100
*/
package main

import "fmt"

func arraySign(nums []int) int {
	if nums[0] < 0 {
		nums[0] = -1
	} else if nums[0] == 0 {
		nums[0] = 0
	} else {
		nums[0] = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			nums[0] *= -1
		} else if nums[i] == 0 {
			nums[0] *= 0
		}
	}
	return nums[0]
}

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(arraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
}
