/*
https://leetcode.com/problems/single-number/
136. Single Number
Easy
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:
Input: nums = [2,2,1]
Output: 1

Example 2:
Input: nums = [4,1,2,1,2]
Output: 4

Example 3:
Input: nums = [1]
Output: 1
*/

package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			i++
			continue
		}
		return nums[i]
	}
	return nums[len(nums)-1]
}

// using temporary space
// func singleNumber(nums []int) int {
//     temps := make(map[int]string, 0)
//     for i := 0; i < len(nums); i++ {
//         if temps[nums[i]] == "" {
//             temps[nums[i]] = "a"
//         } else {
//             temps[nums[i]] = "b"
//         }
//     }
//     for k, v := range temps {
//         if v == "a" {
//             return k
//         }
//     }
//     return 0
// }

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}
