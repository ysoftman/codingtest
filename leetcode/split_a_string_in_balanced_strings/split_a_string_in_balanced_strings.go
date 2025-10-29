/*
https://leetcode.com/problems/split-a-string-in-balanced-strings/description/
1221. Split a String in Balanced Strings
Easy

Balanced strings are those that have an equal quantity of 'L' and 'R' characters.

Given a balanced string s, split it into some number of substrings such that:

Each substring is balanced.
Return the maximum number of balanced strings you can obtain.

Example 1:
Input: s = "RLRRLLRLRL"
Output: 4
Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

Example 2:
Input: s = "RLRRRLLRLL"
Output: 2
Explanation: s can be split into "RL", "RRRLLRLL", each substring contains same number of 'L' and 'R'.
Note that s cannot be split into "RL", "RR", "RL", "LR", "LL", because the 2nd and 5th substrings are not balanced.

Example 3:
Input: s = "LLLLRRRR"
Output: 1
Explanation: s can be split into "LLLLRRRR".

Constraints:
2 <= s.length <= 1000
s[i] is either 'L' or 'R'.
s is a balanced string.
*/

package main

import "fmt"

func balancedStringSplit(s string) int {
	result := 0
	lrCnt := 0
	for i := range s {
		if s[i] == 'L' {
			lrCnt++
		} else if s[i] == 'R' {
			lrCnt--
		}
		if lrCnt == 0 {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
	fmt.Println(balancedStringSplit("LLLLRRRR"))
}
