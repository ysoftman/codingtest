/*
https://leetcode.com/problems/kth-missing-positive-number/
1539. Kth Missing Positive Number
Easy
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
Find the kth positive integer that is missing from this array.

Example 1:
Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

Example 2:
Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.

Constraints:
1 <= arr.length <= 1000
1 <= arr[i] <= 1000
1 <= k <= 1000
arr[i] < arr[j] for 1 <= i < j <= arr.length
*/

package main

import "fmt"

// brute-force
func findKthPositive2(arr []int, k int) int {
	n := 1
	missingArr := []int{}
	for i := 0; i < len(arr); i++ {
		for arr[i] != n {
			missingArr = append(missingArr, n)
			n++
		}
		n++
	}
	for i := 0; i < k; i++ {
		missingArr = append(missingArr, n)
		n++
	}
	fmt.Println(missingArr)
	return missingArr[k-1]
}

// binary search
func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (right-left)/2 + left
		// 원래 빠진게 없다면 mid(빠진 숫자 없을때 값)와 실제 arr[mid] 는 같아야 한다.
		// 빠진게 있다면 arr[mid] > mid 상태
		// mid 는 0 기준이고, arr 값은 1부터 시작이라 -1 해준다.
		// ex) [2,3,4,7,9], k=5
		// mid = (5-0)/2+0 = 2, arr[2](4)-1-2 < 5,  4(mid=2) => left=mid+1, left=2+1
		// mid = (5-3)/2+3 = 4, arr[4](11)-1-4 >=5, 11(mid=4) => right=mid, right=4
		// mid = (4-3)/2+3 = 3, arr[3](7)-1-3 < 5,  7(mid=3) => left=mid+1, left=4
		if (arr[mid]-1)-mid < k {
			left = mid + 1
		} else {
			right = mid
		}
		fmt.Printf("mid:%v left:%v right:%v\n", mid, left, right)
	}
	// fmt.Println("left:",left)
	return left + k
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 9}, 5))
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}
