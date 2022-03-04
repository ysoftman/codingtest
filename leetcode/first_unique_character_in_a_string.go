/*
https://leetcode.com/problems/first-unique-character-in-a-string/
387. First Unique Character in a String
Easy

Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Example 1:
Input: s = "leetcode"
Output: 0

Example 2:
Input: s = "loveleetcode"
Output: 2

Example 3:
Input: s = "aabb"
Output: -1
*/
package main

import "fmt"

func firstUniqChar(s string) int {
	hashmap := make(map[byte]uint, 0)
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if hashmap[s[i]] == 1 {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
