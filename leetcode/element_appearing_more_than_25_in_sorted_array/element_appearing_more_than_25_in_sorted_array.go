/*
https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
1287. Element Appearing More Than 25% In Sorted Array
Easy
Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time, return that integer.

Example 1:
Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6

Example 2:
Input: arr = [1,1]
Output: 1

Constraints:
1 <= arr.length <= 104
0 <= arr[i] <= 105
*/
package main

import "fmt"

func findSpecialInteger(arr []int) int {
	maxnum := arr[0]
	maxcnt := 0
	cnt := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if maxcnt < cnt {
				maxcnt = cnt
				maxnum = arr[i-1]
			}
			cnt = 0
		}
		cnt++
	}
	if maxcnt < cnt {
		maxcnt = cnt
		maxnum = arr[len(arr)-1]
	}
	return maxnum
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
	fmt.Println(findSpecialInteger([]int{1, 1}))
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 3, 4, 5, 5, 5, 5}))
	fmt.Println(findSpecialInteger([]int{1}))
	fmt.Println(findSpecialInteger([]int{1, 2, 3, 3}))
	fmt.Println(findSpecialInteger([]int{5668, 5668, 5668, 5668, 22011}))
}
