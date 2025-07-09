/*
https://leetcode.com/problems/numbers-with-same-consecutive-differences/
967. Numbers With Same Consecutive Differences
Medium
Return all non-negative integers of length n such that the absolute difference between every two consecutive digits is k.

Note that every number in the answer must not have leading zeros. For example, 01 has one leading zero and is invalid.

You may return the answer in any order.

Example 1:
Input: n = 3, k = 7
Output: [181,292,707,818,929]
Explanation: Note that 070 is not a valid number, because it has leading zeroes.

Example 2:
Input: n = 2, k = 1
Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]

Constraints:
2 <= n <= 9
0 <= k <= 9
*/

package main

import (
	"fmt"
	"strconv"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func numSliceToString(nums []int) string {
	str := ""
	for i := 0; i < len(nums); i++ {
		str += fmt.Sprint(nums[i])
	}
	// fmt.Println(str)
	return str
}

/*
dfs
가자리(tree 에서의 level) 가 0~9로 했을때 이전 자리의 수와 차이가 k 인 경로만 결과 포함

		    1            2         ...           9
		.. 8 ..       ..  9                ..  2 ..
	 .. 1 ..           .. 2 ..              .. 9
*/
func recursiveNumsSameConsecDiff(n int, k int, curidx int, cur []int, r *[]int) {
	if len(cur) > n {
		return
	}
	// 0으로 시작하는 경우 결과에서 제외
	if len(cur) == n && cur[0] != 0 {
		curn, _ := strconv.Atoi(numSliceToString(cur))
		// fmt.Println("-curn-:", curn)
		*r = append(*r, curn)
		return
	}
	for i := 0; i <= 9; i++ {
		if abs(cur[curidx]-i) == k {
			recursiveNumsSameConsecDiff(n, k, curidx+1, append(cur, i), r)
		}
	}
}

func numsSameConsecDiff(n int, k int) []int {
	r := []int{}
	for j := 1; j <= 9; j++ {
		recursiveNumsSameConsecDiff(n, k, 0, []int{j}, &r)
	}
	return r
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(5, 6))
	fmt.Println(numsSameConsecDiff(5, 6))
	fmt.Println(numsSameConsecDiff(7, 3))
}
