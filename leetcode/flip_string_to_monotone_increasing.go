/*
https://leetcode.com/problems/flip-string-to-monotone-increasing/
926. Flip String to Monotone Increasing
Medium

A binary string is monotone increasing if it consists of some number of 0's (possibly none), followed by some number of 1's (also possibly none).

You are given a binary string s. You can flip s[i] changing it from 0 to 1 or from 1 to 0.

Return the minimum number of flips to make s monotone increasing.

Example 1:
Input: s = "00110"
Output: 1
Explanation: We flip the last digit to get 00111.

Example 2:
Input: s = "010110"
Output: 2
Explanation: We flip to get 011111, or alternatively 000111.

Example 3:
Input: s = "00011000"
Output: 2
Explanation: We flip to get 00000000.

Constraints:
1 <= s.length <= 105
s[i] is either '0' or '1'.
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFlipsMonoIncr(s string) int {
	cntFlip := 0
	cntOne := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cntOne++
		} else {
			cntFlip++
			// 0 일때, 0->1로 뒤집는 횟수와 1->0으로 뒤집는 횟수 중 작은쪽을 선택
			cntFlip = min(cntFlip, cntOne)
		}
	}
	return cntFlip
}

func main() {
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("010110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
}
