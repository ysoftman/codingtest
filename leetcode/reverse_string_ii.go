/*
https://leetcode.com/problems/reverse-string-ii/
541. Reverse String II
Easy

Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.

If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and leave the other as original.

Example 1:
Input: s = "abcdefg", k = 2
Output: "bacdfeg"

Example 2:
Input: s = "abcd", k = 2
Output: "bacd"

Constraints:
1 <= s.length <= 104
s consists of only lowercase English letters.
1 <= k <= 104
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func reverseStr(s string, k int) string {
	r := make([]byte, len(s))
	copy(r, s)
	for i := 0; i < len(s); i += 2 * k {
		a := i
		b := min(i+k-1, len(r)-1)
		for a < b {
			temp := r[a]
			r[a] = r[b]
			a++
			r[b] = temp
			b--
		}
	}
	return string(r)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
}
