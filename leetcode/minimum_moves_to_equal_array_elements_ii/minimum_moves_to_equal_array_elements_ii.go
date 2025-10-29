/*
https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/
462. Minimum Moves to Equal Array Elements II
Medium
Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.

In one move, you can increment or decrement an element of the array by 1.
Test cases are designed so that the answer will fit in a 32-bit integer.

Example 1:
Input: nums = [1,2,3]
Output: 2
Explanation:
Only two moves are needed (remember each move increments or decrements one element):
[1,2,3]  =>  [2,2,3]  =>  [2,2,2]

Example 2:
Input: nums = [1,10,2,9]
Output: 16

Constraints:
n == nums.length
1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/
package main

import (
	"fmt"
	"sort"
)

/*
[1,10] => 1+9 or 10-9 or 1+4,10-5  ... => 9 moves
[1,2,3] => 1+1,2,3-1 .. => 2 moves
[1,2,3,4] => 1+2,2+1,3,4-1 .. => 4 moves
... 배열 값들의 중간값(median)을 찾자
*/
func minMoves2(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	cnt := 0
	for left < right {
		cnt += nums[right] - nums[left]
		left++
		right--
	}
	return cnt
}

func main() {
	fmt.Println(minMoves2([]int{1, 2, 3}))
	fmt.Println(minMoves2([]int{1, 2, 3, 4}))
	fmt.Println(minMoves2([]int{1, 10, 2, 9}))
}
