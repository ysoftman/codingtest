/*
https://leetcode.com/problems/contains-duplicate-iii/
220. Contains Duplicate III
Medium
Given an integer array nums and two integers k and t, return true if there are two distinct indices i and j in the array such that abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.

Example 1:
Input: nums = [1,2,3,1], k = 3, t = 0
Output: true

Example 2:
Input: nums = [1,0,1,1], k = 1, t = 2
Output: true

Example 3:
Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false

Constraints:
1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
0 <= k <= 104
0 <= t <= 231 - 1
*/
package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// brute-force, accepted
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= t && abs(i-j) <= k {
				return true
			}
		}
	}
	return false
}

// bucket approach, accepted
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	window := t + 1
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		bucket_id := nums[i] / window
		if nums[i] < 0 {
			bucket_id -= 1
		}
		if _, exist := m[bucket_id]; exist {
			return true
		}
		if _, exist := m[bucket_id-1]; exist {
			if abs(m[bucket_id-1]-nums[i]) < window {
				return true
			}
		}
		if _, exist := m[bucket_id+1]; exist {
			if abs(m[bucket_id+1]-nums[i]) < window {
				return true
			}
		}
		m[bucket_id] = nums[i]
		if i >= k {
			delete(m, nums[i-k]/window)
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	fmt.Println(containsNearbyAlmostDuplicate([]int{2147483647, -1, 2147483647}, 1, 2147483647))
}
