/*
https://leetcode.com/problems/valid-perfect-square/
367. Valid Perfect Square
Easy
Given a positive integer num, write a function which returns True if num is a perfect square else False.
Follow up: Do not use any built-in library function such as sqrt.

Example 1:
Input: num = 16
Output: true

Example 2:
Input: num = 14
Output: false

Constraints:
1 <= num <= 2^31 - 1
*/
package main

import "fmt"

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (right-left)/2 + left
		if (mid)*(mid) == num {
			return true
		}
		if (mid)*(mid) > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(36))
	fmt.Println(isPerfectSquare(121))
}
