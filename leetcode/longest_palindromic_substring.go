/*
https://leetcode.com/problems/longest-palindromic-substring/
5. Longest Palindromic Substring

Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"
*/

package main

import "fmt"

func longestPalindrome(s string) string {
	longestPal := ""
	for i := range s {
		minlen := len(s[:i])
		if minlen > len(s[i:]) {
			minlen = len(s[i:])
		}
		// fmt.Println("minlen:", minlen)
		temp := string(s[i])
		// bb 인 경우
		for j := 0; j <= minlen; j++ {
			if i-j-1 < 0 || i+j >= len(s) {
				break
			}
			// fmt.Printf("i: %v j: %v  s[i-j-1]:%v  s[i+j]:%v\n", i, j, string(s[i-j-1]), string(s[i+j]))
			if s[i-j-1] == s[i+j] {
				temp = s[i-j-1 : i+j+1]
			} else {
				break
			}
		}
		if len(longestPal) < len(temp) {
			longestPal = temp
		}
		temp = ""
		// bab 인 경우
		for j := 1; j <= minlen; j++ {
			if i+1+j > len(s) {
				break
			}
			// fmt.Printf("i: %v j: %v  s[i-j]:%v  s[i+j]:%v\n", i, j, string(s[i-j]), string(s[i+j]))
			if s[i-j] == s[i+j] {
				temp = s[i-j : i+j+1]
			} else {
				break
			}
		}
		if len(longestPal) < len(temp) {
			longestPal = temp
		}
	}

	return longestPal
}

func main() {
	fmt.Println("longestPalindrome")
	fmt.Println("babad => ", longestPalindrome("babad"))
	fmt.Println("abba => ", longestPalindrome("abba"))
	fmt.Println("cbbc => ", longestPalindrome("cbbc"))
	fmt.Println("a => ", longestPalindrome("a"))
	fmt.Println("bb => ", longestPalindrome("bb"))
	fmt.Println("ccc => ", longestPalindrome("ccc"))
	fmt.Println("abcba => ", longestPalindrome("abcba"))
}
