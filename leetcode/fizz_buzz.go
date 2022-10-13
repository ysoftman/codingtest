/*
https://leetcode.com/problems/fizz-buzz/
412. Fizz Buzz
Easy

Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.

Example 1:
Input: n = 3
Output: ["1","2","Fizz"]

Example 2:
Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]

Example 3:
Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]

Constraints:
1 <= n <= 104
*/
package main

import "fmt"

func fizzBuzz2(n int) []string {
	r := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			r = append(r, "FizzBuzz")
		} else if i%3 == 0 {
			r = append(r, "Fizz")
		} else if i%5 == 0 {
			r = append(r, "Buzz")
		} else {
			r = append(r, fmt.Sprintf("%d", i))
		}
	}
	return r
}

func fizzBuzz(n int) []string {
	r := []string{}
	for i := 1; i <= n; i++ {
		temp := ""
		if i%3 == 0 {
			temp = "Fizz"
		}
		if i%5 == 0 {
			temp += "Buzz"
		}
		if temp == "" {
			temp = fmt.Sprintf("%d", i)
		}
		r = append(r, temp)
	}
	return r
}

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}
