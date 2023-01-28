/*
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
28. Find the Index of the First Occurrence in a String
Medium
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

Constraints:
1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
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
