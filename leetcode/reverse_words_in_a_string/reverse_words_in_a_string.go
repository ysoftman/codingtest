/*
https://leetcode.com/problems/reverse-words-in-a-string/
151. Reverse Words in a String
Medium
Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

Example 1:
Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

Constraints:
1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/
package main

import (
	"fmt"
	"strings"
)

// string 기본 패키지 사용시
// time complexity: O(N)
// space complexity: O(N)
func reverseWords2(s string) string {
	r := ""
	s = strings.Trim(s, " ")
	temp := strings.Split(s, " ")
	r += strings.Trim(temp[len(temp)-1], " ")
	for i := len(temp) - 2; i >= 0; i-- {
		word := strings.Trim(temp[i], " ")
		if len(word) > 0 {
			r += " " + word
		}
	}
	return r
}

// time complexity: O(N)
// space complexity: O(1)
func reverseWords(s string) string {
	r := ""
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
			if i < len(s)-1 {
				continue
			}
		}
		if s[i] == ' ' && len(word) == 0 {
			continue
		}
		r = word + " " + r
		word = ""
	}
	return r[:len(r)-1]
}
func main() {
	fmt.Println(reverseWords("the sky is blue") + "|")
	fmt.Println(reverseWords("  hello world  ") + "|")
	fmt.Println(reverseWords("a good   example") + "|")
}
