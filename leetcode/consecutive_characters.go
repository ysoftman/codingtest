/*
https://leetcode.com/problems/consecutive-characters/
1446. Consecutive Characters
Easy
The power of the string is the maximum length of a non-empty substring that contains only one unique character.
Given a string s, return the power of s.

Example 1:
Input: s = "leetcode"
Output: 2
Explanation: The substring "ee" is of length 2 with the character 'e' only.

Example 2:
Input: s = "abbcccddddeeeeedcba"
Output: 5
Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
*/

package main

import "fmt"

func maxPower(s string) int {
	max := -1
	length := 0
	var pre rune
	for _, v := range s {
		length++

		if pre != v {
			length = 1
		}
		pre = v
		if max < length {
			max = length
		}
	}
	return max
}
func main() {
	fmt.Println(maxPower("leetcode"))
	fmt.Println(maxPower("abbcccddddeeeeedcba"))
}
