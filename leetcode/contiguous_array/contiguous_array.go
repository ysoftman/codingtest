/*
https://leetcode.com/problems/contiguous-array
525. Contiguous Array
Medium
Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.

Example 1:
Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

Example 2:
Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

Example 3:
Input: nums = [0,1,1,1,1,1,0,0,0]
Output: 6
Explanation: [1,1,1,0,0,0] is the longest contiguous subarray with equal number of 0 and 1.

Constraints:
1 <= nums.length <= 105
nums[i] is either 0 or 1.
*/
package main

import "fmt"

// brute force : O(n^2), time limit exceeded
// func findMaxLength(nums []int) int {
// 	max := 0
// 	for i := 0; i < len(nums); i++ {
// 		sum := 0
// 		for j := i; j < len(nums); j++ {
// 			if nums[j] == 0 {
// 				sum += -1
// 			} else {
// 				sum += nums[j]
// 			}
// 			if sum == 0 {
// 				if max < (j-i)+1 {
// 					max = (j - i) + 1
// 				}
// 			}
// 		}
// 	}
// 	return max
// }

func findMaxLength(nums []int) int {
	max := 0
	m := make(map[int]int)
	m[0] = -1 // [0,1] -> 인덱스에서의 sum [-1, 0] 로 0일때의 길이 파악을 위해 m[0]=-1 로 설정해 1-(-1)=2 의 길이로 찾을 수 있도록 한다.
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum += -1
		} else {
			sum += 1
		}

		if index, ok := m[sum]; ok {
			if max < i-index {
				max = i - index
			}
		} else {
			m[sum] = i
		}
	}

	return max
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
	fmt.Println(findMaxLength([]int{0, 1, 1, 1, 1, 1, 0, 0, 0}))
	fmt.Println(findMaxLength([]int{0, 1, 1, 1, 1, 1, 0, 0, 0, 0}))
}
