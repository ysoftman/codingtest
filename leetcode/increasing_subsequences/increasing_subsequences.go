/*
https://leetcode.com/problems/non-decreasing-subsequences/
491. Non-decreasing Subsequences
Medium
Given an integer array nums, return all the different possible non-decreasing subsequences of the given array with at least two elements. You may return the answer in any order.

Example 1:
Input: nums = [4,6,7,7]
Output: [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

Example 2:
Input: nums = [4,4,3,2,1]
Output: [[4,4]]

Constraints:
1 <= nums.length <= 15
-100 <= nums[i] <= 100
*/
package main

import (
	"fmt"
)

func makeKey(n []int) string {
	key := ""
	for i := 0; i < len(n); i++ {
		key += fmt.Sprintf("%v-", n[i])
	}
	return key
}
func recursive(nums []int, idx int, temp *[]int, r *map[string][]int) {
	if len(*temp) >= 2 {
		curResult := make([]int, len(*temp))
		copy(curResult, *temp)
		// 4 6 7(2번째)
		// 4 6 7(3번째) 로 중복저장을 피하기 위해 map에 저장
		key := makeKey(curResult)
		// fmt.Println("key", key)
		(*r)[key] = curResult
	}
	for i := idx; i < len(nums); i++ {
		// temp 에 처음 값을 추가 하는 경우
		// temp 마지막 원소가 현재 nums[i] 작을때 추가하는 경우
		if len(*temp) == 0 || len(*temp) > 0 && (*temp)[len(*temp)-1] <= nums[i] {
			temp2 := append(*temp, nums[i])
			recursive(nums, i+1, &temp2, r)
		}
	}
}
func findSubsequences(nums []int) [][]int {
	// 중복 제거를 위해 map 사용(golang 엔 set이 없으니..)
	m := make(map[string][]int)
	recursive(nums, 0, &[]int{}, &m)
	r := [][]int{}
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(findSubsequences([]int{4, 4, 3, 2, 1}))
}
