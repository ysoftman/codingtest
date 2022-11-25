/*
https://leetcode.com/problems/repeated-substring-pattern/
459. Repeated Substring Pattern
Easy
Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

Example 1:
Input: s = "abab"
Output: true
Explanation: It is the substring "ab" twice.

Example 2:
Input: s = "aba"
Output: false

Example 3:
Input: s = "abcabcabcabc"
Output: true
Explanation: It is the substring "abc" four times or the substring "abcabc" twice.

Constraints:
1 <= s.length <= 104
s consists of lowercase English letters.
*/
package main

import (
	"fmt"
	"strings"
)

// brute forace --> passed
func repeatedSubstringPattern2(s string) bool {
	for i := len(s) / 2; i >= 1; i-- {
		pattern := s[:i]
		temp := s[i:]
		if len(temp)%len(pattern) != 0 {
			continue
		}
		matched := true
		for j := i; j < len(s); j += i {
			if pattern != s[j:j+len(pattern)] {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}
	return false
}

/*
simple solution
ex) abab
abab+abab=abababab 에서, 앞1글자, 끝1글제외, bababa
bababa 에서 abab 를 찾으면 true 아니면 false
*/
func repeatedSubstringPattern(s string) bool {
	ss := (s + s)[1 : len(s+s)-1]
	// fmt.Println(ss)
	return strings.Contains(ss, s)
}

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("ababab"))
}
