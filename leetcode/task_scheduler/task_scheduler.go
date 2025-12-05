/*
https://leetcode.com/problems/task-scheduler
621. Task Scheduler
Medium
You are given an array of CPU tasks, each labeled with a letter from A to Z, and a number n. Each CPU interval can be idle or allow the completion of one task. Tasks can be completed in any order, but there's a constraint: there has to be a gap of at least n intervals between two tasks with the same label.

Return the minimum number of CPU intervals required to complete all tasks.

Example 1:
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.
After completing task A, you must wait two intervals before doing A again. The same applies to task B. In the 3rd interval, neither A nor B can be done, so you idle. By the 4th interval, you can do A again as 2 intervals have passed.

Example 2:
Input: tasks = ["A","C","A","B","D","B"], n = 1
Output: 6
Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.
With a cooling interval of 1, you can repeat a task after just one other task.

Example 3:
Input: tasks = ["A","A","A", "B","B","B"], n = 3
Output: 10
Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.
There are only two types of tasks, A and B, which need to be separated by 3 intervals. This leads to idling twice between repetitions of these tasks.

Constraints:
1 <= tasks.length <= 104
tasks[i] is an uppercase English letter.
0 <= n <= 100
*/
package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	maxFreq := 0
	m := make(map[byte]int)
	for _, t := range tasks {
		m[t]++
		if maxFreq < m[t] {
			maxFreq = m[t]
		}
	}

	// 가장 많은 빈도 작업이 interval(n) 간격만큼 있는것으로 파악할 수 있다.
	// A__A__A  --> maxFreq(3-1) * n+1(A포함한 간격)
	result := (maxFreq - 1) * (n + 1)

	for _, v := range m {
		if v == maxFreq {
			result++
		}
	}

	// 예외: task 보다 작은수는 없다.
	if result < len(tasks) {
		result = len(tasks)
	}
	return result
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}
