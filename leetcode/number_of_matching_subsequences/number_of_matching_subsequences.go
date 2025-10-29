/*
https://leetcode.com/problems/number-of-matching-subsequences/
792. Number of Matching Subsequences
Medium
Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".

Example 1:
Input: s = "abcde", words = ["a","bb","acd","ace"]
Output: 3
Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".

Example 2:
Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
Output: 2

Constraints:
1 <= s.length <= 5 * 104
1 <= words.length <= 5000
1 <= words[i].length <= 50
s and words[i] consist of only lowercase English letters.
*/
package main

import (
	"fmt"
)

// time limit exceeded
func numMatchingSubseq2(s string, words []string) int {
	num := 0
	mword := make(map[int]string, len(words))
	for i := 0; i < len(words); i++ {
		mword[i] = words[i]
	}
	// s 의 문자 에 대해 n 개의 word 내의 문자와 매치되는지 확인
	for i := 0; i < len(s); i++ {
		for k, v := range mword {
			if len(v) == 0 {
				continue
			}
			// 매치되면 매치 문자를 삭제한다.
			if s[i] == v[0] {
				mword[k] = v[1:]
			}
			// 매치로 삭제된 word만 카운트한다.
			if len(mword[k]) == 0 {
				num++
				delete(mword, k)
			}
		}
	}
	// fmt.Println("mword:", mword)
	return num
}

// brute force, time complexity: O(m*n)
// 통과는 됐는데 느림(2800ms)
func numMatchingSubseq(s string, words []string) int {
	num := 0
	for _, word := range words {
		j := 0
		for i := 0; i < len(s); i++ {
			if s[i] == word[j] {
				j++
			}
			if j == len(word) {
				num++
				break
			}
		}
	}
	return num
}

func main() {
	s := "abcde"
	words := []string{"a", "a", "ab", "ab"}
	fmt.Println(numMatchingSubseq(s, words))
	s = "abcde"
	words = []string{"a", "bb", "acd", "ace"}
	fmt.Println(numMatchingSubseq(s, words))
	s = "dsahjpjauf"
	words = []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}
	fmt.Println(numMatchingSubseq(s, words))
}
