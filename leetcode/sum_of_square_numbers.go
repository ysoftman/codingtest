/*
https://leetcode.com/problems/sum-of-square-numbers/
633. Sum of Square Numbers
Medium
Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:
Input: c = 5
Output: true
Explanation: 1 * 1 + 2 * 2 = 5

Example 2:
Input: c = 3
Output: false

Constraints:
0 <= c <= 231 - 1
*/
package main

import "fmt"

func binarySearch(left, right, target int) int {
	for left <= right {
		mid := (right-left)/2 + left
		// fmt.Println("mid:", mid)
		if mid*mid == target {
			return mid
		}
		if mid*mid < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		// fmt.Println(a, c-(a*a))
		if binarySearch(a, c, c-(a*a)) >= 0 {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(judgeSquareSum(0))
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(310))
	fmt.Println(judgeSquareSum(20000))
	fmt.Println(judgeSquareSum(2147482647))
}
