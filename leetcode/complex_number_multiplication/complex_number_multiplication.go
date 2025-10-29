/*
https://leetcode.com/problems/complex-number-multiplication/
537. Complex Number Multiplication
Medium

A complex number can be represented as a string on the form "real+imaginaryi" where:
real is the real part and is an integer in the range [-100, 100].
imaginary is the imaginary part and is an integer in the range [-100, 100].
i2 == -1.
Given two complex numbers num1 and num2 as strings, return a string of the complex number that represents their multiplications.

Example 1:
Input: num1 = "1+1i", num2 = "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.

Example 2:
Input: num1 = "1+-1i", num2 = "1+-1i"
Output: "0+-2i"
Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.

Constraints:
num1 and num2 are valid complex numbers.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// complexNumber(복수수) = a+bi 꼴로 표현되는 수(a,b실수, i허수)
// complexNumber multiplications(복소수 곱셈 표현)
// (a+bi)*(x+yi)
// = ax+ayi+bxi+byi2(제곱)
// = ax+byi2(제곱)+(ay+bx)i, i2(제곱)=>-1
// = ax−by+(ay+bx)i
func complexNumberMultiply(num1 string, num2 string) string {
	n1 := strings.Split(num1, "+")
	n1[1] = n1[1][:len(n1[1])-1]
	n2 := strings.Split(num2, "+")
	n2[1] = n2[1][:len(n2[1])-1]
	// fmt.Println(n1)
	// fmt.Println(n2)
	n11, _ := strconv.Atoi(n1[0])
	n12, _ := strconv.Atoi(n1[1])
	n21, _ := strconv.Atoi(n2[0])
	n22, _ := strconv.Atoi(n2[1])
	// ax−by+(ay+bx)i
	return fmt.Sprintf("%d+%di", (n11*n21 - n12*n22), (n11*n22 + n12*n21))
}

func main() {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}
