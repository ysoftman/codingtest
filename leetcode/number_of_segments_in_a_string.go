/*
https://leetcode.com/problems/number-of-segments-in-a-string/
434. Number of Segments in a String
Easy

Given a string s, return the number of segments in the string.
A segment is defined to be a contiguous sequence of non-space characters.

Example 1:
Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]

Example 2:
Input: s = "Hello"
Output: 1

Constraints:
0 <= s.length <= 300
s consists of lowercase and uppercase English letters, digits, or one of the following characters "!@#$%^&*()_+-=',.:".
The only space character in s is ' '.
*/
package main

import "fmt"

func countSegments(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		// 공백만 검색하면 안된다.
		// i=0일때 문자가 있거나
		// 공백 이전 글자가 공백일때 새로운 세그먼트 시작
		if i == 0 || s[i-1] == ' ' {
			cnt++
		}
	}
	return cnt
}
func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments("aaa"))
	fmt.Println(countSegments(""))
	fmt.Println(countSegments("aksdjfasdkfj klj asdkljfsdl,ffjkdsa jdfsajfkj,jfdskjfksadjk"))
	fmt.Println(countSegments("                "))
}
