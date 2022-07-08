/*
https://leetcode.com/problems/intersection-of-two-arrays/
349. Intersection of Two Arrays
Easy
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.

Constraints:
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/
package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		if _, exist := m[nums2[i]]; exist {
			m[nums2[i]]++
		}
	}

	r := []int{}
	for k, v := range m {
		if v >= 2 {
			r = append(r, k)
		}
	}
	return r
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4, 1, 1, 1}))
}
