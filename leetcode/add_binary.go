/*
https://leetcode.com/problems/add-binary/
67. Add Binary
Easy
Given two binary strings a and b, return their sum as a binary string.

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"

*/
package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		temp := ""
		for i := 0; i < len(b)-len(a); i++ {
			temp += "0"
		}
		a = temp + a
	} else {
		temp := ""
		for i := 0; i < len(a)-len(b); i++ {
			temp += "0"
		}
		b = temp + b
	}
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	result := ""
	addup := 0
	for i := len(a) - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(string(a[i]))
		b, _ := strconv.Atoi(string(b[i]))
		result += strconv.Itoa((a + b + addup) % 2)
		addup = (a + b + addup) / 2
	}
	if addup > 0 {
		result += strconv.Itoa(addup)
	}

	result2 := ""
	for i := len(result) - 1; i >= 0; i-- {
		result2 += string(result[i])
	}

	return result2
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
