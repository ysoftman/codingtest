/*
https://leetcode.com/problems/pascals-triangle-ii/
119. Pascal's Triangle II
Easy
Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:
Input: rowIndex = 3
Output: [1,3,3,1]

Example 2:
Input: rowIndex = 0
Output: [1]

Example 3:
Input: rowIndex = 1
Output: [1,1]

Constraints:
0 <= rowIndex <= 33
*/
package main

import "fmt"

func getRow(rowIndex int) []int {
	result := [][]int{}
	for i := 0; i <= rowIndex; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, temp)
	}
	return result[rowIndex]
}

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
	fmt.Println(getRow(5))
}
