/*
https://leetcode.com/problems/fraction-addition-and-subtraction/
592. Fraction Addition and Subtraction
Medium
Given a string expression representing an expression of fraction addition and subtraction, return the calculation result in string format.
The final result should be an irreducible fraction.
If your final result is an integer, change it to the format of a fraction that has a denominator 1.
So in this case, 2 should be converted to 2/1.

Example 1:
Input: expression = "-1/2+1/2"
Output: "0/1"

Example 2:
Input: expression = "-1/2+1/2+1/3"
Output: "1/3"

Example 3:
Input: expression = "1/3-1/2"
Output: "-1/6"

Constraints:
The input string only contains '0' to '9', '/', '+' and '-'. So does the output.
Each fraction (input and output) has the format ±numerator/denominator. If the first input fraction or the output is positive, then '+' will be omitted.
The input only contains valid irreducible fractions, where the numerator and denominator of each fraction will always be in the range [1, 10].
If the denominator is 1, it means this fraction is actually an integer in a fraction format defined above.
The number of given fractions will be in the range [1, 10].
The numerator and denominator of the final result are guaranteed to be valid and in the range of 32-bit int.
*/

package main

import (
	"fmt"
	"strconv"
)

// 최소공배수(Least common multiple)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 최대공약수(gcd)
func gcd(a, b int) int {
	// fmt.Println(a, b)
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func fractionAddition(expression string) string {
	// 분자
	numerator := []int{}
	// 분모
	denominator := []int{}
	pre := 0
	// 분자, 분모 값들을 구분해 저장
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '/':
			n, _ := strconv.Atoi(expression[pre:i])
			numerator = append(numerator, n)
			pre = i + 1
		case '+', '-':
			if i == 0 {
				continue
			}
			n, _ := strconv.Atoi(expression[pre:i])
			denominator = append(denominator, n)
			pre = i
		}
	}
	n, _ := strconv.Atoi(expression[pre:])
	denominator = append(denominator, n)

	// 분모들 최소공배수로 만들기(통분)
	commonDenominator := denominator[0]
	for i := 0; i < len(denominator); i++ {
		commonDenominator = lcm(commonDenominator, denominator[i])
	}
	// fmt.Println("commonDenominator:", commonDenominator)

	// 분자,분모를 gcd 로 약분
	sumNumerator := 0
	for i := 0; i < len(numerator); i++ {
		sumNumerator += (commonDenominator / denominator[i]) * numerator[i]
	}

	// fmt.Println(numerator)
	// fmt.Println(denominator)
	if sumNumerator == 0 {
		return "0/1"
	}

	sumNumeratorTemp := sumNumerator
	if sumNumeratorTemp < 0 {
		sumNumeratorTemp *= -1
	}
	commonDenominatorTemp := commonDenominator
	if commonDenominatorTemp < 0 {
	}

	// 약분
	gcd := gcd(sumNumeratorTemp, commonDenominatorTemp)
	sumNumerator /= gcd
	commonDenominator /= gcd
	r := fmt.Sprintf("%d/%d", sumNumerator, commonDenominator)
	return r
}

func main() {
	fmt.Println(fractionAddition("-1/2+1/2"))
	fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("1/3-1/2"))
}
