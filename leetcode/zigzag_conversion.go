/*
// https: //leetcode.com/problems/zigzag-conversion/
6. Zigzag Conversion
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I

Example 3:
Input: s = "A", numRows = 1
Output: "A"

*/
package main

import "fmt"

func convert(s string, numRows int) string {
	temp := make([]string, numRows)
	reverseDir := false
	row := 0
	for _, v := range s {
		temp[row] += string(v)
		if row == numRows-1 {
			reverseDir = true
		}
		if row == 0 {
			reverseDir = false
		}

		if reverseDir {
			row--
		} else {
			row++
		}

		if row > numRows-1 {
			row = numRows - 1
		}
		if row < 0 {
			row = 0
		}
	}
	result := ""
	for _, v := range temp {
		result += v
		// fmt.Println(v)
	}

	return result
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
