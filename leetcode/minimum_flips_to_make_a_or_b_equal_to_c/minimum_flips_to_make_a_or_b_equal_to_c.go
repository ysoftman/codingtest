/*
https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
1318. Minimum Flips to Make a OR b Equal to c
Medium
Given 3 positives numbers a, b and c. Return the minimum flips required in some bits of a and b to make ( a OR b == c ). (bitwise OR operation).
Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.

Example 1:
Input: a = 2, b = 6, c = 5
Output: 3
Explanation: After flips a = 1 , b = 4 , c = 5 such that (a OR b == c)

Example 2:
Input: a = 4, b = 2, c = 7
Output: 1

Example 3:
Input: a = 1, b = 2, c = 3
Output: 0

Constraints:
1 <= a <= 10^9
1 <= b <= 10^9
1 <= c <= 10^9
*/

package main

import (
	"fmt"
	"strconv"
)

const maxlen = 10

func invert(n byte) byte {
	if n == 1 {
		return 0
	}
	return 1
}
func toByteSlice(binStr string) []byte {
	r := make([]byte, maxlen)
	l := maxlen
	for i := len(binStr) - 1; i >= 0; i-- {
		l--
		r[l] = binStr[i] - '0'
	}
	return r
}
func minFlips(a int, b int, c int) int {
	abin := toByteSlice(strconv.FormatInt(int64(a), 2))
	bbin := toByteSlice(strconv.FormatInt(int64(b), 2))
	cbin := toByteSlice(strconv.FormatInt(int64(c), 2))
	// fmt.Println(abin)
	// fmt.Println(bbin)
	// fmt.Println(cbin)
	cnt := 0
	for i := 0; i < maxlen; i++ {
		if abin[i]|bbin[i] == cbin[i] {
			continue
		}
		if invert(abin[i])|bbin[i] == cbin[i] {
			cnt++
			continue
		}
		if abin[i]|invert(bbin[i]) == cbin[i] {
			cnt++
			continue
		}
		if invert(abin[i])|invert(bbin[i]) == cbin[i] {
			cnt += 2
		}
	}
	return cnt
}

func main() {
	fmt.Println(minFlips(2, 6, 5))
	fmt.Println(minFlips(4, 2, 7))
	fmt.Println(minFlips(1, 2, 3))
	fmt.Println(minFlips(8, 3, 5))
}
