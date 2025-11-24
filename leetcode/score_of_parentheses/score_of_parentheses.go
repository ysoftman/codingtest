/*
https://leetcode.com/problems/score-of-parentheses/submissions
856. Score of Parentheses
Medium
Given a balanced parentheses string s, return the score of the string.

The score of a balanced parentheses string is based on the following rule:

"()" has score 1.
AB has score A + B, where A and B are balanced parentheses strings.
(A) has score 2 * A, where A is a balanced parentheses string.

Example 1:
Input: s = "()"
Output: 1

Example 2:
Input: s = "(())"
Output: 2

Example 3:
Input: s = "()()"
Output: 2

Constraints:
2 <= s.length <= 50
s consists of only '(' and ')'.
s is a balanced parentheses string.
*/
package main

import "fmt"

func scoreOfParentheses(s string) int {
	m := make(map[int]int)
	st := 0
	for i := range s {
		if s[i] == '(' {
			st++
		} else if s[i] == ')' {
			if s[i-1] == ')' {
				m[st] += m[st+1] * 2
				m[st+1] = 0
			} else {
				m[st] += 1
			}
			st--
		}
	}
	// fmt.Println(m)
	return m[1]
}

func main() {
	fmt.Println(scoreOfParentheses("()"))
	fmt.Println(scoreOfParentheses("(())"))
	fmt.Println(scoreOfParentheses("()()"))
	fmt.Println(scoreOfParentheses("(()())()()()((()))"))
	fmt.Println(scoreOfParentheses("((()))"))
	fmt.Println(scoreOfParentheses("(()(()))"))
	fmt.Println(scoreOfParentheses("(()())()()()((()))(((())))()()()(())(())"))
}
