/*
https://leetcode.com/problems/intersection-of-two-arrays-ii/

350. Intersection of Two Arrays II
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.
*/

package main

import (
	"fmt"
	"sort"
)

// brute-force
func intersect2(nums1 []int, nums2 []int) []int {
	result := []int{}
	// O(mn)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				nums2 = append(nums2[:j], nums2[j+1:]...)
				result = append(result, nums1[i])
				break
			}

		}
	}
	return result
}

// using sort
func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	sort.Ints(nums1) // O(nLogn)
	sort.Ints(nums2) // O(mLogm)
	j := 0
	// O(n)
	for i := 0; i < len(nums1); {
		if j == len(nums2) {
			break
		}
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
