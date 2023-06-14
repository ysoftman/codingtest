/*
2390. Removing Stars From a String
Medium
You are given a string s, which contains stars *.
In one operation, you can:
Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:
The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.

Example 1:
Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".

Example 2:
Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters and stars *.
The operation above can be performed on s
*/
package main

import (
	"fmt"
)

// using stack, but time limit exceeded!!!
// func removeStars(s string) string {
// 	result := ""
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '*' {
// 			if len(result) > 0 {
// 				result = result[:len(result)-1]
// 			}
// 		} else {
// 			result += string(s[i])
// 		}
// 	}
// 	return result
// }

func reverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func removeStars(s string) string {
	result := make([]byte, 0)
	right := len(s) - 1
	cntStar := 0
	// 뒤쪽부터 기록한다.
	for right >= 0 {
		if s[right] == '*' {
			cntStar++
		} else if cntStar > 0 {
			cntStar--
		} else {
			result = append(result, s[right])
		}
		right--
	}
	reverseString(result)
	return string(result)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
	fmt.Println(removeStars("abb*cdfg*****x*"))
}
