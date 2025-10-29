/*
https://leetcode.com/problems/single-element-in-a-sorted-array/
540. Single Element in a Sorted Array
Medium
You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

Return the single element that appears only once.
Your solution must run in O(log n) time and O(1) space.

Example 1:
Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2

Example 2:
Input: nums = [3,3,7,7,10,11,11]
Output: 10

Constraints:
1 <= nums.length <= 105
0 <= nums[i] <= 105
*/
package main

import "fmt"

// binary search(logN)
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right + left) / 2
		// 2개씩 있다면
		// 0,1번째, 2,3 번째.. 등으로 같아야 함
		// mid+1 비교하기 위해 0,2,4,6 번째로 수정
		if mid%2 == 1 {
			mid--
		}
		// 중복 이면 왼쪽부분은 중복이었음, 그래서 오른쪽 부분에서 탐색해야 함
		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else { // 중복 아니면 왼쪽부분에 1개까지가 있는것이러, 왼쪽 부분에서 탐색해야 함
			right = mid
		}
	}
	return nums[left]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}
