/*
https://leetcode.com/problems/find-k-closest-elements/
658. Find K Closest Elements
Medium

Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array. The result should also be sorted in ascending order.

An integer a is closer to x than an integer b if:

|a - x| < |b - x|, or
|a - x| == |b - x| and a < b

Example 1:
Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]

Example 2:
Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]

Constraints:
1 <= k <= arr.length
1 <= arr.length <= 104
arr is sorted in ascending order.
-104 <= arr[i], x <= 104
*/
package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// brute-force, time complexity: O(nk)
// 배열의 모든 원소-x 를 구해서 작은 순서대로 결과에 넣는다.
func findClosestElements2(arr []int, k int, x int) []int {
	r := []int{}
	for i := 0; i < k; i++ {
		r = append(r, 1<<31-1)
	}

	for i := 0; i < len(arr); i++ {
		diff := abs(arr[i] - x)
		for j := 0; j < len(r); j++ {
			if diff < abs(r[j]-x) || (diff == abs(r[j]-x) && arr[i] < r[j]) {
				// r = append(r, 0)
				copy(r[j+1:], r[j:])
				r[j] = arr[i]
				// fmt.Println(r)
				break
			}
		}
	}
	sort.Ints(r)
	return r
}

// binary search, time complexity: O(logN)
func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	// k 길이 만큼을 하나의 묶음으로 계산하기 위해
	right := len(arr) - k
	for left < right {
		// mid := (left + right) / 2
		mid := (right-left)/2 + left
		// mid 위치와 mid+k 위치에서의 x와의 차이를 중 작은쪽으로 left, right 를 조정
		// 2,2,2,2 으로 연속으로 같은 원소가 있는 경우가 있어 abs 로 차이를 계산하면 안된다.
		// 같은 원소의 경우 x 에 따라서 왼쪽,오른쪽으로 움직이도록 해야 한다.
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// left ~ left+k 중에 가장 치이가 작은 원소가 있다.
	return arr[left : left+k]
}

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 1))
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
	fmt.Println(findClosestElements([]int{1, 1, 2, 2, 2, 2, 2, 3, 3}, 3, 3))
	fmt.Println(findClosestElements([]int{1, 1, 2, 2, 2, 2, 2, 3, 3}, 3, 1))
}
