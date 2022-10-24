/*
https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
239. Maximum Length of a Concatenated String with Unique Characters
Medium

You are given an array of strings arr. A string s is formed by the concatenation of a subsequence of arr that has unique characters.

Return the maximum possible length of s.

A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

Example 1:
Input: arr = ["un","iq","ue"]
Output: 4
Explanation: All the valid concatenations are:
- ""
- "un"
- "iq"
- "ue"
- "uniq" ("un" + "iq")
- "ique" ("iq" + "ue")
Maximum length is 4.

Example 2:
Input: arr = ["cha","r","act","ers"]
Output: 6
Explanation: Possible longest valid concatenations are "chaers" ("cha" + "ers") and "acters" ("act" + "ers").

Example 3:
Input: arr = ["abcdefghijklmnopqrstuvwxyz"]
Output: 26
Explanation: The only string in arr has all 26 characters.

Constraints:
1 <= arr.length <= 16
1 <= arr[i].length <= 26
arr[i] contains only lowercase English letters.
*/
package main

import "fmt"

// DFS 탐색
func recursiveMaxLen(arr []string, idx int, str string) int {
	m := make(map[byte]bool)
	// unique 문자만 있어야만 길이 비교 대상이 된다.
	for i := 0; i < len(str); i++ {
		m[str[i]] = true
	}
	if len(m) != len(str) {
		return 0
	}

	// unique 한 경우
	r := len(str)
	for i := idx; i < len(arr); i++ {
		temp := recursiveMaxLen(arr, i, str+arr[i])
		if r < temp {
			r = temp
		}
	}
	return r
}

func maxLength(arr []string) int {
	return recursiveMaxLen(arr, 0, "")
}

func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Println(maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
}
