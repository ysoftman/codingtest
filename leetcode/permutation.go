/*
https://leetcode.com/problems/permutations/
46. Permutations
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]


                                      123
0번째고정할값들    (0<->0)123                 (0<->1)213                   (0<->2)312
1번째고정할값들 (1<->1)123 (1<->2)132     (1<->1)213 (1<->2)231      (1<->1)312 (1<->2)321
*/

package main

import "fmt"

func doPermute(nums []int, a, b int, result *[][]int) {
	if a == b {
		// fmt.Println("nums:", nums)
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		*result = append(*result, nums2)
		// fmt.Println("result:", result)
		return
	}
	for i := a; i <= b; i++ {
		nums[a], nums[i] = nums[i], nums[a]
		doPermute(nums, a+1, b, result)
		// 다음 루프에서 인덱스 순서가 제대로 동작하기 위해 원상 복구
		nums[a], nums[i] = nums[i], nums[a]
	}
}

func permute(nums []int) [][]int {
	result := [][]int{}
	doPermute(nums, 0, len(nums)-1, &result)
	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
