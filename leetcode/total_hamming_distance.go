/*
https://leetcode.com/problems/total-hamming-distance/
477. Total Hamming Distance
Medium

The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given an integer array nums, return the sum of Hamming distances between all the pairs of the integers in nums.

Example 1:
Input: nums = [4,14,2]
Output: 6
Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
showing the four bits relevant in this case).
The answer will be:
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

Example 2:
Input: nums = [4,14,4]
Output: 4

Constraints:
1 <= nums.length <= 104
0 <= nums[i] <= 109
The answer for the given input will fit in a 32-bit integer.
*/
package main

import "fmt"

func cntBits(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}

// distance(n)+distance(n-1)...distance(n)
// 1 + 2 + ... + (n - 1) = n(n-1)/2
// XOR 연산후 비트 카운트
// time complexity: n(n-1)/2
// time limit exceeded!!!
func totalHammingDistance2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum += cntBits(nums[i] ^ nums[j])
		}
	}
	return sum
}

// time complexity: O(n)
func totalHammingDistance(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sum := 0
	for i := 0; i < 32; i++ {
		cnt := 0
		// 최하위 비트(least significant bit, LSB)는 이진 정수에서 짝수인지 홀수인지를 결정하는 단위값이 되는 비트 위치
		// nums 모든 수의 LSB 의 1 카운트
		for j := 0; j < len(nums); j++ {
			cnt += (nums[j] >> i) & 1
		}
		// (len(nums) - cnt):1아닌 숫자의 개수 * cnt:1인 숫자의 개수 = 현재 자리(i)에서 distance 다.
		// (len(nums) - cnt):1아닌 숫자의 개수가 0 이면 현재 자리(i)에서 모든 숫자가 같아서 distance 는 0 이 된다.
		sum += cnt * (len(nums) - cnt)
	}
	return sum
}

func main() {
	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
	fmt.Println(totalHammingDistance([]int{4, 14, 4}))
}
