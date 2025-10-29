/*
https://leetcode.com/problems/sort-array-by-parity/
905. Sort Array By Parity
Easy
Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
Return any array that satisfies this condition.

Example 1:
Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Example 2:
Input: nums = [0]
Output: [0]


Constraints:
1 <= nums.length <= 5000
0 <= nums[i] <= 5000
*/
package main

import (
	"fmt"
	"sort"
)

// even, odd 두 저장공간으로 분리하는 방법
func sortArrayByParity3(nums []int) []int {
	even := []int{}
	odd := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	return append(even, odd...)
}

// even 값만 찾는 루프, odd 값만 찾는 루프 방법
func sortArrayByParity2(nums []int) []int {
	r := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			r = append(r, nums[i])
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			r = append(r, nums[i])
		}
	}
	return r
}

// 왼쪽 odd, 오른쪽 even 찾아서 스왑하는 방법
func sortArrayByParity1(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		for nums[left]%2 == 0 {
			left++
			if left == len(nums)-1 {
				break
			}
		}
		for nums[right]%2 == 1 {
			right--
			if right == 0 {
				break
			}
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}

// odd, even 판단해 정렬하는 방법
func sortArrayByParity(nums []int) []int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a]%2 < nums[b]%2
	})
	return nums
}
func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0, 2}))
	fmt.Println(sortArrayByParity([]int{0, 1}))
}
