/*
https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
378. Kth Smallest Element in a Sorted Matrix
Medium
Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

You must find a solution with a memory complexity better than O(n2).


Example 1:
Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13

Example 2:
Input: matrix = [[-5]], k = 1
Output: -5

Constraints:
n == matrix.length == matrix[i].length
1 <= n <= 300
-109 <= matrix[i][j] <= 109
All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
1 <= k <= n2

Follow up:
Could you solve the problem with a constant memory (i.e., O(1) memory complexity)?
Could you solve the problem in O(n) time complexity? The solution may be too advanced for an interview but you may find reading this paper fun.
*/
package main

import (
	"fmt"
	"sort"
)

// sumit 성공하지만 memory complexity better than O(n2) 를 만족못함
func kthSmallest2(matrix [][]int, k int) int {
	arr := make([]int, len(matrix)*len(matrix[0]))
	idx := 0
	// space complexity: O(m*n)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			arr[idx] = matrix[i][j]
			idx++
		}
	}
	// time complexity: O(nlogn)
	sort.Ints(arr)
	fmt.Println(arr)
	return arr[k-1]
}

/*
 1  5  9
10 11 13
12 13 15

binary search
1+(15-1) = mid = 15, matrix[0][0] 부터 차례대로 15 보다 작으면 카운트 count++
count < k , left = mid 로 범위를 좁혀 다시 위 과정 반복
count > k , right = mid 로 범위를 좁혀 다시 위 과정 반복
*/
func kthSmallest(matrix [][]int, k int) int {
	left := matrix[0][0]
	right := matrix[len(matrix)-1][len(matrix[0])-1]
	for left < right {
		mid := left + (right-left)/2
		cnt := 0
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] <= mid {
					cnt++
				} else {
					break
				}
			}
		}
		if cnt < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
	fmt.Println(kthSmallest([][]int{{-5, 1}}, 1))
}
