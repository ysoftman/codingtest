/*
https://leetcode.com/problems/gray-code/
89. Gray Code
Medium
An n-bit gray code sequence is a sequence of 2n integers where:
Every integer is in the inclusive range [0, 2n - 1],
The first integer is 0,
An integer appears no more than once in the sequence,
The binary representation of every pair of adjacent integers differs by exactly one bit, and
The binary representation of the first and last integers differs by exactly one bit.
Given an integer n, return any valid n-bit gray code sequence.

Example 1:
Input: n = 2
Output: [0,1,3,2]
Explanation:
The binary representation of [0,1,3,2] is [00,01,11,10].
- 00 and 01 differ by one bit
- 01 and 11 differ by one bit
- 11 and 10 differ by one bit
- 10 and 00 differ by one bit
[0,2,3,1] is also a valid gray code sequence, whose binary representation is [00,10,11,01].
- 00 and 10 differ by one bit
- 10 and 11 differ by one bit
- 11 and 01 differ by one bit
- 01 and 00 differ by one bit

Example 2:
Input: n = 1
Output: [0,1]

Constraints:
1 <= n <= 16
*/

package main

import "fmt"

/*
시작
000 0 arr[0]
001 1 arr[1]

011 3 arr[n-1] + 2^1
010 2 arr[n-2] + 2^1

110 6 arr[n-1] + 2^2
111 7 arr[n-2] + 2^2
101 5 arr[n-3] + 2^2
100 4 arr[n-4] + 2^2
*/

func pow(base, n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r *= base
	}
	return r
}

func grayCode(n int) []int {
	arr := []int{0, 1}
	if n == 1 {
		return arr
	}
	power := 1
	for len(arr) < pow(2, n) {
		for i := len(arr) - 1; i >= 0; i-- {
			arr = append(arr, arr[i]+pow(2, power))
		}
		power++
	}
	return arr
}

func main() {
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
	fmt.Println(grayCode(4))
	fmt.Println(grayCode(5))
	fmt.Println(grayCode(6))
	fmt.Println(grayCode(7))
}
