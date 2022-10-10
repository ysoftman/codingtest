/*
https://leetcode.com/problems/break-a-palindrome/
1328. Break a Palindrome
Medium

Given a palindromic string of lowercase English letters palindrome, replace exactly one character with any lowercase English letter so that the resulting string is not a palindrome and that it is the lexicographically smallest one possible.

Return the resulting string. If there is no way to replace a character to make it not a palindrome, return an empty string.

A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ, a has a character strictly smaller than the corresponding character in b. For example, "abcc" is lexicographically smaller than "abcd" because the first position they differ is at the fourth character, and 'c' is smaller than 'd'.

Example 1:
Input: palindrome = "abccba"
Output: "aaccba"
Explanation: There are many ways to make "abccba" not a palindrome, such as "zbccba", "aaccba", and "abacba".
Of all the ways, "aaccba" is the lexicographically smallest.

Example 2:
Input: palindrome = "a"
Output: ""
Explanation: There is no way to replace a single character to make "a" not a palindrome, so return an empty string.


Constraints:
1 <= palindrome.length <= 1000
palindrome consists of only lowercase English letters.
*/
package main

import "fmt"

// palindrom 이 아니게 만드는 방법중, 사전적으로 가장 smallest 방법을 리턴
func breakPalindrome(palindrome string) string {
	// a 와 같이 길이가 1개이면 palindrome 이 아니라서 empty string 리턴
	if len(palindrome) == 1 {
		return ""
	}
	// palindrom 이 아니게 만들때 중간 길이중에 하나를 변경하면 된다.
	for i := 0; i < len(palindrome)/2; i++ {
		// 사전전으로 a 로 바꾸는것이 smallest 방법이다.
		// 따라서 a 가 아닌것이 있다면 a로 바꿔서 리턴하면 smallest 방법이다.
		if palindrome[i] != 'a' {
			r := ""
			r += palindrome[:i]
			r += "a"
			r += palindrome[i+1:]
			return r
		}
	}
	// 모두 a 였다면 마지막 문자만 b 로 바꾸는것이 smallest 방법이다.
	r := palindrome[:len(palindrome)-1]
	r += "b"
	return r
}

func main() {
	fmt.Println(breakPalindrome("abccba"))
	fmt.Println(breakPalindrome("a"))
}
