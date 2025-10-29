/*
https://leetcode.com/problems/maximum-erasure-value/
1695. Maximum Erasure Value
Medium
You are given an array of positive integers nums and want to erase a subarray containing unique elements. The score you get by erasing the subarray is equal to the sum of its elements.
Return the maximum score you can get by erasing exactly one subarray.
An array b is called to be a subarray of a if it forms a contiguous subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).

Example 1:
Input: nums = [4,2,4,5,6]
Output: 17
Explanation: The optimal subarray here is [2,4,5,6].

Example 2:
Input: nums = [5,2,1,2,5,2,1,2,5]
Output: 8
Explanation: The optimal subarray here is [5,2,1] or [1,2,5].

Constraints:
1 <= nums.length <= 105
1 <= nums[i] <= 104
*/
package main

import "fmt"

func findDuplicates(m *map[int]int, target int) bool {
	if _, ok := (*m)[target]; ok {
		return true
	}
	return false
}

// hashmap
func maximumUniqueSubarray(nums []int) int {
	m := make(map[int]int, len(nums))
	max_sum := 0
	sum := 0
	left, right := 0, 0
	for right < len(nums) {
		// hashmap 에 중복된것이 있다면, 현재 num[left] 를 제외시킨다.
		for findDuplicates(&m, nums[right]) {
			sum -= nums[left]
			delete(m, nums[left])
			left++
		}
		sum += nums[right]
		m[nums[right]] = 0
		right++
		if max_sum < sum {
			max_sum = sum
		}
	}
	return max_sum
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
}
