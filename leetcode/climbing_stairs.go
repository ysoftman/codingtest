/*
https://leetcode.com/problems/climbing-stairs/

70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps


Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:

1 <= n <= 45
*/
package main

import "fmt"

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	temp := []int{}

	// n == 1, 1
	temp = append(temp, 1)
	// n == 2, 1+1, 2
	temp = append(temp, 2)
	// n == 3, 1+1+1, 2+1, 1+2
	// n == 4, 1+1+1+1, 2+2, 2+1+1, 1+2+1, 1+1+2
	// n == 5, 1+1+1+1+1, 2+2+1, 2+1+2, 1+2+2, 2+1+1+1, 1+2+1+1, 1+1+2+1, 1+1+1+2
	// .. n = pre + ppre
	for i := 2; i < n; i++ {
		temp = append(temp, temp[i-1]+temp[i-2])
	}
	fmt.Println("temp:", temp)

	return temp[len(temp)-1]
}

func main() {
	for i := 2; i < 5; i++ {
		fmt.Printf("%v -> %v\n", i, climbStairs(i))
	}
}
