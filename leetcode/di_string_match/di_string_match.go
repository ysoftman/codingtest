/*
https://leetcode.com/problems/di-string-match
Q67. DI String Match
Easy
A permutation perm of n + 1 integers of all the integers in the range [0, n] can be represented as a string s of length n where:

s[i] == 'I' if perm[i] < perm[i + 1], and
s[i] == 'D' if perm[i] > perm[i + 1].
Given a string s, reconstruct the permutation perm and return it. If there are multiple valid permutations perm, return any of them.

Example 1:
Input: s = "IDID"
Output: [0,4,1,3,2]

Example 2:
Input: s = "III"
Output: [0,1,2,3]

Example 3:
Input: s = "DDI"
Output: [3,2,0,1]

Constraints:
1 <= s.length <= 105
s[i] is either 'I' or 'D'.
*/
package main

import "fmt"

func diStringMatch(s string) []int {
	r := []int{}
	I := 0
	D := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			r = append(r, I)
			I++
		} else {
			r = append(r, D)
			D--
		}
	}
	r = append(r, I)
	return r
}

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
}
