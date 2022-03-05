/*
https://leetcode.com/problems/longest-common-prefix/
14. Longest Common Prefix
Easy
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minStr := strs[0]
	for _, v := range strs {
		if len(minStr) > len(v) {
			minStr = v
		}
	}

	commonPrefix := []byte{}
	for i := range minStr {
		for _, v := range strs {
			if minStr[i] != v[i] {
				return string(commonPrefix)
			}
		}
		commonPrefix = append(commonPrefix, minStr[i])
	}

	return string(commonPrefix)
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
