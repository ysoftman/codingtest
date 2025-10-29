/*
https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
1523. Count Odd Numbers in an Interval Range
Easy
Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).

Example 1:
Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].

Example 2:
Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].

Constraints:
0 <= low <= high <= 10^9
*/
package main

import "fmt"

// brute-force, time complexity: O(1+high-low)
func countOdds2(low int, high int) int {
	result := 0
	for i := low; i <= high; i++ {
		if i%2 == 1 {
			result++
		}
	}
	return result
}

// time complexity: O(1)
/*
0,0 0
1,1 1

0,1 1
0,2 1
0,3 2
0,4 2
0,5 3

1,2 1
1,3 2
1,4 2
1,5 3
*/
func countOdds(low int, high int) int {
	add := 0
	if low%2 == 1 {
		add++
	}
	if low == high {
		return add
	}
	if high%2 == 1 {
		add++
		high--
	}
	return (high-low)/2 + add
}

func main() {
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
	fmt.Println(countOdds(0, 0))
	fmt.Println(countOdds(1, 1))
	fmt.Println(countOdds(0, 1))
	fmt.Println(countOdds(0, 2))
	fmt.Println(countOdds(0, 3))
	fmt.Println(countOdds(0, 4))
	fmt.Println(countOdds(0, 5))
}
