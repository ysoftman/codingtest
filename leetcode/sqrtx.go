/*
https://leetcode.com/problems/sqrtx/
69. Sqrt(x)
Easy
Given a non-negative integer x, compute and return the square root of x.
Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.
Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.

Example 1:
Input: x = 4
Output: 2

Example 2:
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.
*/
package main

import "fmt"

func mySqrt(x int) int {
	for i := 0; i < (1<<31 - 1); i++ {
		if i*i > x {
			return i - 1
		}
	}
	return 0
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
