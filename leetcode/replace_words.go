/*
https://leetcode.com/problems/replace-words
648. Replace Words
Medium

In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word derivative. For example, when the root "help" is followed by the word "ful", we can form a derivative "helpful".

Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, replace all the derivatives in the sentence with the root forming it. If a derivative can be replaced by more than one root, replace it with the root that has the shortest length.

Return the sentence after the replacement.

Example 1:
Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"

Example 2:
Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
Output: "a a b c"
*/
package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	m := make(map[string]bool)
	for i := 0; i < len(dictionary); i++ {
		m[dictionary[i]] = true
	}

	r := []string{}
	replaced := false
	sen := strings.Split(sentence, " ")
	for _, word := range sen {
		replaced = false
		for i := 1; i < len(word); i++ {
			if exist := m[word[:i]]; exist {
				r = append(r, word[:i])
				replaced = true
				break
			}
		}
		if !replaced {
			r = append(r, word)
		}
	}

	return strings.Join(r, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}
