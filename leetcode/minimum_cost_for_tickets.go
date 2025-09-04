/*
https://leetcode.com/problems/minimum-cost-for-tickets/
983. Minimum Cost For Tickets
Medium

You have planned some train traveling one year in advance. The days of the year in which you will travel are given as an integer array days. Each day is an integer from 1 to 365.

Train tickets are sold in three different ways:

a 1-day pass is sold for costs[0] dollars,
a 7-day pass is sold for costs[1] dollars, and
a 30-day pass is sold for costs[2] dollars.
The passes allow that many days of consecutive travel.

For example, if we get a 7-day pass on day 2, then we can travel for 7 days: 2, 3, 4, 5, 6, 7, and 8.
Return the minimum number of dollars you need to travel every day in the given list of days.

Example 1:

Input: days = [1,4,6,7,8,20], costs = [2,7,15]
Output: 11
Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
In total, you spent $11 and covered all the days of your travel.
Example 2:

Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
Output: 17
Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 30-day pass for costs[2] = $15 which covered days 1, 2, ..., 30.
On day 31, you bought a 1-day pass for costs[0] = $2 which covered day 31.
In total, you spent $17 and covered all the days of your travel.

Constraints:
1 <= days.length <= 365
1 <= days[i] <= 365
days is in strictly increasing order.
costs.length == 3
1 <= costs[i] <= 1000
*/

/*
days = [1,4,6,7,8,20], costs = [2,7,15]
1일 이용권 비용 : 2
7일 일이용권 비용 : 7
30일 이용권 비용 : 15

x = min( (1일전 최소 비용 + 1일 이용권 비용 2), (7일전 최소 이용 비용 + 7일 이용권 비용 7), (30일전 최소 비용 + 30일 이용권 비용 15) )

1 일 -> x = 2
1,4 일 -> x = 4
1,4,6 일 -> x = 6
1,4,6,7 일 -> x = 7
1,4,6,7,8 일 -> x = 9
1,4,6,7,8,20 일 -> x = 11
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// time complexity : O(n)
func mincostTickets(days []int, costs []int) int {
	// 찾기을 빠르게 하기 위해 해시맵 생성, 시간 복잡도 O(1)
	m := make(map[int]bool)
	for _, d := range days {
		m[d] = true
	}

	dp := make([]int, 366)
	for i := 1; i < 366; i++ {
		if _, exist := m[i]; !exist {
			// days 없는 날짜는 이전 dp 값을 유지
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = min(dp[i-1]+costs[0], dp[max(i-7, 0)]+costs[1])
		dp[i] = min(dp[i], dp[max(i-30, 0)]+costs[2])
	}
	return dp[365]
}

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}
