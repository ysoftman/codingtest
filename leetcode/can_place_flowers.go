/*
https://leetcode.com/problems/can-place-flowers/
605. Can Place Flowers
Easy

You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.

Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true

Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: false

Constraints:
1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length
*/
package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		plant := false
		if flowerbed[i] == 1 {
			continue
		}
		if i == 0 {
			if i+1 < len(flowerbed) && flowerbed[i+1] == 0 {
				plant = true
			} else if i == len(flowerbed)-1 {
				plant = true
			}
		} else if i-1 >= 0 && flowerbed[i-1] == 0 {
			if (i+1 < len(flowerbed) && flowerbed[i+1] == 0) || i == len(flowerbed)-1 {
				plant = true
			}
		}
		if plant {
			flowerbed[i] = 1
			n--
		}
		if n == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0}, 1))
	fmt.Println(canPlaceFlowers([]int{1}, 0))
	fmt.Println(canPlaceFlowers([]int{1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}
