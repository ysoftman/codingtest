/*
https://leetcode.com/problems/duplicate-zeros/
1089. Duplicate Zeros
Easy
Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
Note that elements beyond the length of the original array are not written.
Do the above modifications to the input array in place and do not return anything.

Example 1:
Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Example 2:
Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]

Constraints:
1 <= arr.length <= 104
0 <= arr[i] <= 9
*/

package main

import "fmt"

func duplicateZeros(arr []int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	j := 0
	for i := 0; i < len(temp); i++ {
		if j == len(temp) {
			break
		}
		arr[j] = temp[i]
		j++
		if temp[i] == 0 {
			if j == len(temp) {
				break
			}
			arr[j] = 0
			j++
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
	arr = []int{1, 2, 3}
	duplicateZeros(arr)
	fmt.Println(arr)
	arr = []int{1, 5, 2, 0, 6, 8, 0, 6, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
	arr = []int{0, 0, 0, 0, 0, 0, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}
