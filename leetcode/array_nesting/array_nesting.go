/*
https://leetcode.com/problems/array-nesting/
565. Array Nesting
Medium
You are given an integer array nums of length n where nums is a permutation of the numbers in the range [0, n - 1].

You should build a set s[k] = {nums[k], nums[nums[k]], nums[nums[nums[k]]], ... } subjected to the following rule:

The first element in s[k] starts with the selection of the element nums[k] of index = k.
The next element in s[k] should be nums[nums[k]], and then nums[nums[nums[k]]], and so on.
We stop adding right before a duplicate element occurs in s[k].
Return the longest length of a set s[k].

Example 1:
Input: nums = [5,4,0,3,1,6,2]
Output: 4
Explanation:
nums[0] = 5, nums[1] = 4, nums[2] = 0, nums[3] = 3, nums[4] = 1, nums[5] = 6, nums[6] = 2.
One of the longest sets s[k]:
s[0] = {nums[0], nums[5], nums[6], nums[2]} = {5, 6, 2, 0}

Example 2:
Input: nums = [0,1,2]
Output: 1

Constraints:
1 <= nums.length <= 105
0 <= nums[i] < nums.length
All the values of nums are unique.
*/
package main

import "fmt"

// brute force, time limit exceeded
// time complexity: O(n*n)
func arrayNesting2(nums []int) int {
	maxCnt := 0
	for i := 0; i < len(nums); i++ {
		// start 값이 되기전까지 루프 카운트
		start := nums[nums[i]]
		cnt := 1
		for start != nums[i] {
			start = nums[start]
			cnt++
		}
		if maxCnt < cnt {
			maxCnt = cnt
		}
	}
	return maxCnt
}

// 이미 방문했던 곳은 스킵
// time complexity:  O(n)
func arrayNesting(nums []int) int {
	maxCnt := 0
	visited := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		// start 값이 되기전까지 루프 카운트
		start := nums[nums[i]]
		cnt := 1
		// 이미 방문했다면 이후 연결들은 이미 방문했던곳들이라, 스킵한다.
		if visited[start] {
			continue
		}
		for start != nums[i] {
			start = nums[start]
			cnt++
			visited[start] = true
		}
		if maxCnt < cnt {
			maxCnt = cnt
		}
	}
	return maxCnt
}

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
	fmt.Println(arrayNesting([]int{0, 1, 2}))
}
