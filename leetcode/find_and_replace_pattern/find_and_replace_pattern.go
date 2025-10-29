/*
https://leetcode.com/problems/find-and-replace-pattern/
890. Find and Replace Pattern
Medium
Given a list of strings words and a string pattern, return a list of words[i] that match pattern. You may return the answer in any order.

A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.

Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.

Example 1:
Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]
Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation, since a and b map to the same letter.

Example 2:
Input: words = ["a","b","c"], pattern = "a"
Output: ["a","b","c"]


Constraints:
1 <= pattern.length <= 20
1 <= words.length <= 50
words[i].length == pattern.length
pattern and words[i] are lowercase English letters.
*/
package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	r := []string{}
	for _, word := range words {
		m := make(map[byte]byte, 0)
		m2 := make(map[byte]byte, 0)
		idx := 0
		for i := 0; i < len(word); i++ {
			if _, exist := m[pattern[idx]]; !exist {
				// 중복 매핑되지 않도록 _ 문자를 매핑한다.
				if _, ok := m2[word[i]]; !ok {
					m[pattern[idx]] = word[i]
				} else {
					m[pattern[idx]] = '_'
				}
				m2[word[i]] = '_'
			}
			if m[pattern[idx]] != word[i] {
				break
			}
			idx++
		}
		if idx >= len(word) {
			r = append(r, word)
		}
	}
	return r
}

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
	fmt.Println(findAndReplacePattern([]string{"a", "b", "c"}, "a"))
}
