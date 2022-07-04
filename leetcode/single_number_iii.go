/*
https://leetcode.com/problems/single-number-iii/
260. Single Number III
Medium
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order.

You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.

Example 1:
Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.

Example 2:
Input: nums = [-1,0]
Output: [-1,0]

Example 3:
Input: nums = [0,1]
Output: [1,0]

Constraints:
2 <= nums.length <= 3 * 104
-2^31 <= nums[i] <= 2^31 - 1
Each integer in nums will appear twice, only two integers will appear once.
*/
package main

import "fmt"

/*
1  01
2  10
1  01
3  11
2  10
5 101
----- (xor)
  110 => 모두 xor 한 결과
  010 => -110 을 2의 보수(1의보수 +1, 001 > 010)
----- (and)
   10 => and 한 결과

1   1 &  10 =   0 , r[0] ^=   1 >   1(1)
2  10 &  10 =  10 , r[1] ^=  10 >  10(2)
1   1 &  10 =   0 , r[0] ^=   1 >   0(0)
3  11 &  10 =  10 , r[1] ^=  11 >   1(1)
2  10 &  10 =  10 , r[1] ^=  10 >  11(3)
5 101 &  10 =   0 , r[0] ^= 101 > 101(5)
같은 값이 2번 나오면 xor 로 0이 되어 마지막 r[0], r[1] 에 한번씩만 나온 수가 기록된다.
*/

// time complexity: O(n)
// space complexity: O(1)
func singleNumber(nums []int) []int {
	r := []int{0, 0}
	xor := nums[0]
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}
	// fmt.Printf("xor: %b\n", xor)
	// fmt.Printf("-xor: %b\n", -xor)
	xor &= -xor
	// fmt.Printf("xor &= -xor: %b\n", xor)

	for i := 0; i < len(nums); i++ {
		if nums[i]&xor == 0 {
			r[0] ^= nums[i]
			// fmt.Printf("%v %3b & %3b = %3b , r[0] ^= %3b > %3b(%v)\n", nums[i], nums[i], xor, nums[i]&xor, nums[i], r[0], r[0])
		} else {
			r[1] ^= nums[i]
			// fmt.Printf("%v %3b & %3b = %3b , r[1] ^= %3b > %3b(%v)\n", nums[i], nums[i], xor, nums[i]&xor, nums[i], r[1], r[1])
		}
	}
	return r
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{-1, 0}))
	fmt.Println(singleNumber([]int{0, 1}))
}
