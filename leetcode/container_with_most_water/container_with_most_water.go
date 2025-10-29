/*
https://leetcode.com/problems/container-with-most-water/
11. Container With Most Water
Medium
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1

*/

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// brute force approach, O(n*n) ---> time limit exceeded!
func maxArea2(height []int) int {
	area, temp := 0, 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			temp = min(height[i], height[j]) * (j - i)
			if area < temp {
				area = temp
			}
		}
	}

	return area
}

// O(n)
func maxArea(height []int) int {
	area, temp := 0, 0
	i, j := 0, len(height)-1
	for i < j {
		temp = min(height[i], height[j]) * (j - i)
		if area < temp {
			area = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return area
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
