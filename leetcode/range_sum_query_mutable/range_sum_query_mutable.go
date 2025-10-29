/*
https://leetcode.com/problems/range-sum-query-mutable
307. Range Sum Query - Mutable
Medium
Given an integer array nums, handle multiple queries of the following types:

Update the value of an element in nums.
Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
void update(int index, int val) Updates the value of nums[index] to be val.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).


Example 1:
Input
["NumArray", "sumRange", "update", "sumRange"]
[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
Output
[null, 9, null, 8]

Explanation
NumArray numArray = new NumArray([1, 3, 5]);
numArray.sumRange(0, 2); // return 1 + 3 + 5 = 9
numArray.update(1, 2);   // nums = [1, 2, 5]
numArray.sumRange(0, 2); // return 1 + 2 + 5 = 8
*/
package main

import "fmt"

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{}
	na.data = make([]int, len(nums))
	copy(na.data, nums)
	return na
}

func (this *NumArray) Update(index int, val int) {
	if index >= 0 && index < len(this.data) {
		this.data[index] = val
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	// fmt.Println("this.data:", this.data)
	if left > right || right >= len(this.data) {
		return 0
	}
	if 0 > left {
		return 0
	}
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.data[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func main() {
	obj := Constructor([]int{1, 3, 5})
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println("update")
	fmt.Println(obj.SumRange(0, 2))
}
