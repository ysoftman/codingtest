/*
https://leetcode.com/problems/delete-operation-for-two-strings/
583. Delete Operation for Two Strings
Medium
Share
Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.
In one step, you can delete exactly one character in either string.

Example 1:
Input: word1 = "sea", word2 = "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".

Example 2:
Input: word1 = "leetcode", word2 = "etco"
Output: 4

Constraints:
1 <= word1.length, word2.length <= 500
word1 and word2 consist of only lowercase English letters.
*/
package main

import "fmt"

/*
LongestCommonSubsequence(LCS)

같으면 왼쪽위([i-1][j-1])+1
아니면 max(왼쪽, 위)
    e a t
  0 0 0 0
s 0 0 0 0
e 0 1 1 1
a 0 1 2 2

lcs = 2
result = (word1 - lcs) +  (word2 - lcs)
*/

// for debugging
func PrintDP(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			fmt.Printf("%v ", dp[i][j])
		}
		fmt.Println()
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// for debugging
	// PrintDP(dp)
	lcs := dp[len(word1)][len(word2)]
	return (len(word1) - lcs) + (len(word2) - lcs)
}

func main() {
	fmt.Println(minDistance("sea", "eat"))
	fmt.Println(minDistance("leetcode", "etco"))
	fmt.Println(minDistance("leetcode", "etcoaaaaaaaaaa"))
}
