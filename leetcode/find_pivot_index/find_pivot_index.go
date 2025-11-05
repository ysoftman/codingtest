/*
https://leetcode.com/problems/find-pivot-index/description
724. Find Pivot Index
Easy
Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.

Example 1:
Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11

Example 2:
Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.

Example 3:
Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0

Constraints:
1 <= nums.length <= 104
-1000 <= nums[i] <= 1000
*/
package main

import "fmt"

func pivotIndex(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	ar1 := []int{}
	ar2 := []int{}
	ar1 = append(ar1, nums[0])
	for i := 1; i < len(nums); i++ {
		ar1 = append(ar1, ar1[i-1]+nums[i])
	}
	ar2 = append(ar2, nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		ar2 = append(ar2, ar2[len(ar2)-1]+nums[i])
	}
	// reverse ar2
	for i, j := 0, len(ar2)-1; i < j; i, j = i+1, j-1 {
		ar2[i], ar2[j] = ar2[j], ar2[i]
	}

	// fmt.Println(ar1)
	// fmt.Println(ar2)
	pivot := -1
	for i := 0; i < len(ar1); i++ {
		if i+2 > len(ar2)-1 {
			break
		}
		if ar1[i] == ar2[i+2] {
			pivot = i + 1
			break
		}
	}

	if ar2[1] == 0 {
		return 0
	}
	if len(ar1) >= 2 && ar1[len(ar1)-2] == 0 {
		if pivot == -1 {
			return len(ar1) - 1
		}
		// leftmost pivot
		if pivot > len(ar1)-1 {
			return len(ar1) - 1
		}
	}
	return pivot
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, 0, 1}))
	fmt.Println(pivotIndex([]int{-1, -1, 0, 1, 1, 0}))
	fmt.Println(pivotIndex([]int{-1, -1, 1, 1, 0, 0}))
	fmt.Println(pivotIndex([]int{-1, -1, 0, 0, -1, -1}))
	fmt.Println(pivotIndex([]int{0}))
}
