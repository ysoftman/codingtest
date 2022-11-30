/*
https://leetcode.com/problems/number-complement/
476. Number Complement
Easy

The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.

For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
Given an integer num, return its complement.

Example 1:
Input: num = 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.

Example 2:
Input: num = 1
Output: 0
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.

Constraints:
1 <= num < 231

Note: This question is the same as 1009: https://leetcode.com/problems/complement-of-base-10-integer/
*/
package main

import "fmt"

/*
num 보다 작은 수까지 2의 배수로 xor
101(5) xor 001(1) = 100
100 xor 010(2) = 110
110 xor 100(4) = 010(2)
*/
func findComplement(num int) int {
	for i := 1; i <= num; i *= 2 {
		num ^= i
	}
	return num
}

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}
