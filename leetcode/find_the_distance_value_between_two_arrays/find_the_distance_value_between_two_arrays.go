/*
https://leetcode.com/problems/find-the-distance-value-between-two-arrays/
1385. Find the Distance Value Between Two Arrays
Easy
Given two integer arrays arr1 and arr2, and the integer d, return the distance value between the two arrays.
The distance value is defined as the number of elements arr1[i] such that there is not any element arr2[j] where |arr1[i]-arr2[j]| <= d.

Example 1:
Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
Output: 2
Explanation:
For arr1[0]=4 we have:
|4-10|=6 > d=2
|4-9|=5 > d=2
|4-1|=3 > d=2
|4-8|=4 > d=2
For arr1[1]=5 we have:
|5-10|=5 > d=2
|5-9|=4 > d=2
|5-1|=4 > d=2
|5-8|=3 > d=2
For arr1[2]=8 we have:
|8-10|=2 <= d=2
|8-9|=1 <= d=2
|8-1|=7 > d=2
|8-8|=0 <= d=2

Example 2:
Input: arr1 = [1,4,2,3], arr2 = [-4,-3,6,10,20,30], d = 3
Output: 2

Example 3:
Input: arr1 = [2,1,100,3], arr2 = [-5,-2,10,-3,7], d = 6
Output: 1

Constraints:
1 <= arr1.length, arr2.length <= 500
-1000 <= arr1[i], arr2[j] <= 1000
0 <= d <= 100
*/
package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func findTheDistanceValue2(arr1 []int, arr2 []int, d int) int {
	cnt := 0
	for i := 0; i < len(arr1); i++ {
		skip := false
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				skip = true
				break
			}
		}
		if skip == false {
			cnt++
		}
	}
	return cnt
}

// time complexity O(nLogn)+O(LogN) = O(nLogn)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	cnt := 0
	// sort.Ints(arr1)  // O(nLogn)
	sort.Ints(arr2) // O(nLogn)
	// fmt.Println("arr1:",arr1)
	// fmt.Println("arr2:",arr2)

	for i := 0; i < len(arr1); i++ {
		left, right := 0, len(arr2)-1
		skip := false
		// binary search O(LogN)
		for left <= right {
			mid := (right-left)/2 + left
			// fmt.Println("arr1[i]:", arr1[i], "arr2[mid]", arr2[mid])
			if abs(arr1[i]-arr2[mid]) <= d {
				// fmt.Println("--", arr1[i])
				skip = true
				break
			}
			// 값이 음수가 있어 abs 결과대시 비교대상 값들 크기를 비교한다.
			if arr1[i] < arr2[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if skip == false {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
	fmt.Println(findTheDistanceValue([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3))
	fmt.Println(findTheDistanceValue([]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6))
	fmt.Println(findTheDistanceValue([]int{-3, 10, 2, 8, 0, 10}, []int{-9, -1, -4, -9, -8}, 9))
}
