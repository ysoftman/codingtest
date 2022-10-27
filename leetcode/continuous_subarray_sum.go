/*
https://leetcode.com/problems/continuous-subarray-sum/
523. Continuous Subarray Sum
Medium

Given an integer array nums and an integer k, return true if nums has a continuous subarray of size at least two whose elements sum up to a multiple of k, or false otherwise.

An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.

Example 1:
Input: nums = [23,2,4,6,7], k = 6
Output: true
Explanation: [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.

Example 2:
Input: nums = [23,2,6,4,7], k = 6
Output: true
Explanation: [23, 2, 6, 4, 7] is an continuous subarray of size 5 whose elements sum up to 42.
42 is a multiple of 6 because 42 = 7 * 6 and 7 is an integer.

Example 3:
Input: nums = [23,2,6,4,7], k = 13
Output: false

Constraints:
1 <= nums.length <= 105
0 <= nums[i] <= 109
0 <= sum(nums[i]) <= 231 - 1
1 <= k <= 231 - 1
*/
package main

import "fmt"

/*
ex) [23,2,4,6,7], k = 6
23+2 = 25 => 6의 배수 아님
2+4 = 6 => 6의 배수임

i=0, sum=23, remainder=23%6=5, 맵에 나머지에 대한 값이 없어 map[5]=i+1=1 값을 저장
i=1, sum=25, remainder=25%6=1, 맵에 나머지에 대한 값이 없어 map[1]=i+1=2 값을 저장
i=2, sum=29, remainder=29%6=5, 맵에 나머지에 대한 값이 있음 map[5]=1<i(2) 이면 subarray 찾았음(true)
*/
// time complexity: O(n)
// space complexity: O(min(n,k))
func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	// 나머지가 0일 경우가 있어 m[0] = 0 으로 값을 설정해둬야 한다.
	m[0] = 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		remainder := sum % k
		// fmt.Printf("i=%v, sum=%v, remainder=%v%%%v=%v\n", i, sum, sum, k, remainder)
		if _, exist := m[remainder]; !exist {
			m[remainder] = i + 1
			continue
		}
		if m[remainder] < i {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))

	//
	// i=0, sum=23, remainder=23%7=2
	// i=1, sum=25, remainder=25%7=4
	// i=2, sum=29, remainder=29%7=1
	// i=3, sum=35, remainder=35%7=0 나머지가 0일 경우가 있어 m[0] = 0 으로 값을 설정해둬야 한다.
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}
