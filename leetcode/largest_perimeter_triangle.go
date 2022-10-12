/*
https://leetcode.com/problems/largest-perimeter-triangle/
976. Largest Perimeter Triangle
Easy

Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return 0.

Example 1:
Input: nums = [2,1,2]
Output: 5

Example 2:
Input: nums = [1,2,1]
Output: 0

Constraints:
3 <= nums.length <= 104
1 <= nums[i] <= 106
*/
package main

import (
	"fmt"
	"sort"
)

// perimeter(둘레), 삼각형이 되려면 가장큰 c 보다 작은 a+b 가 커야 한다.
// a+b > c
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	// 가장 큰 수부터 3개씩 삼각형 둘레가 될 수 있는지 체크
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1}))
	fmt.Println(largestPerimeter([]int{1, 2, 1, 2, 3, 7, 1, 6, 4, 2, 6, 12, 9}))
}
