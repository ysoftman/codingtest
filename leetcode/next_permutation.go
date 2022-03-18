/*
https://leetcode.com/problems/next-permutation/
31. Next Permutation
Medium
A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

For example, for arr = [1,2,3], the following are considered permutations of arr: [1,2,3], [1,3,2], [3,1,2], [2,3,1].
The next permutation of an array of integers is the next lexicographically greater permutation of its integer. More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

For example, the next permutation of arr = [1,2,3] is [1,3,2].
Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.

The replacement must be in place and use only constant extra memory.

Example 1:
Input: nums = [1,2,3]
Output: [1,3,2]

Example 2:
Input: nums = [3,2,1]
Output: [1,2,3]

Example 3:
Input: nums = [1,1,5]
Output: [1,5,1]

Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 100
*/
package main

import "fmt"

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	// 치환순서가 뒤쪽부터 시작이라 뒤쪽부터 크기순으로 되어 있는지 찾는다.
	i := len(nums) - 2
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	// case: 3,2,1 -> 1,2,3
	if i < 0 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	// case: 1 4 3 2
	// i = 1, j = 2, swap(i,j) => 2 4 3 1
	j := len(nums) - 1
	for j >= i {
		if nums[i] < nums[j] {
			break
		}
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	// reverse(i+1 ~ end of nums) => 2 1 3 4
	reverse(nums, i+1, len(nums)-1)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Print(nums, "=>", func([]int) []int { nextPermutation(nums); return nums }(nums), "\n")

	nums = []int{3, 2, 1}
	fmt.Print(nums, "=>", func([]int) []int { nextPermutation(nums); return nums }(nums), "\n")

	nums = []int{1, 1, 5}
	fmt.Print(nums, "=>", func([]int) []int { nextPermutation(nums); return nums }(nums), "\n")

	nums = []int{1, 4, 3, 2}
	fmt.Print(nums, "=>", func([]int) []int { nextPermutation(nums); return nums }(nums), "\n")

	nums = []int{1, 4, 5, 3, 2}
	fmt.Print(nums, "=>", func([]int) []int { nextPermutation(nums); return nums }(nums), "\n")
}
