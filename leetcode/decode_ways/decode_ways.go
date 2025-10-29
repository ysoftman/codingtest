/*
https://leetcode.com/problems/decode-ways/
91. Decode Ways
Medium
A message containing letters from A-Z can be encoded into numbers using the following mapping:
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:

"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.
The test cases are generated so that the answer fits in a 32-bit integer.

Example 1:
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).

Example 2:
Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Example 3:
Input: s = "06"
Output: 0
Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").

Constraints:
1 <= s.length <= 100
s contains only digits and may contain leading zero(s).
*/

package main

import "fmt"

/*
dp[0] = 1
1 -> 1      dp[1] = dp[1-1]
11 -> 2     dp[2] = dp[2-1] + dp[2-2]
111 -> 3    dp[3] = dp[3-1] + dp[3-2]
1111 -> 5   dp[4] = dp[4-1] + dp[4-2]


if i-1 == 0, skip adding dp[i-1]
if (i-2 i-2) 10~26, adding dp[i-2]
110 -> 1    dp[3] = dp[3-2]
*/
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	if s[0] != '0' {
		dp[0] = 1
	}
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if i < 2 {
			continue
		}
		if (s[i-2] == '1' && s[i-1] >= '0' && s[i-1] <= '9') ||
			(s[i-2] == '2' && s[i-1] >= '0' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("06"))
	fmt.Println(numDecodings("10"))
	fmt.Println(numDecodings("2101"))
	fmt.Println(numDecodings("111111111111111111111111111111111111111111111"))
	fmt.Println(numDecodings("27"))
	fmt.Println(numDecodings("00005"))
	fmt.Println(numDecodings("190"))
	fmt.Println(numDecodings("2067"))
}
