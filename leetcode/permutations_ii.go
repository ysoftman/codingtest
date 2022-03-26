/*
https://leetcode.com/problems/permutations-ii/
47. Permutations II
Medium
Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

Example 1:
Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]

Example 2:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Constraints:
1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
package main

import "fmt"

/*

                                     1,1,2
(0번째교환)           1,1,2          1,1,2             2,1,1
(0고정, 1교환)    1,1,2  1,2,1      1,1,2  1,2,1   2,1,1  2,1,1


                                   1,2,3
(0번째교환)           1,2,3           2,1,3             3,1,2
(0고정, 1교환)    1,2,3  1,3,2      2,1,3  2,3,1      3,1,2  3,2,1


                                   1,2,2
(0번째교환)           1,2,2           2,1,2             2,2,1
(0고정, 1교환)    1,2,2  1,2,2      2,1,2  2,2,1      2,2,1  2,1,2
*/

func getkey(nums []int) string {
	k := ""
	for _, v := range nums {
		k += string(v + '0')
	}
	// fmt.Println(k)
	return k
}
func recursivePermuteUnique(nums []int, a, b int, duplicated map[string]bool, result *[][]int) {
	if a >= b {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	for i := a; i <= b; i++ {
		nums[a], nums[i] = nums[i], nums[a]
		key := getkey(nums)
		if _, exist := duplicated[key]; !exist {
			recursivePermuteUnique(nums, a+1, b, duplicated, result)
			duplicated[key] = true
		}

		// 현재 루프에서는 원위치시켜 a 자리에 다음 i 가 위치할 수 있도록 한다.
		nums[a], nums[i] = nums[i], nums[a]
	}
}
func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	duplicated := map[string]bool{}
	recursivePermuteUnique(nums, 0, len(nums)-1, duplicated, &result)
	return result
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 2, 2}))
	fmt.Println(permuteUnique([]int{1, 3, 3, 2, 1, 5}))
}
