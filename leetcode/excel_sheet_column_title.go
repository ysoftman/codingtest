/*
https://leetcode.com/problems/excel-sheet-column-title/
168. Excel Sheet Column Title
Easy
Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

For example:
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

Example 1:
Input: columnNumber = 1
Output: "A"

Example 2:
Input: columnNumber = 28
Output: "AB"

Example 3:
Input: columnNumber = 701
Output: "ZY"

Constraints:
1 <= columnNumber <= 231 - 1
*/
package main

import (
	"fmt"
)

/*
A~Z 26개(0~25) 26진수로 생각
26 28-1
   ---
    1-1      1
   0(A) 1(B)

26 701-1
   ---
    26-1     24
   25(Z) 24(Y)

26 24898
   -----
     957-1   15
      36-1   20
       1-1   9
  0(A) 9(J) 21(U) 16(P)

*/
func convertToTitle(columnNumber int) string {
	ret := ""
	for columnNumber > 0 {
		columnNumber--
		reminder := columnNumber % 26
		ret = string(byte('A'+reminder)) + ret
		columnNumber /= 26
	}

	return ret
}

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(24898))
}
