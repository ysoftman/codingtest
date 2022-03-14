/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/
17. Letter Combinations of a Phone Number
Medium
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []

Example 3:
Input: digits = "2"
Output: ["a","b","c"]
*/
package main

import (
	"fmt"
	"strconv"
)

func phoneCombinations(digits string, pre string, i int, m map[string][]string, result *[]string) {
	if i == len(digits) {
		*result = append(*result, pre)
		return
	}
	num := string(digits[i])
	for _, v := range m[num] {
		phoneCombinations(digits, pre+string(v), i+1, m, result)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := make(map[string][]string, 0)
	ch := 'a'
	cnt := 3
	for i := 2; i <= 9; i++ {
		if i == 7 || i == 9 {
			cnt = 4
		} else {
			cnt = 3
		}
		for j := 0; j < cnt; j++ {
			str := strconv.Itoa(i)
			m[str] = append(m[str], string(ch))
			ch++
		}
	}
	// fmt.Println(m)
	result := []string{}
	pre := ""
	phoneCombinations(digits, pre, 0, m, &result)

	return result
}

func main() {
	fmt.Println(letterCombinations("235"))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("5572"))
	fmt.Println(letterCombinations("8888"))
}
