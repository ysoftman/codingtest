/*
https://leetcode.com/problems/set-mismatch/
645. Set Mismatch
Easy

You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.

Example 1:
Input: nums = [1,2,2,4]
Output: [2,3]

Example 2:
Input: nums = [1,1]
Output: [1,2]

Constraints:
2 <= nums.length <= 104
1 <= nums[i] <= 104
*/
package main

import (
	"fmt"
	"sort"
)

// missing -> 1~n 합산에서 찾기
func findErrorNums2(nums []int) []int {
	sort.Ints(nums)
	// fmt.Println(nums)
	duplicated := 0
	missing := 0
	sum1 := 1
	sum2 := nums[0]
	for i := 1; i < len(nums); i++ {
		sum1 += i + 1
		if nums[i] == nums[i-1] {
			duplicated = nums[i]
			continue
		}
		sum2 += nums[i]
	}
	missing = sum1 - sum2
	return []int{duplicated, missing}
}

// missing -> 이전 숫자와 차이가 1이상 나는것으로 찾기
func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	// fmt.Println(nums)
	duplicated := 0

	// 기본 1로 설정
	missing := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			duplicated = nums[i]
		} else if nums[i] > nums[i-1]+1 {
			missing = nums[i-1] + 1
		}
	}
	// missing == 1 일때 마지막이 빠진것인지 체크해야 한다.
	// 마지막 숫자가 없는 경우 체크
	if nums[len(nums)-1] != len(nums) {
		missing = len(nums)
	}

	return []int{duplicated, missing}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
	fmt.Println(findErrorNums([]int{3, 2, 3, 4, 6, 5}))
}
