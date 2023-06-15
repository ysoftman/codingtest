/*
https://leetcode.com/problems/extra-characters-in-a-string/
2707. Extra Characters in a String
Medium
You are given a 0-indexed string s and a dictionary of words dictionary.
You have to break s into one or more non-overlapping substrings such that each substring is present in dictionary.
There may be some extra characters in s which are not present in any of the substrings.

Return the minimum number of extra characters left over if you break up s optimally.

Example 1:
Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
Output: 1
Explanation:
We can break s in two substrings: "leet" from index 0 to 3 and "code" from index 5 to 8.
There is only 1 unused character (at index 4), so we return 1.

Example 2:
Input: s = "sayhelloworld", dictionary = ["hello","world"]
Output: 3
Explanation: We can break s in two substrings: "hello" from index 3 to 7 and "world" from index 8 to 12.
The characters at indices 0, 1, 2 are not used in any substring and thus are considered as extra characters.
Hence, we return 3.

Constraints:
1 <= s.length <= 50
1 <= dictionary.length <= 50
1 <= dictionary[i].length <= 50
dictionary[i] and s consists of only lowercase English letters
dictionary contains distinct words
*/
package main

import (
	"fmt"
)

func recursive(s string, m map[string]int, index int, dp []int) int {
	if index >= len(s) {
		return 0
	}
	// dp(dynamic programming / memoization) 하지 않으면 time limit exceeded!!!
	if dp[index] != -100 {
		return dp[index]
	}
	curStr := ""
	minExtra := len(s)
	for i := index; i < len(s); i++ {
		curStr += (string)(s[i])
		curExtra := 0
		// 사전에 현재 문자열이 없으면, 현재 문자열전체가 extra 가 된다.
		if _, ok := m[curStr]; !ok {
			curExtra = len(curStr)
		}
		// i+1 이후로 같은 과정을 반복(recursive)
		nextExtra := recursive(s, m, i+1, dp)
		// 현재 extra 개수
		totalExtra := curExtra + nextExtra
		if minExtra > totalExtra {
			minExtra = totalExtra
		}
	}
	dp[index] = minExtra
	return minExtra
}

func minExtraChar(s string, dictionary []string) int {
	dp := make([]int, len(s))
	// dp 에 값이 설정되지 않았음을 표시하기 위해 -100 값을 사용
	for i := 0; i < len(dp); i++ {
		dp[i] = -100
	}
	m := make(map[string]int, 0)
	for _, dic := range dictionary {
		m[dic]++
	}
	return recursive(s, m, 0, dp)
}

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
	fmt.Println(minExtraChar("dwmodizxvvbosxxw", []string{"ox", "lb", "diz", "gu", "v", "ksv", "o", "nuq", "r", "txhe", "e", "wmo", "cehy", "tskz", "ds", "kzbu"}))
}
