/*
https://leetcode.com/problems/minimum-increment-to-make-array-unique
945. Minimum Increment to Make Array Unique
Medium
You are given an integer array nums. In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.

Return the minimum number of moves to make every value in nums unique.

The test cases are generated so that the answer fits in a 32-bit integer.

Example 1:
Input: nums = [1,2,2]
Output: 1
Explanation: After 1 move, the array could be [1, 2, 3].

Example 2:
Input: nums = [3,2,1,2,1,7]
Output: 6
Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
It can be shown that it is impossible for the array to have all unique values with 5 or less moves.

Constraints:
1 <= nums.length <= 105
0 <= nums[i] <= 105
*/
package main

import (
	"fmt"
	"sort"
)

// 중복된 원소들만 변경해야 된다면
func minIncrementForUnique2(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	sort.Ints(nums)
	moves := 0
	candi := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			continue
		}
		if candi < nums[i] {
			candi = nums[i] + 1
		}
		for {
			if _, ok := m[candi]; ok {
				candi++
			} else {
				break
			}
		}
		moves += candi - nums[i]
		m[candi]++
		m[nums[i]]--
		candi++
	}
	return moves
}

// no hashmap
// 모든 원소의 값이 변경되도 상관 없다면
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	moves := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			moves += (nums[i-1] + 1) - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return moves
}

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
	fmt.Println(minIncrementForUnique([]int{0, 0}))
	fmt.Println(minIncrementForUnique([]int{2, 2, 2, 2, 0}))
	fmt.Println(minIncrementForUnique([]int{1, 10, 10}))
}
