/*
https://leetcode.com/problems/hamming-distance/
461. Hamming Distance
Easy

The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, return the Hamming distance between them.

Example 1:
Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)

The above arrows point to positions where the corresponding bits are different.

Example 2:
Input: x = 3, y = 1
Output: 1

Constraints:
0 <= x, y <= 231 - 1
*/
package main

import "fmt"

func hammingDistance(x int, y int) int {
	// 비트 다르면 1인 xor 특성이용
	z := x ^ y
	cnt := 0
	// 2진수(bit)만들때 1카운트
	for z >= 2 {
		if z%2 == 1 {
			cnt++
		}
		z /= 2
	}
	if z == 1 {
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance(3, 1))
	fmt.Println(hammingDistance(1234541, 1251478))
}
