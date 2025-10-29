/*
https://leetcode.com/problems/word-break-ii/
140. Word Break II
Hard

Given a string s and a dictionary of strings wordDict, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences in any order.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
Output: ["cats and dog","cat sand dog"]

Example 2:
Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: []

Constraints:
1 <= s.length <= 20
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 10
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.
*/
package main

import (
	"fmt"
	"strings"
)

func dfs(s string, wordDict []string, m map[string][]string) []string {
	if val, ok := m[s]; ok {
		return val
	}
	r := []string{}
	if len(s) == 0 {
		// r 원소가 0이 아니게 만들어야 caller 에서 word 를 추가 할 수 있다.
		r = append(r, "")
		return r
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			// fmt.Println("----", s[:len(word)], wordDict, m)
			sublist := dfs(s[len(word):], wordDict, m)
			for _, sub := range sublist {
				temp := word
				if len(sub) == 0 {
					temp += sub
				} else {
					temp += " " + sub
				}
				r = append(r, temp)
			}
		}
	}
	m[s] = r
	return r
}

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string][]string)
	return dfs(s, wordDict, m)
}

func main() {
	fmt.Printf("%q\n", wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Printf("%q\n", wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Printf("%q\n", wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
