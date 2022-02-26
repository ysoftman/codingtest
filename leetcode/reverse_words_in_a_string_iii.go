/*
https://leetcode.com/problems/reverse-words-in-a-string-iii/

557. Reverse Words in a String III
Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Example 2:
Input: s = "God Ding"
Output: "doG gniD"
*/

package main

import (
	"fmt"
	"strings"
)

func reverseWord(word string) string {
	s := []byte(word)
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = reverseWord(words[i])
	}
	return strings.Join(words, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
	s = "God Ding"
	fmt.Println(reverseWords(s))
}
