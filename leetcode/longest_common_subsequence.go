/*
https://leetcode.com/problems/longest-common-subsequence/
1143. Longest Common Subsequence
Medium
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

Example 1:
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
*/

package main

import "fmt"

/*
LongestCommonSubsequence(LCS)

같으면 왼쪽위([i-1][j-1])+1
아니면 max(왼쪽, 위)
   a b c d e
 a 1 1 1 1 1
 c 1 1 2 2 2
 e 1 1 2 2 3
*/
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			a, b := 0, 0
			if i > 0 {
				a = dp[i-1][j]
			}
			if j > 0 {
				b = dp[i][j-1]
			}
			c := 0
			if i > 0 && j > 0 {
				c = dp[i-1][j-1]
			}
			dp[i][j] = max(a, b)
			if text1[i] == text2[j] {
				dp[i][j] = c + 1
			}
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("bsbininm", "jmjkbkjkv"))
}
