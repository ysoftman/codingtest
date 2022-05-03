/*
https://leetcode.com/problems/peak-index-in-a-mountain-array/
852. Peak Index in a Mountain Array
Easy
Let's call an array arr a mountain if the following properties hold:
arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
Given an integer array arr that is guaranteed to be a mountain, return any i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

Example 1:
Input: arr = [0,1,0]
Output: 1

Example 2:
Input: arr = [0,2,1,0]
Output: 1

Example 3:
Input: arr = [0,10,5,2]
Output: 1

Constraints:
3 <= arr.length <= 104
0 <= arr[i] <= 106
arr is guaranteed to be a mountain array.

Follow up: Finding the O(n) is straightforward, could you find an O(log(n)) solution?
*/

package main

import "fmt"

func peakIndexInMountainArray2(arr []int) int {
	peak := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			peak = i + 1
		}
	}
	return peak
}

// O(logN)
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := (right-left)/2 + left
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
}
