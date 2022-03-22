/*
https://leetcode.com/problems/combination-sum-ii/
40. Combination Sum II
Medium
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.
Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.

Example 1:
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/
package main

import (
	"fmt"
	"sort"
)

func sum(n []int, idx int) int {
	if idx >= len(n) {
		return 0
	}
	return n[idx] + sum(n, idx+1)
}
func recursiveCombinationSum2(candi, temp []int, idx, target int, result *[][]int) {
	temp = append(temp, candi[idx])
	// fmt.Println(temp)
	s := sum(temp, 0)
	if s == target {
		r := make([]int, len(temp))
		copy(r, temp)
		*result = append(*result, r)
		return
	}
	if s > target {
		return
	}
	pre := -1
	for i := idx + 1; i < len(candi); i++ {
		if pre == candi[i] {
			continue
		}
		recursiveCombinationSum2(candi, temp, i, target, result)
		pre = candi[i]
	}

}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	temp := []int{}
	pre := -1
	for i := 0; i < len(candidates); i++ {
		if pre == candidates[i] {
			continue
		}
		recursiveCombinationSum2(candidates, temp, i, target, &result)
		pre = candidates[i]
	}
	return result
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 7))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	fmt.Println(combinationSum2([]int{4, 2, 5, 2, 5, 3, 1, 5, 2, 2}, 9))
}
