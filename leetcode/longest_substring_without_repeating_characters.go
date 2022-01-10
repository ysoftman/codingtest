/*
// https: //leetcode.com/problems/longest-substring-without-repeating-characters/
3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	temp := ""
	max := 0
	for i := 0; i < len(s); i++ {
		for j := range temp {
			if temp[j] == s[i] {
				if max < len(temp) {
					max = len(temp)
				}
				i -= len(temp) - 1
				i += j
				temp = ""
				break
			}
		}
		temp += string(s[i])
	}
	if max < len(temp) {
		max = len(temp)
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
