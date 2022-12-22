/*
https://leetcode.com/problems/smallest-even-multiple/
2413. Smallest Even Multiple
Easy

Given a positive integer n, return the smallest positive integer that is a multiple of both 2 and n.

Example 1:
Input: n = 5
Output: 10
Explanation: The smallest multiple of both 5 and 2 is 10.

Example 2:
Input: n = 6
Output: 6
Explanation: The smallest multiple of both 6 and 2 is 6. Note that a number is a multiple of itself.

Constraints:
1 <= n <= 150
*/
package main

import "fmt"

// 최소공배수(Least common multiple)
func smallestEvenMultiple2(n int) int {
	if n == 1 {
		return 2
	}
	r := n
	for {
		// 2로 나둬지지 않으면 n 을 계속 더해 간다.
		if r%2 == 0 {
			break
		}
		r += n
	}
	return r
}

/*
2 5
5 (2%5)2
2 (5%2)1
1 (2%2)0 --> 1 최대공약수(gcd)
*/
func gcd(a, b int) int {
	// fmt.Println(a, b)
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// gcd(최대공약수)를 이용한 방법
// lcd(a,b) = a*b / gcd(a,b)
func smallestEvenMultiple(n int) int {
	if n == 1 {
		return 2
	}
	return 2 * n / gcd(2, n)
}

func main() {
	fmt.Println(smallestEvenMultiple(5))
	fmt.Println(smallestEvenMultiple(6))
	fmt.Println(smallestEvenMultiple(1))
	fmt.Println(smallestEvenMultiple(11))
}
