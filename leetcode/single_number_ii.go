/*
https://leetcode.com/problems/single-number-ii/
137. Single Number II
Medium
Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:
Input: nums = [2,2,3,2]
Output: 3

Example 2:
Input: nums = [0,1,0,1,0,1,99]
Output: 99

Constraints:
1 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
Each element in nums appears exactly three times except for one element which appears once.
*/
package main

import (
	"fmt"
	"sort"
)

// sorting(O(nlogn)) --> linear 아님
func singleNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] && nums[i+1] == nums[i+2] {
			i += 2
		} else {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

// O(N)
func singleNumber(nums []int) int {
	ones, twos, threes := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		// 두번 나타난 경우 체크를 위해 저장 => 한번 나타난 경우 & 현재 값
		twos |= ones & nums[i]

		// 한번 나타난 경우 체크를 위해 저장 => ones 에 현재값과 같은 값이 있다면 xor 연산에 의해 0
		ones ^= nums[i]

		// 세번 나타난 경우 => 한번 나타난 경우 & 두번 나타난 경우
		threes = ones & twos

		// 현재 값이 세번번 나타났다면 not threes 와 and 로 ones, twor 가 0 이된다.
		// 참고로 golang bitwise not (~) 이 따로 없어 xor 로  계산 not 을 구한다.
		// ones &= ~threes
		// twos &= ~threes
		ones &= (-1 ^ threes)
		twos &= (-1 ^ threes)
	}
	return ones
}
func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{30000, 500, 100, 30000, 100, 30000, 100}))
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
