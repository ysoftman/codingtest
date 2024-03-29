/*
https://leetcode.com/problems/find-the-difference/
389. Find the Difference
Easy
You are given two strings s and t.
String t is generated by random shuffling string s and then add one more letter at a random position.

Return the letter that was added to t.

Example 1:
Input: s = "abcd", t = "abcde"
Output: "e"
Explanation: 'e' is the letter that was added.

Example 2:
Input: s = "", t = "y"
Output: "y"

Constraints:
0 <= s.length <= 1000
t.length == s.length + 1
s and t consist of lowercase English letters.
*/
package main

import "fmt"

// counting
func findTheDifference2(s string, t string) byte {
	a := make(map[byte]int, 26)
	b := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		a[s[i]]++
		b[t[i]]++
	}
	b[t[len(t)-1]]++
	for i := 'a'; i <= 'z'; i++ {
		if a[byte(i)] != b[byte(i)] {
			return byte(i)
		}
	}
	return byte('0')
}

// using xor
func findTheDifference(s string, t string) byte {
	var ch byte
	i := 0
	// 같은 값을 xor 하면 0이 되는 성질을 이용
	// s, t 는 문자 하나만 다르고(홀수개수) 같은 문자가 짝수번 반복되어
	// xor 하면 홀수개만 있는 문자 하나만 남게된다.
	for i < len(s) {
		ch ^= s[i]
		ch ^= t[i]
		i++
	}
	ch ^= t[i]
	return ch
}

func main() {
	fmt.Printf("%c\n", findTheDifference("abcd", "abcde"))
	fmt.Printf("%c\n", findTheDifference("", "y"))
}
