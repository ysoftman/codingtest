/*
https://leetcode.com/problems/excel-sheet-column-number/
171. Excel Sheet Column Number
Easy
Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.
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
Input: columnTitle = "A"
Output: 1

Example 2:
Input: columnTitle = "AB"
Output: 28

Example 3:
Input: columnTitle = "ZY"
Output: 701

Constraints:
1 <= columnTitle.length <= 7
columnTitle consists only of uppercase English letters.
columnTitle is in the range ["A", "FXSHRXW"].
*/
package main

import "fmt"

/*
26진수 -> 10진수로 변경
AB => A(1)*26^1 + B(2)*26^0 = 26+2=28
*/
func titleToNumber(columnTitle string) int {
	step := 1
	result := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += (int(columnTitle[i]-'A') + 1) * step
		step *= 26
	}
	return result
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
	fmt.Println(titleToNumber("FXSHRXW"))
}
