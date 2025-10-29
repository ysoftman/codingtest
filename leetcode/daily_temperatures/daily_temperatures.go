/*
https://leetcode.com/problems/daily-temperatures/
739. Daily Temperatures
Medium

Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/
package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := []int{len(temperatures) - 1}
	// 맨끝에는 다음 날이 없음으로 0
	r := []int{0}
	// 맨끝 -2 날 부터 처음순으로 파악
	for i := len(temperatures) - 2; i >= 0; i-- {
		// 스택에 쌓인 날의 온도 중 현재보다 큰 날을 찾는다.
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if temperatures[top] <= temperatures[i] {
				// pop
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		// fmt.Println("stack:", stack)
		// 뒤에 온도가 큰 날이 없는 경우 0 으로 표시
		if len(stack) == 0 {
			r = append(r, 0)
		} else {
			top := stack[len(stack)-1]
			// fmt.Println("top:", top, "i:", i)
			r = append(r, top-i)
		}
		stack = append(stack, i)
	}
	// 맨 긑 날짜부터 계산한거라 reverse 한다.
	r2 := []int{}
	for i := len(r) - 1; i >= 0; i-- {
		r2 = append(r2, r[i])
	}
	return r2

}
func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}
