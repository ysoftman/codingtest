/*
https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
1047. Remove All Adjacent Duplicates In String
Easy

You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.

We repeatedly make duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.

Example 1:
Input: s = "abbaca"
Output: "ca"
Explanation:
For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".

Example 2:
Input: s = "azxxzy"
Output: "ay"

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters.
*/
package main

import "fmt"

/*
s="abbaca" 일때
stack top 과 s[i] 가 중복이면 stack pop, 다르면 push s[i] to stack...
결국 스택에 남은것이 인접한 중복 글자를 제외한 문자들만 남는다.
*/
// using stack
func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			// pop
			stack = stack[:len(stack)-1]
			continue
		}
		// push
		stack = append(stack, s[i])
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("azxxzy"))
	fmt.Println(removeDuplicates("aaa"))
	fmt.Println(removeDuplicates("aaaa"))
	fmt.Println(removeDuplicates("aaaabaaaa"))
}
