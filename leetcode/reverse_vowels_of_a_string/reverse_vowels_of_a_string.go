/*
https://leetcode.com/problems/reverse-vowels-of-a-string/
345. Reverse Vowels of a String
Easy

Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.

Example 1:
Input: s = "hello"
Output: "holle"

Example 2:
Input: s = "leetcode"
Output: "leotcede"

Constraints:
1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/
package main

import "fmt"

func isVowel(c byte) bool {
	// 소문자로 변환
	if c < 'a' {
		c += ('a' - 'A')
	}
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}

// brute-force
// time complexity: O(n)
// space complexity: O(n)
func reverseVowels2(s string) string {
	vowels := []byte{}
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if isVowel(b[i]) {
			vowels = append(vowels, b[i])
		}
	}
	for i := 0; i < len(s); i++ {
		if isVowel(b[i]) {
			b[i] = vowels[len(vowels)-1]
			vowels = vowels[:len(vowels)-1]
		}
	}
	return string(b)
}

// two pointers
// time complexity: O(n)
// space complexity: O(1)
func reverseVowels(s string) string {
	b := []byte(s)
	start := 0
	end := len(b) - 1
	for start < end {
		for !isVowel(b[start]) && start < end {
			start++
		}
		for !isVowel(b[end]) && start < end {
			end--
		}
		// fmt.Println(start, end)
		// swap
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
	return string(b)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("aA"))
}
