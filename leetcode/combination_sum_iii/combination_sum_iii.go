/*
https://leetcode.com/problems/combination-sum-iii/
216. Combination Sum III
Medium
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:

Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.

Example 1:
Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.

Example 2:
Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.

Example 3:
Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.

Constraints:
2 <= k <= 9
1 <= n <= 60
*/

package main

import "fmt"

func sum(nums []int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}

/*
k=3,n=7
1 -> [2~9] -> [3~9] ... 각 단계마다 원소개수가 3, 합이 7 이 되는지 확인
2 -> [3~9] -> [4~9] ... 각 단계마다 원소개수가 3, 합이 7 이 되는지 확인
...
*/
func recursiveCombinationSum3(k int, n int, curi int, temp []int, r *[][]int) {
	if len(temp) > k {
		return
	}
	if len(temp) == k {
		if sum(temp) == n {
			t := make([]int, k)
			copy(t, temp)
			*r = append(*r, t)
		}
		return
	}
	for i := curi; i < 10; i++ {
		recursiveCombinationSum3(k, n, i+1, append(temp, i), r)
	}
}
func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	recursiveCombinationSum3(k, n, 1, []int{}, &result)
	return result
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
	fmt.Println(combinationSum3(9, 45))
}
