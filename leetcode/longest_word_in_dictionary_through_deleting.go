/*
https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/
524. Longest Word in Dictionary through Deleting
Medium
Given a string s and a string array dictionary, return the longest string in the dictionary that can be formed by deleting some of the given string characters. If there is more than one possible result, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.

Example 1:
Input: s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
Output: "apple"

Example 2:
Input: s = "abpcplea", dictionary = ["a","b","c"]
Output: "a"

Constraints:
1 <= s.length <= 1000
1 <= dictionary.length <= 1000
1 <= dictionary[i].length <= 1000
s and dictionary[i] consist of lowercase English letters.
*/
package main

import (
	"fmt"
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	// 사전순으로 정렬
	sort.Strings(dictionary)
	// fmt.Println(dictionary)
	r := ""
	for i := 0; i < len(dictionary); i++ {
		dicPos := 0
		for j := 0; j < len(s); j++ {
			if dicPos < len(dictionary[i]) && s[j] == dictionary[i][dicPos] {
				dicPos++
			}
			// s 남은 길이가 dictionary[i] 의 남은 길이보다 작으면 더 확인할필요 없다.
			if len(s)-j < len(dictionary[i])-dicPos {
				break
			}
		}
		if dicPos == len(dictionary[i]) {
			if len(r) < len(dictionary[i]) {
				r = dictionary[i]
			}
		}
	}
	return r
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
	fmt.Println(findLongestWord("abce", []string{"abe", "abc"}))
}
