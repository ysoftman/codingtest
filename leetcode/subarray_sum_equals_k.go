/*
https://leetcode.com/problems/subarray-sum-equals-k
560. Subarray Sum Equals K
Medium
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:
Input: nums = [1,1,1], k = 2
Output: 2

Example 2:
Input: nums = [1,2,3], k = 3
Output: 2

Constraints:
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/
package main

import "fmt"

/*
0부탁 각 index 까지 누적합
index 마다 index까지 누적합-k 한 값이 map 있으면 (찾은곳 부터 index까지 합이 k가 되기 때문) 찾아 있으면 value(count) 를 누적
index 까지 누적합을 맵으로 기록
*/
func subarraySum(nums []int, k int) int {
	result := 0
	m := make(map[int]int)
	// m[sum] = count
	// k - k = 0 으로 1개가 존재한다는 의미
	m[0] = 1
	sum := 0
	for _, n := range nums {
		sum += n
		if v, ok := m[sum-k]; ok {
			result += v
		}

		// 현재까지 누적값 맵에 기록하되, 기존 맵에 존재하는 카운트에 합산해 기록
		if v, ok := m[sum]; ok {
			m[sum] = v + 1
		} else {
			m[sum] = 1
		}
	}
	return result
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1, 3, 4, 2, -1, 2, 1, 1}, 2))
	fmt.Println(subarraySum([]int{2, 2, -2, 2, 2}, 2))
}
