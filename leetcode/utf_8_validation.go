/*
https://leetcode.com/problems/utf-8-validation/
393. UTF-8 Validation
Medium
Given an integer array data representing the data, return whether it is a valid UTF-8 encoding (i.e. it translates to a sequence of valid UTF-8 encoded characters).

A character in UTF8 can be from 1 to 4 bytes long, subjected to the following rules:

For a 1-byte character, the first bit is a 0, followed by its Unicode code.
For an n-bytes character, the first n bits are all one's, the n + 1 bit is 0, followed by n - 1 bytes with the most significant 2 bits being 10.
This is how the UTF-8 encoding would work:

	  Number of Bytes   |        UTF-8 Octet Sequence
	                    |              (binary)
	--------------------+-----------------------------------------
	         1          |   0xxxxxxx
	         2          |   110xxxxx 10xxxxxx
	         3          |   1110xxxx 10xxxxxx 10xxxxxx
	         4          |   11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

x denotes a bit in the binary form of a byte that may be either 0 or 1.

Note: The input is an array of integers. Only the least significant 8 bits of each integer is used to store the data. This means each integer represents only 1 byte of data.

Example 1:
Input: data = [197,130,1]
Output: true
Explanation: data represents the octet sequence: 11000101 10000010 00000001.
It is a valid utf-8 encoding for a 2-bytes character followed by a 1-byte character.

Example 2:
Input: data = [235,140,4]
Output: false
Explanation: data represented the octet sequence: 11101011 10001100 00000100.
The first 3 bits are all one's and the 4th bit is 0 means it is a 3-bytes character.
The next byte is a continuation byte which starts with 10 and that's correct.
But the second continuation byte does not start with 10, so it is invalid.

Constraints:
1 <= data.length <= 2 * 104
0 <= data[i] <= 255
*/
package main

import "fmt"

func validUtf8(data []int) bool {
	// 체크할 남은 바이트 수
	cnt := 0
	for _, d := range data {
		// 1~4 중 몇 바이트 utf8 인지 체크
		if cnt == 0 {
			if d>>5 == 0b110 { // 2 바이트 utf8 이라면 남은건 바이트 1개(cnt)
				cnt = 1
			} else if d>>4 == 0b1110 { // 3 바이트 utf8 이라면 남은건 바이트 2개(cnt)
				cnt = 2
			} else if d>>3 == 0b11110 { // 4 바이트 utf8 이라면 남은건 바이트 3개(cnt)
				cnt = 3
			} else if d>>7 != 0b0 { // 1바이트 utf8 경우 체크
				return false
			}
		} else {
			// 남은 바이트들은 10xxxxxx 인지 체크한다.
			if d>>6 == 0b10 {
				cnt--
			} else {
				return false
			}
		}
	}
	return cnt == 0
}

func main() {
	fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(validUtf8([]int{235, 140, 4}))
	fmt.Println(validUtf8([]int{255}))
	fmt.Println(validUtf8([]int{240, 162, 138, 147, 145}))
}
