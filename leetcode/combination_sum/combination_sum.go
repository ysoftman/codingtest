/*
https://leetcode.com/problems/combination-sum/
39. Combination Sum
Medium

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.
It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:
Input: candidates = [2], target = 1
Output: []

Constraints:
1 <= candidates.length <= 30
1 <= candidates[i] <= 200
All elements of candidates are distinct.
1 <= target <= 500
*/
package main

import "fmt"

func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
func recursiveCombinationSum(candidates []int, idx, target int, temp []int, result *[][]int) {
	temp = append(temp, candidates[idx])
	sumTemp := sumSlice(temp)
	// fmt.Println("temp:", temp, sumTemp)
	if sumTemp == target {
		// slice 크기가 변경되지 않으면 메모리가 공유되 때문에 temp의 값이 변경될 수 있어 temp 값을 복사해야 한다.
		copyTemp := make([]int, len(temp))
		copy(copyTemp, temp)
		*result = append(*result, copyTemp)
		return
	}
	if sumTemp > target {
		return
	}

	for i := idx; i < len(candidates); i++ {
		recursiveCombinationSum(candidates, i, target, temp, result)
	}
}
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	temp := []int{}
	for i := 0; i < len(candidates); i++ {
		recursiveCombinationSum(candidates, i, target, temp, &result)
	}
	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{3, 5, 8}, 11))
}
