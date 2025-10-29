/*
https://leetcode.com/problems/powx-n/
50. Pow(x, n)
Medium
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:
Input: x = 2.00000, n = 10
Output: 1024.00000

Example 2:
Input: x = 2.10000, n = 3
Output: 9.26100

Example 3:
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25

Constraints:
-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104
*/
package main

import "fmt"

/*
2^4 = 1 * 2 * 2 * 2 * 2 = 16
// 음의 지수는 1/밑 을 곱하는것
2^-4 = 1 * 1/2 * 1/2 * 1/2 * 1/2 = 1/16 = (1/2)^4
2^-4 = 1/(2^4) = 1/16
3^-3 = 1/(3^3) = 1/27
-2^-3 = 1/(-2^3) = 1/-8
(5/8)^-2 = 1/((5/8)^2) = 1/(25/64) = 64/25 = (8/5)^2
*/
func myPow(x float64, n int) float64 {
	r := 1.0
	cnt := n
	if n < 0 {
		cnt *= -1
	}
	// time limit exceeded
	// for i:=0; i<cnt; i++ {
	//     r *= x
	// }

	/*
	   거듭제곱의 거듭제곱으로 표현하면 곱하기 회수를 줄일 수 있다.
	   2^10 = (2^2)^5 = (4)^5 = 4^4 * 4^1 = 1024
	   2^20 = (2^2)^10 = ((2^2)^2)^5 = (4^2)^5 = 16^5 = 16^4 * 16^1 = 1048576
	   x^cnt = (x^2)^(cnt/2)...
	*/
	for cnt > 0 {
		// case1) 남은 거듭제곱 회수가 홀수 인 경우, 현재까지 계산된 밑을 곱한다.
		// case2) 마지막 cnt 1이 될때 현재까지 계산된 밑을 곱합고 루프 종료.
		if cnt&1 == 1 { // if (cnt % 2) == 1
			r *= x
		}
		x *= x
		cnt /= 2 // x^2 했기 때문에 cnt 회수를 반으로 나눈다.
	}
	if n < 0 {
		r = 1 / r
	}
	return r
}

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
	fmt.Println(myPow(-2.00000, -3))
	fmt.Println(myPow(5.0000, -245678))
	fmt.Println(myPow(1.00000, -2147483648))
}
