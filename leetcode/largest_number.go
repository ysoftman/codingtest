/*
https://leetcode.com/problems/largest-number/
179. Largest Number
Medium
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.

Example 1:
Input: nums = [10,2]
Output: "210"

Example 2:
Input: nums = [3,30,34,5,9]
Output: "9534330"

Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 109
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
두개를 더했을때 큰 것을 선태
10, 2 의 경우
210, 102 중 210이 크니 2를 우선

3, 30, 34, 5, 9 경우
330, 303 중 330이 크니 3을 우선
..
59, 95 중 95가 크니 9를 우선
...
*/
func largestNumber(nums []int) string {
	// convert int to string
	strnums := []string{}
	for _, v := range nums {
		n := strconv.Itoa(v)
		strnums = append(strnums, n)
	}
	// fmt.Println(strnums)
	sort.Slice(strnums, func(i, j int) bool {
		a, _ := strconv.Atoi(strnums[i] + strnums[j])
		b, _ := strconv.Atoi(strnums[j] + strnums[i])
		if a > b {
			return true
		}
		return false
	})
	// fmt.Println(strnums)
	r := ""
	for _, v := range strnums {
		r += v
	}
	if r[0] == '0' {
		return "0"
	}
	return r
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{34323, 3432}))
	fmt.Println(largestNumber([]int{0, 0}))
	fmt.Println(largestNumber([]int{0, 0, 0}))
}
