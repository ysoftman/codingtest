/*
https://leetcode.com/problems/power-of-three/
326. Power of Three
Easy
Given an integer n, return true if it is a power of three. Otherwise, return false.

An integer n is a power of three, if there exists an integer x such that n == 3x.

Example 1:
Input: n = 27
Output: true

Example 2:
Input: n = 0
Output: false
Example 3:

Input: n = 9
Output: true


Constraints:

-231 <= n <= 231 - 1

Follow up: Could you solve it without loops/recursion?
*/
package main

import "fmt"

func isPowerOfThree2(n int) bool {
	// 3^0 = 1
	if n == 1 {
		return true
	}
	if n <= 0 || n == 2 {
		return false
	}

	temp := 1
	for temp < n {
		temp *= 3
		if temp == n {
			return true
		}
	}
	return false
}

func isPowerOfThree3(n int) bool {
	// 3^0 = 1
	if n == 1 {
		return true
	}
	if n <= 0 || n == 2 {
		return false
	}

	for n >= 3 {
		if n%3 != 0 {
			break
		}
		n /= 3
	}
	if n == 1 {
		return true
	}
	return false
}

// without loop, time complexity: O(1)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	// 2^31 2147483648 범위내에서 가장 큰 3의 거듭제곱의 최대는
	// 3^19 1162261467 이다.
	// n 이 3의 제곱이라면 n 으로 나누어 떨어져야 한다.(n 을 계속 곱하면 1162261467 가 되어야 한다.)
	if 1162261467%n == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(2))
	fmt.Println(isPowerOfThree(99))
	fmt.Println(isPowerOfThree(81))
	fmt.Println(isPowerOfThree(243))
	fmt.Println(isPowerOfThree(729))
	fmt.Println(isPowerOfThree(-1))
	fmt.Println(isPowerOfThree(-3))
	fmt.Println(isPowerOfThree(6))
}
