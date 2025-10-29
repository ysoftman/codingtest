/*
https://leetcode.com/problems/power-of-four/
342. Power of Four
Easy
Given an integer n, return true if it is a power of four. Otherwise, return false.

An integer n is a power of four, if there exists an integer x such that n == 4x.

Example 1:
Input: n = 16
Output: true

Example 2:
Input: n = 5
Output: false
Example 3:

Input: n = 1
Output: true

Constraints:
-231 <= n <= 231 - 1


Follow up: Could you solve it without loops/recursion?
*/
package main

import "fmt"

func isPowerOfFour2(n int) bool {
	if n == 1 {
		return true
	}
	temp := 1
	for temp < n {
		temp *= 4
	}
	if temp > n {
		return false
	}
	return true
}

// shift bit
// 100 (4)
// 10000 (16)
// 100000000 (256)
func isPowerOfFour3(n int) bool {
	if n == 1 {
		return true
	}
	// 비트 2개씩 왼쪽 쉬프트
	for i := 0; i < 32; i += 2 {
		if 1<<i == n {
			return true
		}
	}
	return false
}

// O(1), without loops/recursion
/*
n 과 (n-1) 를 and 연산해서 0이면 2의 제곱인경우
    1(1) 10(2) 100(4) 1000(8)
  & 0(0) 01(1) 011(3) 0111(7)
   --------------------------
    0    00    000    0000

2의 제곱중에서 홀수 자리에 1인 값
1010101010101010101010101010101(0x55555555)을
and 해서 0이 아니면 4의 제곱인 수다.

  1(1)
& 1
----
  1

  10(2)
& 01
----
  11

  100(4)
& 101
-----
  101

  1000(8)
& 0101
------
  0000  --> 0은 4제곱수가 아니다.


  10000(16)
& 10101
-------
  10000
*/
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	// 2의 제곱 수 여부 파악
	if n&(n-1) != 0 {
		return false
	}
	// 4의 제곱 수 여부 파악
	if n&0x55555555 == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(256))
	fmt.Println(isPowerOfFour(12345))
	fmt.Println(isPowerOfFour(4096))
}
