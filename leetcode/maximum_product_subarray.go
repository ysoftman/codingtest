/*
https://leetcode.com/problems/maximum-product-subarray/
152. Maximum Product Subarray
Medium
Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Constraints:
1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
ex) 3, -1, 4

앞쪽에서 곱하는 경
3=3
3*-1=3
3*-1*4=-12
maxleft = 3

뒤쪽에서 곱하는경우
4=4
4*-1=-4
4*-1*3=-12
maxright =4

앞,뒤쪽 중 큰쪽이 결과가 된다.
max(maxleft, maxright)
*/
func maxProduct(nums []int) int {
	result := nums[0]
	tempLeft := 0
	tempRight := 0
	for i := 0; i < len(nums); i++ {
		if tempLeft == 0 {
			tempLeft = nums[i]
		} else {
			tempLeft *= nums[i]
		}
		if tempRight == 0 {
			tempRight = nums[len(nums)-1-i]
		} else {
			tempRight *= nums[len(nums)-1-i]
		}
		// fmt.Println("tempLeft:", tempLeft)
		// fmt.Println("tempRight:", tempRight)
		result = max(result, max(tempLeft, tempRight))
	}
	return result
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{3, -1, 4}))
	fmt.Println(maxProduct([]int{3, -1, 5, -2, 4}))
}
