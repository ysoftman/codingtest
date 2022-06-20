/*
https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
309. Greatest English Letter in Upper and Lower Case
Easy
Share
Given a string of English letters s, return the greatest English letter which occurs as both a lowercase and uppercase letter in s. The returned letter should be in uppercase. If no such letter exists, return an empty string.

An English letter b is greater than another letter a if b appears after a in the English alphabet.

Example 1:
Input: s = "lEeTcOdE"
Output: "E"
Explanation:
The letter 'E' is the only letter to appear in both lower and upper case.

Example 2:
Input: s = "arRAzFif"
Output: "R"
Explanation:
The letter 'R' is the greatest letter to appear in both lower and upper case.
Note that 'A' and 'F' also appear in both lower and upper case, but 'R' is greater than 'F' or 'A'.

Example 3:
Input: s = "AbCdEfGhIjK"
Output: ""
Explanation:
There is no letter that appears in both lower and upper case.

Constraints:
1 <= s.length <= 1000
s consists of lowercase and uppercase English letters.
*/
package main

import "fmt"

func greatestLetter(s string) string {
	g := ""
	m := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 'A' && char <= 'Z' {
			m[char] |= 1
		} else if char >= 'a' && char <= 'z' {
			char -= 32 // convert lowercase to uppercase
			m[char] |= 2
		}
		if m[char] == 3 && string(char) > g {
			g = string(char)
		}
	}
	return g
}

func main() {
	fmt.Println(greatestLetter("lEeTcOdE"))
	fmt.Println(greatestLetter("arRAzFif"))
	fmt.Println(greatestLetter("AbCdEfGhIjK"))
}
