/*
https://leetcode.com/problems/longest-palindrome/
409. Longest Palindrome
Easy

Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

Example 1:
Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

Example 2:
Input: s = "a"
Output: 1
Explanation: The longest palindrome that can be built is "a", whose length is 1.

Constraints:
1 <= s.length <= 2000
s consists of lowercase and/or uppercase English letters only.
*/

package main

import "fmt"

/*
ccc 의 경우 c=3개
cc 짝수 2개 보장 r+=2
r은 짝수(팰린드롬 보장)라면 c가 1개남은것 더할 수 있음
r += 1
r = 3
*/
func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	r := 0
	for _, v := range m {
		// 문자가 짝수개수있으면 palindrome 을 만들 수 있음으로 우선 짝수개수만큼 보장
		r += (v / 2) * 2
		// 현재 결과가 짝수개(패린드롬인 상태)라면 홀수인 v 의 1개를 더할 수 있다.
		if r%2 == 0 && v%2 == 1 {
			r += 1
		}
	}
	return r
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ccc"))
}
