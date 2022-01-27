/*
https://leetcode.com/problems/valid-palindrome/
125. Valid Palindrome
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/
package main

import "fmt"

func isPalindrome(s string) bool {
	temp := ""
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= '0' && v <= '9' {
			temp += string(v)
			continue
		}
		if v < 'a' {
			v = s[i] + ('a' - 'A')
		}
		if v >= 'a' && v <= 'z' {
			temp += string(v)
		}
	}

	fmt.Println(">", temp)
	for i := 0; i < len(temp)/2; i++ {
		if temp[i] != temp[len(temp)-1-i] {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("0P"))
}
