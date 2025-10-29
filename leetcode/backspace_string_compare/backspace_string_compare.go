/*
https://leetcode.com/problems/backspace-string-compare/
844. Backspace String Compare
Easy

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
Note that after backspacing an empty text, the text will continue empty.

Example 1:
Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".

Constraints:
1 <= s.length, t.length <= 200
s and t only contain lowercase letters and '#' characters.

Follow up: Can you solve it in O(n) time and O(1) space?
*/

package main

import "fmt"

func removeChars(s string) []byte {
	s2 := []byte{}
	removeCnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			removeCnt++
			continue
		}
		if removeCnt > 0 {
			removeCnt--
			continue
		}
		s2 = append(s2, s[i])
	}
	return s2
}

func removeCharsUsingStack(s string) []byte {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		// pop
		if s[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, s[i])
	}
	return stack
}

// time complexity O(M+N)
// space complexity O(M+N)
func backspaceCompare2(s string, t string) bool {
	// s2 := removeChars(s)
	// t2 := removeChars(t)
	s2 := removeCharsUsingStack(s)
	t2 := removeCharsUsingStack(t)
	if len(s2) != len(t2) {
		return false
	}
	for i := 0; i < len(s2); i++ {
		if s2[i] != t2[i] {
			return false
		}
	}
	return true
}

// using two pointers
// time complexity O(M+N)
// space complexity O(1)
func backspaceCompare(s string, t string) bool {
	removeCntS := 0
	removeCntT := 0
	j := len(t) - 1
	i := len(s) - 1
	for i >= 0 {
		removeCntS = 0
		for i >= 0 {
			if s[i] == '#' {
				removeCntS++
			} else if removeCntS > 0 {
				removeCntS--
			} else {
				break
			}
			i--
		}

		removeCntT = 0
		for j >= 0 {
			if t[j] == '#' {
				removeCntT++
			} else if removeCntT > 0 {
				removeCntT--
			} else {
				break
			}
			j--
		}

		// fmt.Println(i, j)
		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}
		// i >=0 이면 j>=0, j<0 이면 j<0 이어야 한다.
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}

	return true
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#c", "b"))
	fmt.Println(backspaceCompare("bxj##tw", "bxj###tw"))
}
