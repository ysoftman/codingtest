/*
https://leetcode.com/problems/check-if-the-sentence-is-pangram/
1832. Check if the Sentence Is Pangram
Easy

A pangram is a sentence where every letter of the English alphabet appears at least once.

Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

Example 1:
Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Explanation: sentence contains at least one of every letter of the English alphabet.

Example 2:
Input: sentence = "leetcode"
Output: false

Constraints:
1 <= sentence.length <= 1000
sentence consists of lowercase English letters.
*/
package main

import "fmt"

// using map
func checkIfPangram2(sentence string) bool {
	m := make(map[rune]bool)
	for _, v := range sentence {
		m[v] = true
	}
	for i := 'a'; i <= 'z'; i++ {
		if m[i] == false {
			return false
		}
	}
	return true
}

// using array
func checkIfPangram(sentence string) bool {
	arr := make([]bool, 26)
	for i := 0; i < len(sentence); i++ {
		arr[sentence[i]-'a'] = true
	}
	for i := 0; i < 26; i++ {
		if arr[i] == false {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
	fmt.Println(checkIfPangram("fkdajfiqowyerckhgsdyfioqfnkhuday"))
}
