/*
https://leetcode.com/problems/ransom-note/
383. Ransom Note
Easy
Given two strings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:
Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:
Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	hashmap := make(map[rune]uint, 0)
	for _, v := range magazine {
		hashmap[v]++
	}
	// fmt.Println(hashmap)
	for _, v := range ransomNote {
		value, exist := hashmap[v]
		if !exist {
			return false
		}
		if value == 0 {
			return false
		}
		hashmap[v]--
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
