/*
https://leetcode.com/problems/convert-a-number-to-hexadecimal/
405. Convert a Number to Hexadecimal
Easy

Given an integer num, return a string representing its hexadecimal representation. For negative integers, two’s complement method is used.

All the letters in the answer string should be lowercase characters, and there should not be any leading zeros in the answer except for the zero itself.

Note: You are not allowed to use any built-in library method to directly solve this problem.

Example 1:
Input: num = 26
Output: "1a"

Example 2:
Input: num = -1
Output: "ffffffff"
Constraints:
-231 <= num <= 231 - 1
*/
package main

import "fmt"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	cnt := 0
	r := ""
	// 2^31 > 32(4*8) 비트만큼만 처리
	for num != 0 && cnt < 8 {
		// 0xf -> 0b1111(15) 로 4비트씩 and 하면 15로 나눈 나머지효과
		r = string(hex[num&0xf]) + r
		// 4비트(16진수)단위로 처리한 부분은 >> 4로 밀어낸다.(삭제)
		num >>= 4
		cnt++
	}
	return r
}

func main() {
	fmt.Println(toHex(26))
	fmt.Println(toHex(123))
	fmt.Println(toHex(10))
	fmt.Println(toHex(0))
	fmt.Println(toHex(3))
	fmt.Println(toHex(12381293))
	fmt.Println(toHex(-1))
	fmt.Println(toHex(-5))
}
