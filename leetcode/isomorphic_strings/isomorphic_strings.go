/*
https://leetcode.com/problems/isomorphic-strings/
205. Isomorphic Strings
Easy
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Example 1:
Input: s = "egg", t = "add"
Output: true

Example 2:
Input: s = "foo", t = "bar"
Output: false

Example 3:
Input: s = "paper", t = "title"
Output: true

Constraints:
1 <= s.length <= 5 * 104
t.length == s.length
s and t consist of any valid ascii character.
*/

package main

import "fmt"

// isomorphic (동형의) string
// egg add => 두 단어 모두 첫문자 이후 두번째 문자 2번 반복 같은 패턴으로 isomorphic
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// s -> t 매핑 테이블 생성
	m := make(map[byte]byte)
	m2 := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		// t[i] 이미 매핑된 값이면 매핑하지 않기(중복 매핑 금지)
		if _, exist := m2[t[i]]; exist {
			continue
		}
		m[s[i]] = t[i]
		m2[t[i]] = true
	}
	// s 매핑 테이블 값으로 t 가 되는지 체크
	for i := 0; i < len(s); i++ {
		if m[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
}
