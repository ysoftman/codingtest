/*
https://leetcode.com/problems/longest-consecutive-sequence/
128. Longest Consecutive Sequence
Medium
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/
package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// using sort: O(NlogN)
func longestConsecutive2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	sort.Ints(nums)
	// fmt.Println(nums)
	maxLC, lc := 0, 1
	for i := 1; i < len(nums); i++ {
		// 0,1,1,2 처럼 같은 숫자가 나오는 경우 카운트 유지하고 스킵
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1]+1 == nums[i] {
			lc++
		} else {
			maxLC = max(maxLC, lc)
			lc = 1
		}
	}
	maxLC = max(maxLC, lc)
	return maxLC
}

func findNum(m map[int]bool, target int) bool {
	if _, ok := m[target]; ok {
		return true
	}
	return false
}

// using hashmap: O(N)
func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	maxLC, lc := 0, 0
	for i := 0; i < len(nums); i++ {
		// 현재값-1 이 존재하면 그 작은 값으로 나중에 해시맵에서 찾게되니 스킵한다.
		// 이 skip 조건이 없으면 시간 초과.
		if findNum(m, nums[i]-1) {
			continue
		}
		lc++
		temp := nums[i] + 1
		for findNum(m, temp) {
			lc++
			temp++
		}
		maxLC = max(maxLC, lc)
		lc = 0
	}
	return maxLC
}

func main() {
	fmt.Println("approach 1")
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{0}))
	fmt.Println(longestConsecutive([]int{0}))
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))

	fmt.Println("approach 2")
	fmt.Println(longestConsecutive2([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive2([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive2([]int{0}))
	fmt.Println(longestConsecutive2([]int{0}))
	fmt.Println(longestConsecutive2([]int{1, 2, 0, 1}))
}
