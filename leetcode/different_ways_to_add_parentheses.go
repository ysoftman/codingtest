/*
https://leetcode.com/problems/different-ways-to-add-parentheses/
241. Different Ways to Add Parentheses
Medium
Given a string expression of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. You may return the answer in any order.

Example 1:
Input: expression = "2-1-1"
Output: [0,2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2

Example 2:
Input: expression = "2*3-4*5"
Output: [-34,-14,-10,-10,10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10

Constraints:
1 <= expression.length <= 20
expression consists of digits and the operator '+', '-', and '*'.
All the integer values in the input expression are in the range [0, 99].
*/
package main

import (
	"fmt"
	"strconv"
)

func recursiveCompute(expression string) []int {
	// fmt.Println("expression:", expression)

	r := []int{}
	for i := 0; i < len(expression); i++ {
		// operator 기준으로 a operator b 형식으로 a, b 2부분으로 나눈다.
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			// a , b
			for _, a := range recursiveCompute(expression[:i]) {
				for _, b := range recursiveCompute(expression[i+1:]) {
					if expression[i] == '+' {
						r = append(r, a+b)
					} else if expression[i] == '-' {
						r = append(r, a-b)
					} else if expression[i] == '*' {
						r = append(r, a*b)
					}
				}
			}
		}
	}

	// operator 가 없는 경우
	if len(r) == 0 {
		v, _ := strconv.Atoi(expression)
		return []int{v}
	}

	// fmt.Println("r:", r)
	return r
}

func diffWaysToCompute(expression string) []int {
	return recursiveCompute(expression)
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
	fmt.Println(diffWaysToCompute("11"))
}
