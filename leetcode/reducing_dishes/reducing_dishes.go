/*
https://leetcode.com/problems/reducing-dishes
1402. Reducing Dishes
Hard
A chef has collected data on the satisfaction level of his n dishes. Chef can cook any dish in 1 unit of time.

Like-time coefficient of a dish is defined as the time taken to cook that dish including previous dishes multiplied by its satisfaction level i.e. time[i] * satisfaction[i].

Return the maximum sum of like-time coefficient that the chef can obtain after preparing some amount of dishes.

Dishes can be prepared in any order and the chef can discard some dishes to get this maximum value.

Example 1:
Input: satisfaction = [-1,-8,0,5,-9]
Output: 14
Explanation: After Removing the second and last dish, the maximum total like-time coefficient will be equal to (-1*1 + 0*2 + 5*3 = 14).
Each dish is prepared in one unit of time.

Example 2:
Input: satisfaction = [4,3,2]
Output: 20
Explanation: Dishes can be prepared in any order, (2*1 + 3*2 + 4*3 = 20)

Example 3:
Input: satisfaction = [-1,-4,-5]
Output: 0
Explanation: People do not like the dishes. No dish is prepared.

Constraints:
n == satisfaction.length
1 <= n <= 500
-1000 <= satisfaction[i] <= 1000
*/
package main

import (
	"fmt"
	"sort"
)

// time  complexity O(n^2)
func maxSatisfaction2(satisfaction []int) int {
	sort.Ints(satisfaction)
	max := 0
	for i := range satisfaction {
		sum := 0
		for j := i; j < len(satisfaction); j++ {
			sum += (j - i + 1) * satisfaction[j]
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

// time complexity O(logN)
func maxSatisfaction(satisfaction []int) int {
	// 내림차순
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	max := 0
	prefixSum := 0
	for _, s := range satisfaction {
		prefixSum += s
		// fmt.Println(s, prefixSum, max)
		if prefixSum <= 0 {
			break
		}
		max += prefixSum
	}
	return max
}

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
}
