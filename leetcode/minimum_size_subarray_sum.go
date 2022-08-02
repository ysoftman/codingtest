/*
https://leetcode.com/problems/minimum-size-subarray-sum/
209. Minimum Size Subarray Sum
Medium
Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:
Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

package main

import "fmt"

/*
taget : 7
2,3,1,2,4,3

2,3,1,2 >= 7 found! then l++
l     r

  3,1,2  not found! then r++
  l   r

  3,1,2,4 >= 7 found! then l++
  l     r

    1,2,4 >= 7 found! then l++
    l   r

      2,4 not found! then r++
      l r

      2,4,3 >= 7 found! then l++
      l   r

        4,3 >= 7 found! then l++
        l r

          3 not found! if end of array, exit loop
         l/r
*/

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] > target {
			return 1
		}
		return 0
	}
	// l, r 두 지점으로 sliding window 로 O(n)으로 가장 작은 minsubarraylen 을 찾는다.
	l, r := 0, 0
	min := 1<<31 - 1
	sum := nums[r]
	for l != len(nums) && r != len(nums) {
		if sum >= target {
			// fmt.Println(nums[l:r+1], sum)
			if min > (r-l)+1 {
				min = (r - l) + 1
			}
			sum -= nums[l]
			l++
		} else {
			r++
			if r < len(nums) {
				sum += nums[r]
			}
		}
	}
	if min == 1<<31-1 {
		min = 0
	}
	return min
}

func main() {
	fmt.Println(minSubArrayLen(6, []int{10, 2, 3}))
	fmt.Println(minSubArrayLen(7, []int{5}))
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
