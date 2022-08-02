/*
https://leetcode.com/problems/integer-to-roman/
12. Integer to Roman
Medium
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.

Example 1:
Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.


Example 2:
Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.

Example 3:
Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:

1 <= num <= 3999
*/
package main

import (
	"fmt"
	"strings"
)

func convertToRoman(num, multiply int) string {
	result := ""

	one := "I"
	five := "V"
	ten := "X"
	if multiply == 1000 {
		one = "M"
	} else if multiply == 100 {
		one = "C"
		five = "D"
		ten = "M"
	} else if multiply == 10 {
		one = "X"
		five = "L"
		ten = "C"
	}

	switch num {
	case 1, 2, 3:
		for i := 0; i < num; i++ {
			result += one
		}
	case 4:
		result += one + five
	case 5:
		result += five
	case 6, 7, 8:
		result += five
		for i := 0; i < num-5; i++ {
			result += one
		}
	case 9:
		result += one + ten
	}

	return result
}

func intToRoman(num int) string {
	result := []string{}
	multiply := 1
	for num > 0 {
		remainder := num % 10
		result = append(result, convertToRoman(remainder, multiply))
		num /= 10
		multiply *= 10

	}

	temp := make([]string, len(result))
	j := 0
	for i := len(result) - 1; i >= 0; i-- {
		temp[j] = result[i]
		j++
	}
	return strings.Join(temp[:], "")
}

func main() {
	fmt.Println("3 -> ", intToRoman(3))
	fmt.Println("58 -> ", intToRoman(58))
	fmt.Println("1994 -> ", intToRoman(1994))

}
