/*
https://leetcode.com/problems/unique-number-of-occurrences/
1207. Unique Number of Occurrences
Easy

Given an array of integers arr, return true if the number of occurrences of each value in the array is unique, or false otherwise.

Example 1:
Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

Example 2:
Input: arr = [1,2]
Output: false

Example 3:
Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true

Constraints:
1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
*/
package main

import (
	"fmt"
	"sort"
)

// 정렬해서 같값이 연속으로 나오는지 체크
func uniqueOccurrences2(arr []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	slice := make([]int, len(m))
	i := 0
	for _, v := range m {
		slice[i] = v
		i++
	}
	sort.Ints(slice)
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			return false
		}
	}
	return true
}

// 찾은 개수로 맵을 구성하여, 중복되는지 체크
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	m2 := make(map[int]int)
	for _, v := range m {
		if _, exist := m2[v]; exist {
			return false
		}
		m2[v]++
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
