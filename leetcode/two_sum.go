/*
https://leetcode.com/problems/two-sum/
1. Two Sum
Easy
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/
package main

import (
	"fmt"
	"sort"
)

// using hash table
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := range nums {
		temp := target - nums[i]
		if v, exist := m[temp]; exist {
			result := []int{}
			result = append(result, i, v)
			return result
		}
		m[nums[i]] = i
	}
	return nil
}

type n struct {
	val int
	idx int
}

// 배열 정렬후, 시작,끝 인덱스를 두고 찾을때
// sort array and using two (index) pointers : log(n)
func twoSum2(nums []int, target int) []int {
	data := []n{}
	for i := 0; i < len(nums); i++ {
		data = append(data, n{nums[i], i})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].val < data[j].val
	})
	i := 0
	j := len(data) - 1
	for i < j {
		if data[i].val+data[j].val == target {
			return []int{data[i].idx, data[j].idx}
		}
		if data[i].val+data[j].val < target {
			i++
		} else if data[i].val+data[j].val > target {
			j--
		}
	}
	return []int{}
}

func main() {
	fmt.Println("twoSum")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{-3, 4, 3, 90}, 0))
	fmt.Println("twoSum2")
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))
	fmt.Println(twoSum2([]int{3, 3}, 6))
	fmt.Println(twoSum2([]int{-3, 4, 3, 90}, 0))
}
