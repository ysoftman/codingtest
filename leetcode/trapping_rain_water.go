/*
https://leetcode.com/problems/trapping-rain-water/
42. Trapping Rain Water
Hard
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Example 1:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

Example 2:
Input: height = [4,2,0,3,2,5]
Output: 9

Constraints:
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// time complexity : O(N)
func trap(height []int) int {
	trappedRain := make([]int, len(height))

	max := 0
	// dynamic programming
	// 오른쪽은 bar 는 현재 max 만큼 있다고 생각하고 계산
	for i := 0; i < len(height); i++ {
		if max < height[i] {
			max = height[i]
		}
		trappedRain[i] = max - height[i]
	}
	max = 0
	// 왼쪽 bar 는 현재 max 만큼 있다고 생각하고
	// 먼저 계산된 영역과 겹치는(min)것이 실제 빗물이다.
	for i := len(height) - 1; i >= 0; i-- {
		if max < height[i] {
			max = height[i]
		}
		trappedRain[i] = min(trappedRain[i], max-height[i])
	}

	// sum of trappedRain
	sum := 0
	for i := 0; i < len(height); i++ {
		sum += trappedRain[i]
	}
	return sum
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
