/*
https://leetcode.com/problems/4sum-ii/
454. 4Sum II
Medium
Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

Example 1:
Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
Output: 2
Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

Example 2:
Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
Output: 1

Constraints:
n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28
*/
package main

import "fmt"

// time complexity: O(n*n)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// nums1*nums2*nums3*nums4 로 하면 O(n*n*n*n) 로 시간 복잡도가 너무 올라간다.

	// 우선 nums1, nums2 의 합의 조합을 맵으로 저장, O(n*n)
	m12 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			// nums1[i]+nums2[j] 가 이미 있을 수도 있기 때문에 중복 개수를 카운트해둔다.
			m12[nums1[i]+nums2[j]]++
		}
	}
	// nums3, nums4 의 합의 조합을 맵으로 저장, O(n*n)
	m34 := make(map[int]int)
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			m34[nums3[i]+nums4[j]]++
		}
	}

	// fmt.Println("m12:", m12)
	// fmt.Println("m34:", m34)

	r := 0
	// m12 + m34 = 0 이 되는 경우
	for k, v := range m12 {
		if cnt, exist := m34[k*-1]; exist {
			// fmt.Println("k:", k, "k*-1:", k*-1, "cnt:", cnt, "v:", v)
			r += cnt * v
		}
	}
	return r
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))

	nums1 = []int{0}
	nums2 = []int{0}
	nums3 = []int{0}
	nums4 = []int{0}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))

	nums1 = []int{1, 2, 3}
	nums2 = []int{-1, -2, -3}
	nums3 = []int{4, 5, 6}
	nums4 = []int{-4, -5, -6}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
