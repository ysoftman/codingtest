/*
https://leetcode.com/problems/add-strings/
415. Add Strings
Easy

Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

Example 1:
Input: num1 = "11", num2 = "123"
Output: "134"

Example 2:
Input: num1 = "456", num2 = "77"
Output: "533"

Example 3:
Input: num1 = "0", num2 = "0"
Output: "0"

Constraints:
1 <= num1.length, num2.length <= 104
num1 and num2 consist of only digits.
num1 and num2 don't have any leading zeros except for the zero itself.
*/
package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	operand1 := 0
	operand2 := 0
	carry := 0
	sum := []byte{}
	i := len(num1) - 1
	j := len(num2) - 1
	for i >= 0 || j >= 0 || carry > 0 {
		operand1 = 0
		operand2 = 0
		if i >= 0 {
			operand1 = int(num1[i] - '0')
		}
		if j >= 0 {
			operand2 = int(num2[j] - '0')
		}
		temp := operand1 + operand2 + carry
		carry = temp / 10
		mod := temp % 10
		sum = append(sum, byte(mod+'0'))
		i--
		j--
	}
	r := ""
	for i := len(sum) - 1; i >= 0; i-- {
		r += string(sum[i])
	}
	return r
}

func main() {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("0", "0"))
	fmt.Println(addStrings("1", "9"))
	fmt.Println(addStrings("12312312234312412354123421341234123531251", "53425897907431523451843294231412341235"))
}
