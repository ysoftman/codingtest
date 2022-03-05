/*
https://leetcode.com/problems/implement-strstr/
28. Implement strStr()
Easy
Implement strStr().
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
Clarification:
What should we return when needle is an empty string? This is a great question to ask during an interview.
For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1

Example 3:
Input: haystack = "", needle = ""
Output: 0
*/

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		j := 0
		for j < len(needle) {
			if i+j >= len(haystack) || haystack[i+j] != needle[j] {
				break
			}
			j++
		}
		if j == len(needle) && j <= len(haystack)-i {
			return i
		}
	}

	return -1

}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
	haystack = "aaaaa"
	needle = "bba"
	fmt.Println(strStr(haystack, needle))
	haystack = ""
	needle = ""
	fmt.Println(strStr(haystack, needle))
}
