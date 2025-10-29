/*
https://leetcode.com/problems/minimum-index-sum-of-two-lists/
599. Minimum Index Sum of Two Lists
Easy
Given two arrays of strings list1 and list2, find the common strings with the least index sum.

A common string is a string that appeared in both list1 and list2.

A common string with the least index sum is a common string such that if it appeared at list1[i] and list2[j] then i + j should be the minimum value among all the other common strings.

Return all the common strings with the least index sum. Return the answer in any order.

Example 1:
Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
Output: ["Shogun"]
Explanation: The only common string is "Shogun".

Example 2:
Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["KFC","Shogun","Burger King"]
Output: ["Shogun"]
Explanation: The common string with the least index sum is "Shogun" with index sum = (0 + 1) = 1.

Example 3:
Input: list1 = ["happy","sad","good"], list2 = ["sad","happy","good"]
Output: ["sad","happy"]
Explanation: There are three common strings:
"happy" with index sum = (0 + 1) = 1.
"sad" with index sum = (1 + 0) = 1.
"good" with index sum = (2 + 2) = 4.
The strings with the least index sum are "sad" and "happy".

Constraints:
1 <= list1.length, list2.length <= 1000
1 <= list1[i].length, list2[i].length <= 30
list1[i] and list2[i] consist of spaces ' ' and English letters.
All the strings of list1 are unique.
All the strings of list2 are unique.
*/
package main

import "fmt"

func findRestaurant2(list1 []string, list2 []string) []string {
	r := []string{}
	leastIndexSum := (len(list1) - 1) + (len(list2) - 1)
	maxIndex := len(list2)
	for i := 0; i < len(list1); i++ {
		for j := 0; j < maxIndex; j++ {
			if list1[i] == list2[j] {
				// 같은 스트링을 찾았으면 다음번에  같은 스트링은 두 인덱스합(i+j) 보다 작거나 같아야 한다.
				maxIndex--
				if i+j == leastIndexSum {
					r = append(r, list1[i])
				} else if i+j < leastIndexSum {
					r = r[:0]
					r = append(r, list1[i])
					leastIndexSum = i + j
				} else {
					break
				}
			}
		}
	}
	return r
}

// using map
func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	for i := 0; i < len(list2); i++ {
		m[list2[i]] = i
	}
	r := []string{}
	leastIndexSum := (len(list1) - 1) + (len(list2) - 1)
	for i := 0; i < len(list1); i++ {
		if idx, ok := m[list1[i]]; ok {
			if i+idx == leastIndexSum {
				r = append(r, list1[i])
			} else if i+idx < leastIndexSum {
				r = r[:0]
				r = append(r, list1[i])
				leastIndexSum = i + idx
			}
		}
	}
	return r
}

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(list1, list2))

	list1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 = []string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant(list1, list2))

	list1 = []string{"happy", "sad", "good"}
	list2 = []string{"sad", "happy", "good"}
	fmt.Println(findRestaurant(list1, list2))

	list1 = []string{"vacag", "KFC"}
	list2 = []string{"fvo", "xrljq", "jrl", "KFC"}
	fmt.Println(findRestaurant(list1, list2))
}
