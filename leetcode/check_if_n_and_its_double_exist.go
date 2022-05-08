/*
https://leetcode.com/problems/check-if-n-and-its-double-exist/
1346. Check If N and Its Double Exist
Easy
Given an array arr of integers, check if there exists two integers N and M such that N is the double of M ( i.e. N = 2 * M).
More formally check if there exists two indices i and j such that :
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

Example 1:
Input: arr = [10,2,5,3]
Output: true
Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.

Example 2:
Input: arr = [7,1,14,11]
Output: true
Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.

Example 3:
Input: arr = [3,1,7,11]
Output: false
Explanation: In this case does not exist N and M, such that N = 2 * M.

Constraints:
2 <= arr.length <= 500
-10^3 <= arr[i] <= 10^3
*/
package main

import (
	"fmt"
	"sort"
)

// using hash table
func checkIfExist2(arr []int) bool {
	m := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = i
	}
	for i := 0; i < len(arr); i++ {
		if v, ok := m[arr[i]*2]; ok {
			// [1,0,2] 0이 한개 있는 경우 자기 자신 0은 안됨
			// [1,0,0] 2개 있는 경우, index 가 다른 0은 ok
			if i != v {
				return true
			}
		}
	}
	return false
}

// using binary search
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		left, right := 0, len(arr)-1
		for left <= right {
			mid := (right-left)/2 + left
			if arr[i]*2 == arr[mid] {
				// [1,0,2] 0이 한개 있는 경우 자기 자신 0은 안됨
				// [1,0,0] 2개 있는 경우, index 가 다른 0은 ok
				if arr[i] == 0 && i == mid {
					break
				}
				return true
			}
			if arr[i]*2 > arr[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{7, 1, 14, 11}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
	fmt.Println(checkIfExist([]int{0, 0}))
	fmt.Println(checkIfExist([]int{-10, 12, -20, -8, 15}))
}
