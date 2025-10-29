/*
https://leetcode.com/problems/group-anagrams/
49. Group Anagrams
Medium
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]


Constraints:
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(str string) string {
	r := []string{}
	for _, v := range str {
		r = append(r, string(v))
	}
	sort.Strings(r)
	return strings.Join(r, ",")
}
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, 0)
	for _, v := range strs {
		s := sortString(v)
		m[s] = append(m[s], v)
	}
	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println(groupAnagrams([]string{""}))
}
