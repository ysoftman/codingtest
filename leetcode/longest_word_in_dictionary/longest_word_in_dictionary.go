/*
https://leetcode.com/problems/longest-word-in-dictionary
720. Longest Word in Dictionary
Solved
Medium
Topics
premium lock icon
Companies
Hint
Given an array of strings words representing an English Dictionary, return the longest word in words that can be built one character at a time by other words in words.

If there is more than one possible answer, return the longest word with the smallest lexicographical order. If there is no answer, return the empty string.

Note that the word should be built from left to right with each additional character being added to the end of a previous word.

Example 1:

Input: words = ["w","wo","wor","worl","world"]
Output: "world"
Explanation: The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
Example 2:

Input: words = ["a","banana","app","appl","ap","apply","apple"]
Output: "apple"
Explanation: Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".

Constraints:

1 <= words.length <= 1000
1 <= words[i].length <= 30
words[i] consists of lowercase English letters.
*/
package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	sort.Slice(words, func(a, b int) bool {
		return words[a] < words[b]
	})
	m := make(map[string]bool)
	for _, w := range words {
		if _, exist := m[w[:len(w)-1]]; exist {
			m[w] = true
		} else if len(w) == 1 {
			m[w] = true
		}
	}
	r := ""
	for k := range m {
		if len(r) < len(k) {
			r = k
		} else if len(r) == len(k) {
			// lexical order
			if r > k {
				r = k
			}
		}
	}
	return r
}

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
