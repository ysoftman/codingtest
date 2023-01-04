/*
https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/
2244. Minimum Rounds to Complete All Tasks
Medium
You are given a 0-indexed integer array tasks, where tasks[i] represents the difficulty level of a task. In each round, you can complete either 2 or 3 tasks of the same difficulty level.

Return the minimum rounds required to complete all the tasks, or -1 if it is not possible to complete all the tasks.

Example 1:
Input: tasks = [2,2,3,3,2,4,4,4,4,4]
Output: 4
Explanation: To complete all the tasks, a possible plan is:
- In the first round, you complete 3 tasks of difficulty level 2.
- In the second round, you complete 2 tasks of difficulty level 3.
- In the third round, you complete 3 tasks of difficulty level 4.
- In the fourth round, you complete 2 tasks of difficulty level 4.
It can be shown that all the tasks cannot be completed in fewer than 4 rounds, so the answer is 4.

Example 2:
Input: tasks = [2,3,3]
Output: -1
Explanation: There is only 1 task of difficulty level 2, but in each round, you can only complete either 2 or 3 tasks of the same difficulty level. Hence, you cannot complete all the tasks, and the answer is -1.

Constraints:
1 <= tasks.length <= 105
1 <= tasks[i] <= 109
*/
package main

import "fmt"

func minimumRounds(tasks []int) int {
	m := make(map[int]int)
	// task level 별 개수 취합
	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
	}
	round := 0
	for _, cnt := range m {
		// level task 1개만 있으면 task 완료(삭제)를 할 수 없다.
		if cnt == 1 {
			return -1
		}
		// 같은 level task 에 대해 2개 또는 3개의 task 를 삭제할 수 있다.
		// task개수 5개 -> 3개처리*1번, 2개처리*1번 -> 총 2회
		// task개수 13개 -> 3개처리*3번, 2개처리*2번 -> 총 5회,
		// --> 3개처리*4번 + 1 -> 총 5회로 계산할 수 도 있다.
		if cnt%3 == 0 {
			round += cnt / 3
		} else {
			round += (cnt / 3) + 1
		}
	}
	return round
}

func main() {
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
	fmt.Println(minimumRounds([]int{2, 3, 3}))
	fmt.Println(minimumRounds([]int{1, 1, 1, 1, 1}))
	fmt.Println(minimumRounds([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minimumRounds([]int{66, 66, 63, 61, 63, 63, 64, 66, 66, 65, 66, 65, 61, 67, 68, 66, 62, 67, 61, 64, 66, 60, 69, 66, 65, 68, 63, 60, 67, 62, 68, 60, 66, 64, 60, 60, 60, 62, 66, 64, 63, 65, 60, 69, 63, 68, 68, 69, 68, 61}))
}
