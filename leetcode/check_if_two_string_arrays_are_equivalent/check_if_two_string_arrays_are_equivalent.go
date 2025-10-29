/*
https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
1662. Check If Two String Arrays are Equivalent
Easy

Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.

A string is represented by an array if the array elements concatenated in order forms the string.

Example 1:
Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Explanation:
word1 represents string "ab" + "c" -> "abc"
word2 represents string "a" + "bc" -> "abc"
The strings are the same, so return true.

Example 2:
Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
Output: false

Example 3:
Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
Output: true

Constraints:
1 <= word1.length, word2.length <= 103
1 <= word1[i].length, word2[i].length <= 103
1 <= sum(word1[i].length), sum(word2[i].length) <= 103
word1[i] and word2[i] consist of lowercase letters.
*/
package main

import "fmt"

// simple
// time complexity: O(n)
// space complexity: O(m+n)
func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	a := ""
	for i := 0; i < len(word1); i++ {
		a += word1[i]
	}
	b := ""
	for i := 0; i < len(word2); i++ {
		b += word2[i]
	}
	if a == b {
		return true
	}
	return false
}

// time complexity: O(n)
// space complexity: O(1)
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i := 0
	j := 0
	a := 0
	b := 0
	for i < len(word1) && j < len(word2) {
		if word1[i][a] != word2[j][b] {
			return false
		}
		if a == len(word1[i])-1 {
			a = 0
			i++
		} else {
			a++
		}
		if b == len(word2[j])-1 {
			b = 0
			j++
		} else {
			b++
		}
	}
	// word1, word2 배열 끝까지 체크했는지
	if i == len(word1) && j == len(word2) {
		return true
	}
	return false
}

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
