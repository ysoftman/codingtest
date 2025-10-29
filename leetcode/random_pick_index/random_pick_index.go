/*
https://leetcode.com/problems/random-pick-index/
398. Random Pick Index
Medium

Given an integer array nums with possible duplicates, randomly output the index of a given target number. You can assume that the given target number must exist in the array.

Implement the Solution class:

Solution(int[] nums) Initializes the object with the array nums.
int pick(int target) Picks a random index i from nums where nums[i] == target. If there are multiple valid i's, then each index should have an equal probability of returning.


Example 1:
Input
["Solution", "pick", "pick", "pick"]
[[[1, 2, 3, 3, 3]], [3], [1], [3]]
Output
[null, 4, 0, 2]

Explanation
Solution solution = new Solution([1, 2, 3, 3, 3]);
solution.pick(3); // It should return either index 2, 3, or 4 randomly. Each index should have equal probability of returning.
solution.pick(1); // It should return 0. Since in the array only nums[0] is equal to 1.
solution.pick(3); // It should return either index 2, 3, or 4 randomly. Each index should have equal probability of returning.


Constraints:

1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
target is an integer from nums.
At most 104 calls will be made to pick.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().Unix())
	sol := Solution{
		m: make(map[int][]int),
	}
	// key: num
	// value: index
	for i := 0; i < len(nums); i++ {
		sol.m[nums[i]] = append(sol.m[nums[i]], i)
	}
	fmt.Println(sol.m)
	return sol
}

func (this *Solution) Pick(target int) int {
	if indexes, exist := this.m[target]; exist {
		return indexes[rand.Intn(len(indexes))]
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
func main() {
	sol := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(sol.Pick(3))
	fmt.Println(sol.Pick(1))
	fmt.Println(sol.Pick(3))
}
