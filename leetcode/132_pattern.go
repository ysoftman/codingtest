/*
https://leetcode.com/problems/132-pattern
456. 132 Pattern
Medium

Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

Return true if there is a 132 pattern in nums, otherwise, return false.


Example 1:
Input: nums = [1,2,3,4]
Output: false
Explanation: There is no 132 pattern in the sequence.

Example 2:
Input: nums = [3,1,4,2]
Output: true
Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

Example 3:
Input: nums = [-1,3,2,0]
Output: true
Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].


Constraints:

n == nums.length
1 <= n <= 2 * 105
-109 <= nums[i] <= 109
*/

package main

import "fmt"

// time complexity O(n^2) --> time limit exceeded!!!
// func find132pattern(nums []int) bool {
// 	for j := 1; j < len(nums)-1; j++ {
// 		right := 1<<32*-1
// 		for k := j + 1; k < len(nums); k++ {
// 			if nums[k] < nums[j] && right < nums[k] {
// 				right = nums[k]
// 			}
// 		}
// 		for i := j - 1; i >= 0; i-- {
// 			if nums[i] < right {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// using stack, time complexity O(n)
func find132pattern(nums []int) bool {
	minArr := make([]int, len(nums))
	min := 1<<32 - 1

	// 이후 nums[i]<nums[j] 조건 만족 체크를 위해
	// 현재 위치에서 가장 작은 값을 기록 해 둔다.
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
		minArr[i] = min
	}

	// nums[k] 가 될 수 있는 후보들을 스택으로 기록
	stack := []int{}
	for j := len(nums) - 1; j > 0; j-- {
		// nums[i]<nums[j] 조건 체크
		if minArr[j] < nums[j] {
			// min 보다 작으면 (왼쪽으로 갈수로 min 이 커지기 때문에) 스택에서 빼서(pop) 탐색 시간을 줄인다.
			for len(stack) > 0 && stack[len(stack)-1] <= minArr[j] {
				stack = stack[:len(stack)-1]
			}
			// nums[j]<nums[k]
			if len(stack) > 0 && stack[len(stack)-1] < nums[j] {
				// fmt.Println(minArr[j], nums[j], stack[len(stack)-1])
				return true
			}
			// nums[k] 를 스택으로 기록해 둔다.
			stack = append(stack, nums[j])
		}
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))
	fmt.Println(find132pattern([]int{-2, 1, 1}))
	fmt.Println(find132pattern([]int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(find132pattern([]int{1, 0, 1, -4, -3}))
}
