/*
https://leetcode.com/problems/find-the-duplicate-number/
287. Find the Duplicate Number
Medium
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
There is only one repeated number in nums, return this repeated number.
You must solve the problem without modifying the array nums and uses only constant extra space.

Example 1:
Input: nums = [1,3,4,2,2]
Output: 2

Example 2:
Input: nums = [3,1,3,4,2]
Output: 3
*/
package main

import "fmt"

// import "sort"

func findDuplicate(nums []int) int {
	// O(NLogN) + O(N) => O(NLogN)
	// sort.Ints(nums)
	// for i := 1; i<len(nums); i++ {
	//     if nums[i] == nums[i-1] {
	//         return nums[i]
	//     }
	// }

	// without sorting
	// 시간복잡도: O(N), 공간복잡도 O(N) 으로 효율적이지만 nums 의 값들이 nums 인덱스 범위를 넘어서면 안되고, nums 값을 -로 마킹하게 된다.
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = nums[i]
		if temp < 0 {
			temp *= -1
		}

		// 해당 인덱스의 값이 음수라면 인덱스(값)은 이미 있는것으로 중복
		if nums[temp-1] < 0 {
			return temp
		}
		// nums[인덱스=값으로취급] = 탐색된 경우 음수값으로 설정
		nums[temp-1] *= -1
	}
	return -1
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
