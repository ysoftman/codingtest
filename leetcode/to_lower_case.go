/*
https://leetcode.com/problems/to-lower-case/
709. To Lower Case
Easy
Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.

Example 1:
Input: s = "Hello"
Output: "hello"

Example 2:
Input: s = "here"
Output: "here"

Example 3:
Input: s = "LOVELY"
Output: "lovely"

Constraints:

1 <= s.length <= 100
s consists of printable ASCII characters.
*/
package main

import "fmt"

func toLowerCase(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			r += string('a' + (s[i] - 'A'))
		} else {
			r += string(s[i])
		}
	}
	return r
}

func main() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
	fmt.Println(toLowerCase("al&phaBET"))
}
