/*
https://leetcode.com/problems/count-binary-substrings
696. Count Binary Substrings
Easy

Given a binary string s, return the number of non-empty substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

Substrings that occur multiple times are counted the number of times they occur.

Example 1:

Input: s = "00110011"
Output: 6
Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
Notice that some of these substrings repeat and are counted the number of times they occur.
Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
Example 2:

Input: s = "10101"
Output: 4
Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.

Constraints:
1 <= s.length <= 105
s[i] is either '0' or '1'.
*/

/*
011 -> 1개
0011 -> 2개
00011 -> 2개
001100 -> 0011 부분에서 2개 1100 부분에서 2개 = 4개
로 0과 1일 연속 부분에서 0과 1의 개수중 작은쪽이 0 과 1 개수가 같은 서브 스트링이 개수가 된다.
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countBinarySubstrings(s string) int {
	zeroCnt := 0
	lastZeroIdx := -1
	oneCnt := 0
	lastOneIdx := -1
	result := 0
	for i, v := range s {
		if v == '0' {
			if lastZeroIdx == i-1 {
				zeroCnt++
			} else {
				if zeroCnt > 0 && oneCnt > 0 {
					result += min(zeroCnt, oneCnt)
				}
				zeroCnt = 1
			}
			lastZeroIdx = i
		} else {
			if lastOneIdx == i-1 {
				oneCnt++
			} else {
				if zeroCnt > 0 && oneCnt > 0 {
					result += min(zeroCnt, oneCnt)
				}
				oneCnt = 1
			}
			lastOneIdx = i
		}
	}
	if zeroCnt > 0 && oneCnt > 0 {
		result += min(zeroCnt, oneCnt)
	}

	return result
}

func main() {
	fmt.Println(countBinarySubstrings("0"))
	fmt.Println(countBinarySubstrings("1"))
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
	fmt.Println(countBinarySubstrings("00100"))
}
