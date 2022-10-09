/*
https://leetcode.com/problems/sum-of-two-integers/
371. Sum of Two Integers
Medium

Given two integers a and b, return the sum of the two integers without using the operators + and -.

Example 1:
Input: a = 1, b = 2
Output: 3

Example 2:
Input: a = 2, b = 3
Output: 5
Constraints:
-1000 <= a, b <= 1000
*/
package main

import "fmt"

/*
ex) a(2) + b(3) = 5

while b != 0 {

and -> a(10) & b(11) = 10(c)
xor -> a(10) ^ b(11) = 01(a)
shift -> c(10) << 1 = 100(b)

and -> a(01) & b(100) = 000(c)
xor -> a(01) ^ b(100) = 101(a)
shift -> c(000) << 1 = 000(b)

}
return a(101) = 5
*/
func getSum(a int, b int) int {
	c := 0
	for b != 0 {
		c = a & b
		a = a ^ b
		b = c << 1
	}
	return a
}

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(2, 3))
}
