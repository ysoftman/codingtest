/*
https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
1502. Can Make Arithmetic Progression From Sequence
Easy
A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.
Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.

Example 1:
Input: arr = [3,5,1]
Output: true
Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.

Example 2:
Input: arr = [1,2,4]
Output: false
Explanation: There is no way to reorder the elements to obtain an arithmetic progression.

Constraints:
2 <= arr.length <= 1000
-106 <= arr[i] <= 106
*/
package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	sort.Ints(arr)
	diff := arr[0] - arr[1]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]-arr[i+1] != diff {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))
	fmt.Println(canMakeArithmeticProgression([]int{23, 4, 65, 51, 5, 5, 4, 15, 4, 7}))
}
