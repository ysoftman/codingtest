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

func sortDictionary(dict []string) {
	sort.Slice(dict, func(a, b int) bool {
		if len(dict[a]) > len(dict[b]) {
			return true
		}
		// 길이가 같으면 오름 차순
		if len(dict[a]) == len(dict[b]) {
			for i := 0; i < len(dict[a]); i++ {
				if dict[a][i] < dict[b][i] {
					return true
				}
				if dict[a][i] == dict[b][i] {
					continue
				}
				return false
			}
		}
		return false
	})
}

func findLongestWord(s string, dictionary []string) string {
	// 방법 1 - 사전순으로 정렬
	// sort.Strings(dictionary)
	// 방법 2 - 길이 우선, 같은 길이면 사전순(오름차순)으로 정렬
	sortDictionary(dictionary)
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
		// 방법1일 경우 dictionary 에 대해서 확인 필요
		// if dicPos == len(dictionary[i]) {
		// 	if len(r) < len(dictionary[i]) {
		// 		r = dictionary[i]
		// 	}
		// }
		// 방법2일 경우 첫번째 매칭되는 dictionary 가 정답
		if dicPos == len(dictionary[i]) {
			r = dictionary[i]
			break
		}
	}
	return r
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
	fmt.Println(findLongestWord("abce", []string{"abe", "abc"}))
	fmt.Println(findLongestWord("abce", []string{"zzzz", "abe", "abc"}))
	fmt.Println(findLongestWord("barfoofoobarthefoobarman", []string{"bbz", "bar", "bbr", "foo", "the"}))
}
