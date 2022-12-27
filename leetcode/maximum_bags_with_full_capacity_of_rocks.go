/*
https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/
2279. Maximum Bags With Full Capacity of Rocks
Medium
You have n bags numbered from 0 to n - 1. You are given two 0-indexed integer arrays capacity and rocks. The ith bag can hold a maximum of capacity[i] rocks and currently contains rocks[i] rocks. You are also given an integer additionalRocks, the number of additional rocks you can place in any of the bags.

Return the maximum number of bags that could have full capacity after placing the additional rocks in some bags.

Example 1:
Input: capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2
Output: 3
Explanation:
Place 1 rock in bag 0 and 1 rock in bag 1.
The number of rocks in each bag are now [2,3,4,4].
Bags 0, 1, and 2 have full capacity.
There are 3 bags at full capacity, so we return 3.
It can be shown that it is not possible to have more than 3 bags at full capacity.
Note that there may be other ways of placing the rocks that result in an answer of 3.

Example 2:
Input: capacity = [10,2,2], rocks = [2,2,0], additionalRocks = 100
Output: 3
Explanation:
Place 8 rocks in bag 0 and 2 rocks in bag 2.
The number of rocks in each bag are now [10,2,2].
Bags 0, 1, and 2 have full capacity.
There are 3 bags at full capacity, so we return 3.
It can be shown that it is not possible to have more than 3 bags at full capacity.
Note that we did not use all of the additional rocks.

Constraints:
n == capacity.length == rocks.length
1 <= n <= 5 * 104
1 <= capacity[i] <= 109
0 <= rocks[i] <= capacity[i]
1 <= additionalRocks <= 109
*/
package main

import (
	"fmt"
	"sort"
)

/*
Greedy algorithm
n 개의 bag이 있고, 각 bag 이 rock 을 담을 수 있는 용량은 capacity 로 나타냈다.
현재 bag(capacity)[i] 에는 rocks[i] 의 rock 이 채워져 있다.
rocks + additionalRocks 모든 rock 을 bag 에 꽉 채울때, 꽉찬 가방의 개수는?

Input: capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2

이미 채워진 돌을 빼고 남은 용량

	capacity [2,3,4,5]

-      rocks [1,2,4,4]
freeCapacity [1,1,0,1]

정렬 후 앞에서 부터 additionalRocks 을 채워 나간다
freeCapacity [0,1,1,1] additionalRocks=2, 첫번째 백 꽉참
freeCapacity [0,0,1,1] additionalRocks=1, 두번째 백 꽉참
freeCapacity [0,0,0,1] additionalRocks=0, 세번째 백 꽉참
남음 rock 없음, 꽉찬 백 개수 3
*/
func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i := 0; i < len(capacity); i++ {
		capacity[i] -= rocks[i]
	}
	sort.Ints(capacity)
	fullbags := 0
	for i := 0; i < len(capacity); i++ {
		if additionalRocks >= capacity[i] {
			additionalRocks -= capacity[i]
			capacity[i] = 0
			fullbags++
		} else {
			break
		}
	}
	return fullbags
}

func main() {
	fmt.Println(maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2))
	fmt.Println(maximumBags([]int{10, 2, 2}, []int{2, 2, 0}, 100))
}
