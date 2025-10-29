/*
https://leetcode.com/problems/multiply-strings/
43. Multiply Strings
Medium
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

Example 1:
Input: num1 = "2", num2 = "3"
Output: "6"

Example 2:
Input: num1 = "123", num2 = "456"
Output: "56088"

Constraints:
1 <= num1.length, num2.length <= 200
num1 and num2 consist of digits only.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
*/
package main

import "fmt"

func runeToInt(r rune) int {
	return int(r) - '0'
}

func byteToInt(r byte) int {
	return int(r) - '0'
}

func intToString(n int) string {
	return string(n + '0')
}

/*
    123
 x  456
 ------
    738   => 738(10의 자리 carry 1, 100의 자리 carry 1)
   615    => 738+6150(100의 자리 carry 1) => 6888
  492     => 6888+49200(1000의 자리 carry 1) => 56088(1000의 자리 carr 1, 10000의 자리 carry 1)
 ------
  56088   => 결과 자리수(길이)는 3+3 이하

=============

  99
 x99
 ---
 891
891
----
9801

*/
func multiply(num1 string, num2 string) string {
	result := make([]int, len(num1)+len(num2))
	startDigit := 0
	carry := 0
	for i := len(num2) - 1; i >= 0; i-- {
		digit := startDigit
		carry = 0
		for j := len(num1) - 1; j >= 0; j-- {
			temp := byteToInt(num2[i]) * byteToInt(num1[j])
			remainder := temp % 10
			digitRemainder := (result[digit] + remainder + carry) % 10
			digitCarry := (result[digit] + remainder + carry) / 10
			result[digit] = digitRemainder
			carry = digitCarry + temp/10
			digit++
		}
		if carry > 0 {
			result[digit] = carry
		}
		startDigit++
	}
	resultString := ""
	for i := len(result) - 1; i >= 0; i-- {
		if len(resultString) == 0 && intToString(result[i]) == "0" {
			continue
		}
		resultString += intToString(result[i])
	}
	if len(resultString) == 0 {
		resultString = "0"
	}
	return resultString
}
func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("0", "1"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("12413243214132413241324132535098098768576", "876587658659795686585768834985768765875688576"))
}
