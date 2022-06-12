/*
https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
1658. Minimum Operations to Reduce X to Zero
Medium
You are given an integer array nums and an integer x. In one operation, you can either remove the leftmost or the rightmost element from the array nums and subtract its value from x. Note that this modifies the array for future operations.

Return the minimum number of operations to reduce x to exactly 0 if it is possible, otherwise, return -1.

Example 1:
Input: nums = [1,1,4,2,3], x = 5
Output: 2
Explanation: The optimal solution is to remove the last two elements to reduce x to zero.

Example 2:
Input: nums = [5,6,7,8,9], x = 4
Output: -1

Example 3:
Input: nums = [3,2,20,1,1,3], x = 10
Output: 5
Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.

Constraints:
1 <= nums.length <= 105
1 <= nums[i] <= 104
1 <= x <= 109
*/
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
two pointer 응용 문제
예) nums = [1,1,4,2,3], x = 5

1+1+4+2+3=11(sum of nums)-5(x) = 6(target)
nums 에서 x 를 제외한 나머지 값들의 합(target)이 되는 지점의 길이를 찾아 len(nums)에 빼야 한다.

taget:6
left=0, right=0
(left=0, right=0)=> sum(nums[left]) ~ nums[right]) = 1
(left=0, right=1)=> sum(nums[left]) ~ nums[right]) = 1+1=2
(left=0, right=2)=> sum(nums[left]) ~ nums[right]) = 1+1+4=6, 6 == target(6)
만약 target 보다 큰 경우는 left 를 증가해서 sum 을 다시 구한다.(이전 sum -= num[left]; left++)

5(nums 길이) - 3(0~2idx=>6, 0~2 의 길이) = 2(나머지 길이)

... 를 right < len(nums) 를 반복하면서 가장 큰 나머지 길이를 아야 최종 len(nums)에서 뺄때 가정 적은(최소회수,minimum-operations)가 된다.
*/

func minOperations(nums []int, x int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	target := sum - x
	// nums 모든 합 == x 랑 딱 맞아 떨어지는 경우
	if target == 0 {
		return len(nums)
	}
	sum = 0
	rest_length := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > target && left < len(nums) {
			sum -= nums[left]
			left++
		}
		if sum == target {
			// rest_length(right-left+1) 중 가장 길어야 최종 len(nums)에서 뺄때 가정 적은(최소회수,minimum-operations)가 된다.
			rest_length = max(rest_length, (right-left)+1)
		}
	}
	if rest_length > 0 {
		return len(nums) - rest_length
	}
	return -1
}

func main() {
	fmt.Println(minOperations([]int{1, 1}, 3))
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
	fmt.Println(minOperations([]int{6016, 5483, 541, 4325, 8149, 3515, 7865, 2209, 9623, 9763, 4052, 6540, 2123, 2074, 765, 7520, 4941, 5290, 5868, 6150, 6006, 6077, 2856, 7826, 9119}, 31841))
}
