/*
https://leetcode.com/problems/valid-anagram/
242. Valid Anagram
Easy
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
*/

package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashmap := make(map[rune]int, 0)
	for _, v := range s {
		hashmap[v]++
	}
	for _, v := range t {
		value, exist := hashmap[v]
		if !exist {
			return false
		}
		if value <= 0 {
			return false
		}
		hashmap[v]--
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("aacc", "ccac"))
}
