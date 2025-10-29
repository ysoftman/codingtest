/*
https://leetcode.com/problems/license-key-formatting/
482. License Key Formatting
Easy

You are given a license key represented as a string s that consists of only alphanumeric characters and dashes. The string is separated into n + 1 groups by n dashes. You are also given an integer k.

We want to reformat the string s such that each group contains exactly k characters, except for the first group, which could be shorter than k but still must contain at least one character. Furthermore, there must be a dash inserted between two groups, and you should convert all lowercase letters to uppercase.

Return the reformatted license key.

Example 1:
Input: s = "5F3Z-2e-9-w", k = 4
Output: "5F3Z-2E9W"
Explanation: The string s has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.

Example 2:
Input: s = "2-5g-3-J", k = 2
Output: "2-5G-3J"
Explanation: The string s has been split into three parts, each part has 2 characters except the first part as it could be shorter as mentioned above.

Constraints:
1 <= s.length <= 105
s consists of English letters, digits, and dashes '-'.
1 <= k <= 104
*/
package main

import "fmt"

func licenseKeyFormatting(s string, k int) string {
	r := ""
	buffer := ""
	// remove -
	// convert all lowercase letters to uppercase.
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			continue
		}
		temp := s[i]
		if 'a' <= s[i] && s[i] <= 'z' {
			temp = s[i] - ('a' - 'A')
		}
		buffer += string(temp)
	}
	// fmt.Println("buffer:", buffer)
	if len(buffer) <= k {
		return buffer
	}

	// 뒤부터 k개를 그룹으로 하고 남은 개수를 처음 그룹으로 설정한다.
	firstIdx := len(buffer) % k
	pre := 0
	for i := firstIdx; i < len(buffer); i += k {
		r += buffer[pre:i]
		if len(r) > 0 {
			r += "-"
		}
		pre = i
	}
	if len(buffer)-k > 0 {
		r += buffer[len(buffer)-k:]
	}
	return r
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
	fmt.Println(licenseKeyFormatting("2-4A0r7-4k", 4))
	fmt.Println(licenseKeyFormatting("2", 2))
	fmt.Println(licenseKeyFormatting("12", 2))
	fmt.Println(licenseKeyFormatting("123", 2))
	fmt.Println(licenseKeyFormatting("1234", 2))
	fmt.Println(licenseKeyFormatting("12345", 2))
}
