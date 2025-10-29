/*
https://leetcode.com/problems/keyboard-row/
500. Keyboard Row
Easy

Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.

In the American keyboard:

the first row consists of the characters "qwertyuiop",
the second row consists of the characters "asdfghjkl", and
the third row consists of the characters "zxcvbnm".

Example 1:
Input: words = ["Hello","Alaska","Dad","Peace"]
Output: ["Alaska","Dad"]

Example 2:
Input: words = ["omk"]
Output: []

Example 3:
Input: words = ["adsdf","sfd"]
Output: ["adsdf","sfd"]

Constraints:
1 <= words.length <= 20
1 <= words[i].length <= 100
words[i] consists of English letters (both lowercase and uppercase).
*/
package main

import (
	"fmt"
	"strings"
)

func findInRow(w, row string) bool {
	for i := 0; i < len(w); i++ {
		bFind := false
		for j := 0; j < len(row); j++ {
			if w[i] == row[j] {
				bFind = true
				break
			}
		}
		if !bFind {
			return false
		}
	}
	return true
}
func findWords(words []string) []string {
	row1 := "qwertyuiop"
	row2 := "asdfghjkl"
	row3 := "zxcvbnm"

	r := []string{}
	for i := 0; i < len(words); i++ {
		w := strings.ToLower(words[i])
		if findInRow(w, row1) || findInRow(w, row2) || findInRow(w, row3) {
			r = append(r, words[i])
		}
	}
	return r
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(findWords([]string{"omk"}))
	fmt.Println(findWords([]string{"adsdf", "sfd"}))
}
