/*
https://leetcode.com/problems/longest-string-chain/
1048. Longest String Chain
Medium
You are given an array of words where each word consists of lowercase English letters.
wordA is a predecessor of wordB if and only if we can insert exactly one letter anywhere in wordA without changing the order of the other characters to make it equal to wordB.
For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor of "bcad".
A word chain is a sequence of words [word1, word2, ..., wordk] with k >= 1, where word1 is a predecessor of word2, word2 is a predecessor of word3, and so on. A single word is trivially a word chain with k == 1.
Return the length of the longest possible word chain with words chosen from the given list of words.

Example 1:
Input: words = ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: One of the longest word chains is ["a","ba","bda","bdca"].

Example 2:
Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
Output: 5
Explanation: All the words can be put in a word chain ["xb", "xbc", "cxbc", "pcxbc", "pcxbcf"].

Example 3:
Input: words = ["abcd","dbqca"]
Output: 1
Explanation: The trivial word chain ["abcd"] is one of the longest word chains.
["abcd","dbqca"] is not a valid word chain because the ordering of the letters is changed.

Constraints:
1 <= words.length <= 1000
1 <= words[i].length <= 16
words[i] only consists of lowercase English letters.
*/
package main

import (
	"fmt"
	"sort"
)

/*
우선 길이기준으 오름차순 정렬
words = ["a","b","ba","bca","bda","bdca"]
words[i] 가 유효가 되려면 => words[0~i-1]에 문하나를 추가해 words[i] 되어야 한다.

a => 앞의 "" 에 a 추가 == a

b => 앞의 "" 에 b 추가 == b or
     앞의 a 에 b 추가 != b or

ba => 앞의 a 에 b 추가 == ba or
      앞의 b 에 a 추가 == ba

bca => 앞의 ba 에 c 추가 == bac

bda => 앞의 ba 에 d 추가 == bda

bdca => 앞의 bca 에 c 추가 != bdca or
        앞의 bda 에 c 추가 == bdca
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dynamic programming
func longestStrChain(words []string) int {
	sort.Slice(words, func(a, b int) bool {
		return len(words[a]) < len(words[b])
	})
	// fmt.Println(words)
	max_result := 0
	m := make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		m[words[i]] = 1
	}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			// j 번째 문자를 뺀 문자열이 기존에 있었는지 확인
			preWord := words[i][:j] + words[i][j+1:]
			if val, ok := m[preWord]; ok {
				// preWord(보다 현재 word 길이가 +1 임) 중 가장 큰 길이 기록
				m[words[i]] = max(m[words[i]], val+1)
			}
		}
		max_result = max(max_result, m[words[i]])
	}
	return max_result
}

func main() {
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
	fmt.Println(longestStrChain([]string{"abcd", "dbqca"}))
}
