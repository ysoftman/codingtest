/*
https://leetcode.com/problems/reduce-array-size-to-the-half/
1338. Reduce Array Size to The Half
Medium
You are given an integer array arr. You can choose a set of integers and remove all the occurrences of these integers in the array.

Return the minimum size of the set so that at least half of the integers of the array are removed.

Example 1:
Input: arr = [3,3,3,3,5,5,5,2,2,7]
Output: 2
Explanation: Choosing {3,7} will make the new array [5,5,5,2,2] which has size 5 (i.e equal to half of the size of the old array).
Possible sets of size 2 are {3,5},{3,2},{5,2}.
Choosing set {2,7} is not possible as it will make the new array [3,3,3,3,5,5,5] which has a size greater than half of the size of the old array.

Example 2:
Input: arr = [7,7,7,7,7,7]
Output: 1
Explanation: The only possible set you can choose is {7}. This will make the new array empty.

Constraints:
2 <= arr.length <= 105
arr.length is even.
1 <= arr[i] <= 105
*/
package main

import (
	"fmt"
	"sort"
)

// time complexity : O(NlogN)
func minSetSize(arr []int) int {
	m := make(map[int]int)
	// 원소들의 개수파악해서 해시맵에 저장
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	// fmt.Println("m:", m)
	v := []int{}

	// 참고로 c++ 등의 multi set 이나 priority queue를 이용하면 데이터를 추가하면서 정렬되서 더 빠를것 같음
	// 가장 많이 카운트된 원소순으로 정렬
	for i := range m {
		v = append(v, m[i])
	}
	// log(n log n)
	sort.Slice(v, func(a, b int) bool {
		return v[a] > v[b]
	})
	// fmt.Println("v:", v)
	sum := 0
	cnt := 0
	// v[i](원소 개수)를 큰것부터 더해서 arr/2 보다 커질때까지의 원소 개수를 카운하면 된다.
	for i := 0; i < len(v); i++ {
		cnt++
		sum += v[i]
		if len(arr)/2 <= sum {
			return cnt
		}
	}
	return 0
}

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
	fmt.Println(minSetSize([]int{7, 7, 7, 7, 7, 7}))
}
