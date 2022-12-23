/*
https://leetcode.com/problems/sort-an-array/
912. Sort an Array
Medium

Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

Example 1:
Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).

Example 2:
Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
Explanation: Note that the values of nums are not necessairly unique.

Constraints:
1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104
*/
package main

import "fmt"

// bubbleSort, time complexity: O(n*n) --> time limit exceeded
func sortArray2(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				// swap
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// quickSort, time complexity: O(nlogn)
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	start := left
	end := right
	pivot := nums[end]
	for left < right {
		// 바꿔야할 왼쪽 지점 찾기
		if nums[left] <= pivot {
			left++
			continue
		}
		// 바꿔야할 오른쪽 지점 찾기
		if nums[right] >= pivot {
			right--
			continue
		}
		// 2 지점 스왑
		nums[left], nums[right] = nums[right], nums[left]
	}
	// pivot 위치의 값도 left 위치와 스왑
	// nums[left] 값1개는 정렬이 끝남
	nums[left], nums[end] = nums[end], nums[left]

	// 정령된 값 기준으로 왼쪽 부분와 오른쪽 부분에 대해서 반복
	quickSort(nums, start, left-1)
	quickSort(nums, left+1, end)
}
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
