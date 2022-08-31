/*
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
448. Find All Numbers Disappeared in an Array
Easy
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Example 2:
Input: nums = [1,1]
Output: [2]

Constraints:
n == nums.length
1 <= n <= 105
1 <= nums[i] <= n

Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
*/
package main

import "fmt"

// time complexity: O(n)
// space complexity: O(n)
func findDisappearedNumbers2(nums []int) []int {
	m := make(map[int]bool, 0)
	r := []int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for i := 1; i <= len(nums); i++ {
		if _, exist := m[i]; !exist {
			r = append(r, i)
		}
	}
	return r
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// time complexity: O(n)
// space complexity: O(1)
func findDisappearedNumbers(nums []int) []int {
	// 존재하는 숫자는 index 의 값에 음수로 표시해 둔다.
	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] *= -1
		}
	}
	// fmt.Println(nums)
	r := []int{}
	// 양수 값의 index 만 없는 숫자다.
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			r = append(r, i+1)
		}
	}
	return r
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
