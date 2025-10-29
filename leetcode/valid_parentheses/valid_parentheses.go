/*
https://leetcode.com/problems/valid-parentheses/
20. Valid Parentheses
Easy
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.


Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false
*/

package main

import "fmt"

func isValid(s string) bool {
	temp := []rune{}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			// push
			temp = append(temp, v)
		} else {
			// pop
			if len(temp) == 0 {
				return false
			}
			if temp[len(temp)-1] == '(' && v == ')' {
				temp = temp[:len(temp)-1]
			} else if temp[len(temp)-1] == '{' && v == '}' {
				temp = temp[:len(temp)-1]
			} else if temp[len(temp)-1] == '[' && v == ']' {
				temp = temp[:len(temp)-1]
			} else {
				return false
			}
		}
	}
	if len(temp) == 0 {
		return true
	}
	return false
}

func main() {
	s := "()"
	fmt.Println(isValid(s))

	s = "()[]{}"
	fmt.Println(isValid(s))

	s = "(]"
	fmt.Println(isValid(s))
}
