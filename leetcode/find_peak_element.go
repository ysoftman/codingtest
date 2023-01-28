/*
https://leetcode.com/problems/find-peak-element/
162. Find Peak Element
Medium
A peak element is an element that is strictly greater than its neighbors.

Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞.

You must write an algorithm that runs in O(log n) time.

Example 1:
Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.

Example 2:
Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

Constraints:
1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
nums[i] != nums[i + 1] for all valid i
*/
package main

import "fmt"

/*
-1 인덱스는 -무한대 값
[-무한대, 1, 2, 3, 4] 일때

nums[1] < nums[2] && nums[2] > nums[3]   => 2번째 인덱스 픽크

결과(픽크 되는 지점)는 여러개가 될 수 있음.
*/

// O(N)
func findPeakElement2(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

// binary seach, O(logN)
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if left == right {
			return left
		}
		mid := (right-left)/2 + left
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return len(nums) - 1
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
