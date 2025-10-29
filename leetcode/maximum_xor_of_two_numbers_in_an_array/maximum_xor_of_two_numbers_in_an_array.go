/*
https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
421. Maximum XOR of Two Numbers in an Array
Medium

Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 <= i <= j < n.

Example 1:
Input: nums = [3,10,5,25,2,8]
Output: 28
Explanation: The maximum result is 5 XOR 25 = 28.

Example 2:
Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
Output: 127

Constraints:
1 <= nums.length <= 2 * 105
0 <= nums[i] <= 231 - 1
*/
package main

import "fmt"

func findMaximumXOR(nums []int) int {
	max := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask |= 1 << i
		m := make(map[int]bool)
		for n := 0; n < len(nums); n++ {
			m[(nums[n] & mask)] = true
		}
		temp := max | (1 << i)
		for k := range m {
			if _, ok := m[k^temp]; ok {
				max = temp
				break
			}
		}
	}
	return max
}

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	fmt.Println(findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}))
}
