/*
https://leetcode.com/problems/occurrences-after-bigram/
1078. Occurrences After Bigram
Easy
Given two strings first and second, consider occurrences in some text of the form "first second third", where second comes immediately after first, and third comes immediately after second.
Return an array of all the words third for each occurrence of "first second third".

Example 1:
Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
Output: ["girl","student"]

Example 2:
Input: text = "we will we will rock you", first = "we", second = "will"
Output: ["we","rock"]

Constraints:
1 <= text.length <= 1000
text consists of lowercase English letters and spaces.
All the words in text a separated by a single space.
1 <= first.length, second.length <= 10
first and second consist of lowercase English letters.
*/
package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	results := []string{}
	strs := strings.Split(text, " ")
	for i := 2; i < len(strs); i++ {
		if strs[i-2] == first && strs[i-1] == second {
			results = append(results, strs[i])
		}
	}
	return results
}

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
	fmt.Println(findOcurrences("we we we we will rock you", "we", "we"))
}
