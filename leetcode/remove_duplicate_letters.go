/*
https://leetcode.com/problems/remove-duplicate-letters/
Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is the smallest in lexicographical order among all possible results.

Example 1:
Input: s = "bcabc"
Output: "abc"

Example 2:
Input: s = "cbacdcbc"
Output: "acdb"

Constraints:
1 <= s.length <= 104
s consists of lowercase English letters.

Note: This question is the same as 1081: https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
*/
package main

import "fmt"

func removeDuplicateLetters(s string) string {
	// 문자별 발생 빈도 파악
	charCnt := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charCnt[s[i]]++
	}
	addedChar := make(map[byte]bool)
	result := ""
	for i := 0; i < len(s); i++ {
		charCnt[s[i]]--
		// 이미 추가된 문자는 skip
		if addedChar[s[i]] {
			continue
		}
		// result 를 stack 으로 생각하고 처리, lexicographical order(사전식 우선순위)로 생성할 수 있다.
		// result 에 추가된 마지막가 중복이면 새로 추가할 문자 s[i] 보다 커야 한다.
		// result 의 마지막 문자 카운트가 남아 있고(중복 문자이고) result 의 마지막 문자가 s[i] 보다 크면
		for len(result) > 0 && charCnt[result[len(result)-1]] > 0 && s[i] < result[len(result)-1] {
			// result 에서 제거된 문자는 다시 탐색(추가)될 수 있도록 표시
			addedChar[result[len(result)-1]] = false
			// result 의 마지막 문자 제거(stack pop)
			result = result[:len(result)-1]
		}
		result += string(s[i])
		// 추가된 문자임을 체크하여 다음에 스킵할 수 있도록 한다.
		addedChar[s[i]] = true
	}
	return result
}

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
	fmt.Println(removeDuplicateLetters("abca"))
	fmt.Println(removeDuplicateLetters("bcaba"))
}
