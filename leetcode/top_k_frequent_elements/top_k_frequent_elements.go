/*
https://leetcode.com/problems/top-k-frequent-elements/
347. Top K Frequent Elements
Medium
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]

Constraints:
1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.

Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/
package main

import "fmt"

// using bucket sort
// time complexity: O(n)
// space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	// len(nums)+1 개수의 array 생성하고
	// index : 출현빈도, bucket[index] : 출현 숫자들(똑같은 출현빈도의 숫자가 n 개 있을 수 있다.)
	buckets := make([][]int, len(nums)+1)
	for i := range m {
		buckets[m[i]] = append(buckets[m[i]], i)
	}
	// fmt.Println("buckets:", buckets)
	r := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		// 결과에 k (등수)개 까지만 채우기
		if len(r) >= k {
			break
		}
		// 개수가 0 인 원소는 placeholder 로 실제 숫자가 없으니 스킵
		if len(buckets[i]) == 0 {
			continue
		}
		r = append(r, buckets[i]...)
	}
	return r
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}
