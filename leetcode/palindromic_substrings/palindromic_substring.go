/*
https://leetcode.com/problems/palindromic-substrings
647. Palindromic Substrings
Medium
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.

Example 1:
Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Example 2:
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

Constraints:
1 <= s.length <= 1000
s consists of lowercase English letters.
*/
package main

import "fmt"

var palStrings = make(map[string]bool)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func subString(s string, cnt *int) {
	// fmt.Println(s) // 요거 출력하면 output limit exceeded 에러 발생
	if _, exist := palStrings[s]; exist {
		*cnt++
		return
	}
	if isPalindrome(s) {
		*cnt++
		palStrings[s] = true
		return
	}
}

func countSubstrings(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			subString(s[i:j], &cnt)
		}
	}
	return cnt
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc"))
}
