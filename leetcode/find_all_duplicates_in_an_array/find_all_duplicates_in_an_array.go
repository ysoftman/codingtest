/*
https://leetcode.com/problems/find-all-duplicates-in-an-array/
442. Find All Duplicates in an Array
Medium
Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.

You must write an algorithm that runs in O(n) time and uses only constant extra space.

Example 1:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]

Example 2:
Input: nums = [1,1,2]
Output: [1]

Example 3:
Input: nums = [1]
Output: []

Constraints:
n == nums.length
1 <= n <= 105
1 <= nums[i] <= n
Each element in nums appears once or twice.
*/

package main

import "fmt"

// time complexity : O(n)
// space complexity : O(n)
func findDuplicates2(nums []int) []int {
	twice := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		twice[nums[i]]++
	}
	r := []int{}
	for k, v := range twice {
		if v >= 2 {
			r = append(r, k)
		}
	}
	return r
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// time complexity : O(n)
// space complexity : O(1)
func findDuplicates(nums []int) []int {
	r := []int{}
	for i := 0; i < len(nums); i++ {
		// 값을 0부터 시작하는 index 로해서 1~n 까지의 출현 여부 flag(음수)로 사용한다.
		index := abs(nums[i]) - 1
		// index 위치의 값을 음수 표시를 해두어 index 값이 한번 나타났음을 표시
		// 이게 가능한 이유는 1~n 내의 숫자만 있고 같은 수가 최대 2번까지만 나올 수 있기때문이다.
		if nums[index] < 0 {
			r = append(r, index+1)
		}
		nums[index] = -nums[index]
	}
	// fmt.Println(nums)
	return r
}

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates([]int{1, 1, 2}))
	fmt.Println(findDuplicates([]int{1}))
}
