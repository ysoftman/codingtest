/*
https://leetcode.com/problems/ugly-number/
263. Ugly Number
Easy
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
Given an integer n, return true if n is an ugly number.

Example 1:
Input: n = 6
Output: true
Explanation: 6 = 2 × 3

Example 2:
Input: n = 1
Output: true
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.

Example 3:
Input: n = 14
Output: false
Explanation: 14 is not ugly since it includes the prime factor 7.

Constraints:
-231 <= n <= 231 - 1
*/

package main

import "fmt"

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	// n이 2,3,5 로 나누어 떨어지는지 본다.
	for _, v := range []int{2, 3, 5} {
		for n%v == 0 {
			n /= v
		}
	}
	// fmt.Println("n:", n)
	return n == 1
}

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(14))
	// 0과 음수 값은 모두 false
	fmt.Println(isUgly(0))
	fmt.Println(isUgly(-6))
	fmt.Println(isUgly(-1))
	fmt.Println(isUgly(-14))
}
