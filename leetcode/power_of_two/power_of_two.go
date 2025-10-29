/*
https://leetcode.com/problems/power-of-two/
231. Power of Two
Easy
Given an integer n, return true if it is a power of two. Otherwise, return false.
An integer n is a power of two, if there exists an integer x such that n == 2x.

Example 1:
Input: n = 1
Output: true
Explanation: 20 = 1

Example 2:
Input: n = 16
Output: true
Explanation: 24 = 16

Example 3:
Input: n = 3
Output: false
*/
package main

import "fmt"

// O(N)
func isPowerOfTwo2(n int) bool {
	for i := 0; i < 31; i++ {
		if 1<<i == n {
			return true
		}
		if 1<<i > n {
			break
		}
	}
	return false
}

// O(1)
/*
    1(1) 10(2) 100(4) 1000(8)
   &0(0) 01(1) 011(3) 0111(7)
   --------------------------
    0    00    000    0000
*/
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPowerOfTwo(-2))
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}
