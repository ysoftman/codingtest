/*
https://leetcode.com/problems/determine-if-string-halves-are-alike/
1704. Determine if String Halves Are Alike
Easy

You are given a string s of even length. Split this string into two halves of equal lengths, and let a be the first half and b be the second half.

Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). Notice that s contains uppercase and lowercase letters.

Return true if a and b are alike. Otherwise, return false.

Example 1:
Input: s = "book"
Output: true
Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.

Example 2:
Input: s = "textbook"
Output: false
Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
Notice that the vowel o is counted twice.

Constraints:
2 <= s.length <= 1000
s.length is even.
s consists of uppercase and lowercase letters.
*/
package main

import "fmt"

func isVowel(b byte) bool {
	for _, v := range []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
		if b == v {
			return true
		}
	}
	return false
}
func halvesAreAlike(s string) bool {
	leftVowels := 0
	rightVowels := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)/2 && isVowel(s[i]) {
			leftVowels++
		}
		if i >= len(s)/2 && isVowel(s[i]) {
			rightVowels++
		}
	}
	return leftVowels == rightVowels
}
func main() {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
}
