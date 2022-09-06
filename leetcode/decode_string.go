/*
https://leetcode.com/problems/decode-string/
394. Decode String
Medium
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.

Example 1:
Input: s = "3[a]2[bc]"
Output: "aaabcbc"

Example 2:
Input: s = "3[a2[c]]"
Output: "accaccacc"

Example 3:
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"

Constraints:
1 <= s.length <= 30
s consists of lowercase English letters, digits, and square brackets '[]'.
s is guaranteed to be a valid input.
All the integers in s are in the range [1, 300].
*/
package main

import (
	"fmt"
	"strconv"
)

// using stack
func decodeString(s string) string {
	r := ""
	cnt := 0
	str := ""
	for i := 0; i < len(s); i++ {
		cntstr := ""
		for s[i] >= '0' && s[i] <= '9' {
			cntstr += string(s[i])

			i++
		}
		if len(cntstr) > 0 {
			cnt, _ = strconv.Atoi(cntstr)
			// fmt.Println(cnt)
		}
		if s[i] == '[' {
			start := i
			end := 0
			a := 0
			for i < len(s) {
				if s[i] == '[' {
					a++
				}
				if s[i] == ']' {
					a--
				}
				if a == 0 {
					end = i
					break
				}
				i++
			}
			if start < end {
				// fmt.Println(s[start+1 : end])
				str = decodeString(s[start+1 : end])
			}
			for j := 0; j < cnt; j++ {
				r += str
			}
		} else {
			r += string(s[i])
		}
	}
	return r
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("100[leetcode]"))
}
