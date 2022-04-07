/*
https://leetcode.com/problems/sort-colors/
75. Sort Colors
Medium
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]

Constraints:
n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.
*/
package main

import "fmt"

// bubble sort
func sortColors(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		changed := false
		for j := 0; j < n-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				changed = true
			}
		}
		if !changed {
			return
		}
	}
}

func sortByPivot(nums []int, left, right *int, pivot int) {
	for *left < *right {
		if nums[*left] == pivot {
			*left++
			continue
		}
		if nums[*right] != pivot {
			*right--
			continue
		}
		nums[*left], nums[*right] = nums[*right], nums[*left]
	}
}
func sortColors2(nums []int) {
	n := len(nums)
	left := 0
	right := n - 1
	// 0~끝 0 과 1 이상으로 정렬
	sortByPivot(nums, &left, &right, 0)
	right = n - 1
	// left ~ 끝, 1 과 2로 정렬
	sortByPivot(nums, &left, &right, 1)
}

// Quick Sort
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	start := left
	end := right
	pivot := nums[end]
	for left < right {
		if nums[left] <= pivot {
			left++
			continue
		}
		if nums[right] >= pivot {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[end] = nums[end], nums[left]

	quickSort(nums, start, left-1)
	quickSort(nums, left+1, end)
}
func sortColors3(nums []int) {
	n := len(nums)
	left := 0
	right := n - 1
	quickSort(nums, left, right)
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
	arr = []int{2, 0, 1}
	sortColors(arr)
	fmt.Println(arr)
	arr = []int{2, 0, 2, 1, 1, 0, 2, 0, 1, 2, 1, 2, 1, 1, 2, 0, 1, 0, 2, 1}
	sortColors(arr)
	fmt.Println(arr)
}
