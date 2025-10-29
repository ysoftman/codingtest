/*
https://leetcode.com/problems/pancake-sorting
969. Pancake Sorting
Medium

Given an array of integers arr, sort the array by performing a series of pancake flips.

In one pancake flip we do the following steps:

Choose an integer k where 1 <= k <= arr.length.
Reverse the sub-array arr[0...k-1] (0-indexed).
For example, if arr = [3,2,1,4] and we performed a pancake flip choosing k = 3, we reverse the sub-array [3,2,1], so arr = [1,2,3,4] after the pancake flip at k = 3.

Return an array of the k-values corresponding to a sequence of pancake flips that sort arr. Any valid answer that sorts the array within 10 * arr.length flips will be judged as correct.

Example 1:

Input: arr = [3,2,4,1]
Output: [4,2,4,3]
Explanation:
We perform 4 pancake flips, with k values 4, 2, 4, and 3.
Starting state: arr = [3, 2, 4, 1]
After 1st flip (k = 4): arr = [1, 4, 2, 3]
After 2nd flip (k = 2): arr = [4, 1, 2, 3]
After 3rd flip (k = 4): arr = [3, 2, 1, 4]
After 4th flip (k = 3): arr = [1, 2, 3, 4], which is sorted.
Example 2:

Input: arr = [1,2,3]
Output: []
Explanation: The input is already sorted, so there is no need to flip anything.
Note that other answers, such as [3, 3], would also be accepted.

Constraints:
1 <= arr.length <= 100
1 <= arr[i] <= arr.length
All integers in arr are unique (i.e. arr is a permutation of the integers from 1 to arr.length).
*/

/*
i=마지막 원소부터 ~ 0순으로
i번째에 제 값을 찾기 위해선 해시팸에서 i번째 값을 위치를 찾아 0번째로 보내고(reverse), 0번째에서 i번째로 보낸다.(reverse)
[3 2 4 1]
[4 2 3 1]
[1 3 2 4]
[3 1 2 4]
[2 1 3 4]
[2 1 3 4]
[1 2 3 4]

example1 과 output 이 다르지만(여러 방법이 있기 때문에) 10*arr.length 길이내에서 정렬이 되면 된다.
*/

package main

import "fmt"

func reverse(arr []int, k int) {
	j := k - 1
	for i := 0; i < k; i++ {
		if i >= j {
			arr[i], arr[j] = arr[j], arr[i]
		}
		j--
	}
	// fmt.Println(arr)
}

func pancakeSort(arr []int) []int {
	result := []int{}

	m := make(map[int]int)

	// update idx position map
	for idx := 0; idx < len(arr); idx++ {
		m[arr[idx]] = idx
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i]-1 == i {
			continue
		}

		// 현재 위치 i에 있어야할 값을 0번째로 reverse 후
		k := m[i+1] + 1
		result = append(result, k)
		reverse(arr, k)
		// reverse 하면 0 번째 값이 i 위치로 제자리를 찾는다.
		result = append(result, i+1)
		reverse(arr, i+1)

		// update idx position map
		for idx := 0; idx < len(arr); idx++ {
			m[arr[idx]] = idx
		}
	}
	return result
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
	fmt.Println(pancakeSort([]int{3, 5, 6, 7, 2, 4, 1}))
}
