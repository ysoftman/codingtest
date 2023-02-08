/*
https://leetcode.com/problems/maximum-average-subarray-i/
643. Maximum Average Subarray I
Easy

You are given an integer array nums consisting of n elements, and an integer k.
Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

Example 1:
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:
Input: nums = [5], k = 1
Output: 5.00000

Constraints:
n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104
*/
package main

import "fmt"

// time limit exceeded
func findMaxAverage2(nums []int, k int) float64 {
	r := float64(1 << 31 * -1)
	for i := 0; i <= len(nums)-k; i++ {
		temp := 0.0
		for j := i; j < i+k; j++ {
			temp += float64(nums[j])
		}
		temp /= float64(k)
		if r < temp {
			r = temp
		}
	}
	return r
}

/*
ex) [1, 12, -5, -6, 50, 3], k=4
i 번째 값 = 0~i 까지의 값을 누적
[1 13 8 2 52 55]
k 범위의 전체를 더하지 않고, 첫번째 마지막 원소의 차이만 계산 후 k로 나누면 된다.
3(4-1)번째 (2-0) / 4 = 8.5
4번째 (52-1) / 4 = 12.75
5번째 (55-13) / 4 = 10.5 ...  중 max = 12.75
*/
func findMaxAverage(nums []int, k int) float64 {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	// fmt.Println("sums=", sums)

	r := float64(sums[k-1]) / float64(k)
	for i := k; i < len(sums); i++ {
		temp := float64(sums[i]-sums[i-k]) / float64(k)
		// fmt.Println(sums[i], "-", sums[i-k], "/", k, "=", temp)
		if r < temp {
			r = temp
		}
	}
	return r
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
