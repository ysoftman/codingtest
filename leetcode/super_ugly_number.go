/*
https://leetcode.com/problems/super-ugly-number/
313. Super Ugly Number
Medium

A super ugly number is a positive integer whose prime factors are in the array primes.
Given an integer n and an array of integers primes, return the nth super ugly number.
The nth super ugly number is guaranteed to fit in a 32-bit signed integer.

Example 1:
Input: n = 12, primes = [2,7,13,19]
Output: 32
Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12 super ugly numbers given primes = [2,7,13,19].

Example 2:
Input: n = 1, primes = [2,3,5]
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are in the array primes = [2,3,5].

Constraints:
1 <= n <= 105
1 <= primes.length <= 100
2 <= primes[i] <= 1000
primes[i] is guaranteed to be a prime number.
All the values of primes are unique and sorted in ascending order.
*/

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
n = 12, primes = [2,7,13,19]
temp = [0,0,0,0]

dp[0] = 1
dp[1] = min(2*dp[0], 7*dp[1], 13*dp[2], 19*dp[3]) = 2
dp[1] 이 된 temp 는 증가, update temp = [1,0,0,0]

dp[2] = min(2*dp[0], 7*dp[1], 13*dp[2], 19*dp[3]) = 4
dp[2] 이 된 temp 는 증가, update temp = [2,0,0,0]
...

*/
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1
	temp := make([]int, len(primes))
	for i := 1; i < n; i++ {
		dp[i] = 1<<31 - 1
		for j := 0; j < len(primes); j++ {
			dp[i] = min(dp[i], primes[j]*dp[temp[j]])
		}
		// fmt.Printf("dp[%v]: %v\n", i, dp[i])
		// update temp values
		for j := 0; j < len(temp); j++ {
			if dp[i] == primes[j]*dp[temp[j]] {
				temp[j]++
			}
		}
		// fmt.Println("updated temp:", temp)
	}
	return dp[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
	fmt.Println(nthSuperUglyNumber(1, []int{2, 3, 5}))
}
