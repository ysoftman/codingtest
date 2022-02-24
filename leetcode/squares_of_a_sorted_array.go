/*
https://leetcode.com/problems/squares-of-a-sorted-array/
977. Squares of a Sorted Array
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.


Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?
*/
package main

import (
	"fmt"
	"sort"
)

// O(nlogn)
func sortedSquares2(nums []int) []int {
	// O(n)
	for i := range nums {
		nums[i] *= nums[i]
	}
	// O(nlogn)
	sort.Ints(nums)
	return nums
}

// O(n)
func sortedSquares(nums []int) []int {
	// find plus-start index
	// O(n)
	left, right := -1, -1
	for i, v := range nums {
		if v >= 0 {
			left = i - 1
			right = i
			break
		}
	}
	// nums are all plus
	if right == 0 {
		left = -1
		right = 0
	}
	// nums are all minus
	if right == -1 {
		left = len(nums) - 1
		right = len(nums)
	}

	// O(n)
	result := []int{}
	for left >= 0 || right <= len(nums)-1 {
		if left >= 0 && right <= len(nums)-1 {
			if nums[left]*nums[left] < nums[right]*nums[right] {
				result = append(result, nums[left]*nums[left])
				left--
			} else {
				result = append(result, nums[right]*nums[right])
				right++
			}
		} else if left >= 0 {
			result = append(result, nums[left]*nums[left])
			left--
		} else {
			result = append(result, nums[right]*nums[right])
			right++
		}

	}
	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-1}))
	fmt.Println(sortedSquares([]int{1}))
	fmt.Println(sortedSquares([]int{-8, -4, -1}))
	fmt.Println(sortedSquares([]int{1, 8, 10}))
}
