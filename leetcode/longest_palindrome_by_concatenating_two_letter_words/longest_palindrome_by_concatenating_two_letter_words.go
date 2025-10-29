/*
https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
2131. Longest Palindrome by Concatenating Two Letter Words
Medium

You are given an array of strings words. Each element of words consists of two lowercase English letters.
Create the longest possible palindrome by selecting some elements from words and concatenating them in any order. Each element can be selected at most once.
Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome, return 0.
A palindrome is a string that reads the same forward and backward.

Example 1:
Input: words = ["lc","cl","gg"]
Output: 6
Explanation: One longest palindrome is "lc" + "gg" + "cl" = "lcggcl", of length 6.
Note that "clgglc" is another longest palindrome that can be created.

Example 2:
Input: words = ["ab","ty","yt","lc","cl","ab"]
Output: 8
Explanation: One longest palindrome is "ty" + "lc" + "cl" + "yt" = "tylcclyt", of length 8.
Note that "lcyttycl" is another longest palindrome that can be created.

Example 3:
Input: words = ["cc","ll","xx"]
Output: 2
Explanation: One longest palindrome is "cc", of length 2.
Note that "ll" is another longest palindrome that can be created, and so is "xx".

Constraints:
1 <= words.length <= 105
words[i].length == 2
words[i] consists of lowercase English letters.
*/

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func reverseWord(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[len(s)-1-i] = s[i]
	}
	return string(r)
}

func longestPalindrome(words []string) int {
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	central := false
	r := 0
	for k, v := range m {
		rw := reverseWord(k)
		// aa, bb 같은 워드는
		if k == rw {
			// 짝수개는 결과에 카운트
			if v%2 == 0 {
				r += v
			} else {
				r += v - 1 // 홀수개는 1개 빼서 짝수개로 카운트
				// aa, aa, aa 가 있다면 aa 2개는 양쪽 끝에서 사용, aa 1개는 중앙에 사용 할 수 있다.
				// bb, bb, bb 가 더 있어도 중앙에 추가되는 것은 1번만 될 수 있음
				central = true
			}
			continue
		}
		// k=ab 일때 rw=ba 가 있다면
		if v2, ok := m[rw]; ok {
			// ab 인 경우로 카운트, ab 에서 카운트했으니 ba 하지 않음
			if k[0] < k[1] {
				// ab, ab, ba 일때 m[ab]=2(v), m[ba]=1(v2) 이고 v2 를 선택
				r += 2 * min(v, v2)
			}
		}
	}
	// 중앙에 하나더 위치할 수 있는 경우
	if central {
		r++
	}
	return r * 2
}

func main() {
	fmt.Println(longestPalindrome([]string{"lc", "cl", "gg"}))
	fmt.Println(longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	fmt.Println(longestPalindrome([]string{"cc", "ll", "xx"}))
}
