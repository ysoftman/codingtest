/*
https://leetcode.com/problems/count-primes/
204. Count Primes
Medium
Given an integer n, return the number of prime numbers that are strictly less than n.

Example 1:
Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

Example 2:
Input: n = 0
Output: 0

Example 3:
Input: n = 1
Output: 0

Constraints:
0 <= n <= 5 * 106
*/
package main

import "fmt"

// title : Sieve of Eratosthenes(에라토스테네스의 체) 방법으로 소수 구하기
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
// https://ko.wikipedia.org/wiki/%EC%97%90%EB%9D%BC%ED%86%A0%EC%8A%A4%ED%85%8C%EB%84%A4%EC%8A%A4%EC%9D%98_%EC%B2%B4
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	// 소수가 될 수 있는 후보 숫자들
	candidate_number := make([]bool, n+1)

	// 처음엔 모두 소수라고 취급
	for i := 2; i <= n; i++ {
		candidate_number[i] = true
	}

	// i*i 까지만 검사하면 됨
	for i := 2; i*i < n; i++ {
		if candidate_number[i] == true {
			// i 의 배수만큼에 해당하는 숫자는 소수가 아님
			for j := i * i; j < n; j += i {
				candidate_number[j] = false
			}
		}
	}
	prime_number := []int{}
	// 배수로 체크되지 않는 숫자가 결국 소수가 됨
	for i := 0; i < n; i++ {
		if candidate_number[i] == true {
			prime_number = append(prime_number, i)
		}
	}
	// fmt.Println(prime_number)
	return len(prime_number)
}

func main() {
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(4))
	fmt.Println(countPrimes(5))
	fmt.Println(countPrimes(6))
	fmt.Println(countPrimes(7))
	fmt.Println(countPrimes(8))
	fmt.Println(countPrimes(9))
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(10000000))
}
