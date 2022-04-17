/*
https://leetcode.com/problems/interleaving-string/
97. Interleaving String
Medium
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:
s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.

Example 1:
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true

Example 2:
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false

Example 3:
Input: s1 = "", s2 = "", s3 = ""
Output: true

Constraints:
0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1, s2, and s3 consist of lowercase English letters.

Follow up: Could you solve it using only O(s2.length) additional memory space?
*/
package main

import "fmt"

// brute force, O(2^m+n), time limit exceeded!
func recursiveIsInterleave(s1, s2, s3 string, s1i, s2i, s3i int) bool {
	if s3i == len(s3) {
		if s1i == len(s1) && s2i == len(s2) {
			// fmt.Println("find!")
			return true
		}
		return false
	}

	a, b := false, false
	if s1i < len(s1) && s3[s3i] == s1[s1i] {
		a = recursiveIsInterleave(s1, s2, s3, s1i+1, s2i, s3i+1)
	}
	if s2i < len(s2) && s3[s3i] == s2[s2i] {
		b = recursiveIsInterleave(s1, s2, s3, s1i, s2i+1, s3i+1)
	}
	return a || b
}
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return recursiveIsInterleave(s1, s2, s3, 0, 0, 0)
}
func printBoolMatrix(matrix [][]bool) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := 'F'
			if matrix[i][j] {
				c = 'T'
			}
			fmt.Printf("%c", c)
			if j != n-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

/*

dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])

aadbbcbcac(s3)

    d b b c a(s2)
  T F F F F F
a T F F F F F
a T T T T T F
b F T T F T F
c F F T T T T
c F F F T F T
(s1)
*/
// 2D dynamic programming
func isInterleave2DP(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	// printBoolMatrix(dp)
	return dp[len(s1)][len(s2)]
}

// 1D dynamic programming
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[j] = true
				continue
			}
			if i == 0 {
				dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) || (dp[j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[len(s2)]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
	fmt.Println(isInterleave("abababababababababababababababababababababababababababababababababababababababababababababababababbb", "abababababababababababababababababababababababababababababababababababababababababababababababababab", "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb"))
}
