/*
https://leetcode.com/problems/find-all-anagrams-in-a-string/
438. Find All Anagrams in a String
Medium
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".

Constraints:
1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.
*/
package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(a, b int) bool {
		return bytes[a] < bytes[b]
	})
	return string(bytes)
}
func findAnagrams(s string, p string) []int {
	r := []int{}
	p = sortString(p)
	l := len(p)
	for i := 0; i <= len(s)-l; i++ {
		// fmt.Println(s[i : i+l])
		temp := sortString(s[i : i+l])
		if temp == p {
			r = append(r, i)
		}
	}
	return r
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
