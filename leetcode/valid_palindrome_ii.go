/*
https://leetcode.com/problems/valid-palindrome-ii/
680. Valid Palindrome II
Given a string s, return true if the s can be palindrome after deleting at most one character from it.

Example 1:
Input: s = "aba"
Output: true

Example 2:
Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.

Example 3:
Input: s = "abc"
Output: false

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters.
*/
package main

import "fmt"

func validPalindrome(s string) bool {
	for delChar := 'a'; delChar <= 'z'; delChar++ {
		i := 0
		j := len(s) - 1
		ok := true
		delCnt := 0
		for i < j {
			if s[i] != s[j] {
				if delCnt == 0 {
					if (byte)(delChar) == s[i] {
						i++
						delCnt++
						continue
					}
					if (byte)(delChar) == s[j] {
						j--
						delCnt++
						continue
					}
				}
				ok = false
				break
			}
			i++
			j--
		}
		if ok {
			// fmt.Println("---", (string)(delChar))
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
	fmt.Println(validPalindrome("abcca"))
	fmt.Println(validPalindrome("eccer"))
	fmt.Println(validPalindrome("pidbliassaqozokmtgahluruufwbjdtayuhbxwoicviygilgzduudzgligyviciowxbhuyatdjbwfuurulhagtmkozoqassailbdip"))
	fmt.Println(validPalindrome("aabbaa"))
}
