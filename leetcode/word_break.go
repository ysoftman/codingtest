/*
https://leetcode.com/problems/word-break/
139. Word Break
Medium
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false

Constraints:
1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.
*/
package main

import "fmt"

/*
dynamic programming
leetcode (s)
01234567 (i)
dp[0] == true  && s[0:0] == x ... > dp[0] = true
dp[1] == false ...  > dp[1] = false
dp[2] == false ...  > dp[2] = false
dp[3] == false ...  > dp[3] = false
dp[0] == true && s[0:4] == leet > dp[4] = true
dp[4] == true && s[4:8] == code > dp[8] = true
return dp[8]
*/
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// fmt.Printf("%v\n", dp[i])
			if !dp[j] {
				continue
			}
			if _, exist := m[s[j:i]]; exist {
				dp[i] = true
				break
			}
		}
	}
	// fmt.Println(dp)
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
