/*
https://leetcode.com/problems/word-pattern/
290. Word Pattern
Easy

Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

Example 1:
Input: pattern = "abba", s = "dog cat cat dog"
Output: true

Example 2:
Input: pattern = "abba", s = "dog cat cat fish"
Output: false

Example 3:
Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false

Constraints:
1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.
*/
package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		val, exist := m1[pattern[i]]
		if !exist {
			if val, exist := m2[strs[i]]; exist {
				if val != pattern[i] {
					return false
				}
			}

			// pattern , s 를 키로 조회 가능하도록 저장해놓는다.
			m1[pattern[i]] = strs[i]
			m2[strs[i]] = pattern[i]
			continue
		}
		if val != strs[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("aaa", "aa aa aa aa"))
}
