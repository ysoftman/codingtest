/*
https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings
712. Minimum ASCII Delete Sum for Two Strings
Medium
Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.

Input: s1 = "sea", s2 = "eat"
Output: 231
Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
Deleting "t" from "eat" adds 116 to the sum.
At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.

Example 2:
Input: s1 = "delete", s2 = "leet"
Output: 403
Explanation: Deleting "dee" from "delete" to turn the string into "let",
adds 100[d] + 101[e] + 101[e] to the sum.
Deleting "e" from "leet" adds 101[e] to the sum.
At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.

Constraints:
1 <= s1.length, s2.length <= 1000
s1 and s2 consist of lowercase English letters.
*/
/*
a=1 b=2 c=3 으로 가정하고
s1 = abc   s2 = abb
  0 a b b
0 0 1 3 4
a 1 0 2 4
b 3 2 0 2
c 6 5 3 5

같으면 대각선 값을 가져오고
다르면 min(왼쪽+s2[j-],  위쪽+s1[i-1])
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i <= len(s2); i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if int(s1[i-1]) == int(s2[j-1]) {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = min((dp[i][j-1] + int(s2[j-1])), (dp[i-1][j] + int(s1[i-1])))
		}
	}

	// for debugging
	// for i := 0; i <= len(s1); i++ {
	// 	for j := 0; j <= len(s2); j++ {
	// 		fmt.Printf("%4d", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }
	return dp[len(s1)][len(s2)]
}

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
}
