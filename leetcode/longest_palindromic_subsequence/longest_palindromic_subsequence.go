/*
https://leetcode.com/problems/longest-palindromic-subsequence/
516. Longest Palindromic Subsequence
Medium

Given a string s, find the longest palindromic subsequence's length in s.

A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

Example 1:
Input: s = "bbbab"
Output: 4
Explanation: One possible longest palindromic subsequence is "bbbb".

Example 2:
Input: s = "cbbd"
Output: 2
Explanation: One possible longest palindromic subsequence is "bb".

Constraints:
1 <= s.length <= 1000
s consists only of lowercase English letters.
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
bottom-up dynamic programming - LPS(longest palindromic subsequence)
2d memoization 을 생성,

문자 매칭되면 왼쪽아래값(이전에서 가장 긴 LPS)+2(매칭됐으니 2개의 문자길이가 추가된것)
문자 매칭 안되면  max(왼쪽값, 아래쪽값)
... 마지막 0번째 라인의 끝이 최대 LPS 길이가 된다.

_ b b b a b
b 1 2 3 3 4
b 0 1 2 2 3
b 0 0 1 1 3
a 0 0 0 1 1
b 0 0 0 0 1
*/
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
}
