/*
https://leetcode.com/problems/detect-capital/
520. Detect Capital
Easy

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital, like "Google".
Given a string word, return true if the usage of capitals in it is right.

Example 1:
Input: word = "USA"
Output: true

Example 2:
Input: word = "FlaG"
Output: false

Constraints:
1 <= word.length <= 100
word consists of lowercase and uppercase English letters.
*/

package main

import "fmt"

func isCapital(b byte) bool {
	if 'A' <= b && b <= 'Z' {
		return true
	}
	return false
}
func detectCapitalUse(word string) bool {
	// a 또는 A 는 옳은 표기
	if len(word) == 1 {
		return true
	}
	// aA 는 잘못된 표기
	if !isCapital(word[0]) && isCapital(word[1]) {
		return false
	}

	allCap := false
	// US 처럼 연속 대문자인 경우
	if isCapital(word[0]) && isCapital(word[1]) {
		allCap = true
	}
	for i := 2; i < len(word); i++ {
		// 연속 대문자라면 이후 계속 대문자여야 한다.
		if allCap && !isCapital(word[i]) {
			return false
		}
		// 연속 대문자가 아니라면 이후 계속 소문자여야 한다.
		if !allCap && isCapital(word[i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("FlaG"))
	fmt.Println(detectCapitalUse("leetcode"))
	fmt.Println(detectCapitalUse("A"))
	fmt.Println(detectCapitalUse("a"))
}
