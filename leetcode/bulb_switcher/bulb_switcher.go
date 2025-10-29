/*
https://leetcode.com/problems/bulb-switcher/
319. Bulb Switcher
Medium

There are n bulbs that are initially off.
You first turn on all the bulbs,
then you turn off every second bulb.
On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on).
For the ith round, you toggle every i bulb. For the nth round, you only toggle the last bulb.

Return the number of bulbs that are on after n rounds.

Example 1:
Input: n = 3
Output: 1
Explanation: At first, the three bulbs are [off, off, off].
After the first round, the three bulbs are [on, on, on].
After the second round, the three bulbs are [on, off, on].
After the third round, the three bulbs are [on, off, off].
So you should return 1 because there is only one bulb is on.

Example 2:
Input: n = 0
Output: 0

Example 3:
Input: n = 1
Output: 1

Constraints:
0 <= n <= 109
*/
package main

import (
	"fmt"
	"math"
)

/*
ex) n=3 0(off), 1(on), n은 전구 개수, ith 라운드
    0 0 0 최초에는 모두 off
1st 1 1 1 첫번째 라운드에서는 모두 on
2nd 1 0 1 두번째 라운드에서는 매2번째(2의배수) 전구 off
3rd 1 0 0 세번째 라운드에서는 매3번째(3의배수) 전구 toggle(on->off, off->on)
*/
// brute-force : Time Limit Exceeded
func bulbSwitch2(n int) int {
	bulbs := make([]int, n)
	// 1st
	for i := 0; i < n; i++ {
		bulbs[i] = 1
	}
	// 2nd
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			bulbs[i] = 0
		}
	}
	// 3rd ~ ith round
	for r := 3; r <= n; r++ {
		for i := 0; i < n; i++ {
			// r 번째 라운드에서는 r의 배수 전구 토글
			if i%r == r-1 {
				// if bulbs[i] == 1 {
				// 	bulbs[i] = 0
				// } else {
				// 	bulbs[i] = 1
				// }
				// not 연산 (1^x)으로 toggle
				bulbs[i] = 1 ^ bulbs[i]
			}
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if bulbs[i] == 1 {
			cnt++
		}
	}
	return cnt
}

/*
n=1 1
n=2 1
n=3 1
n=4 2  2의제곱
n=5 2
n=6 2
n=7 2
n=8 2
n=9 3  3의제곱
n=10 3
n=11 3
n=12 3
n=13 3
n=14 3
n=15 3
n=16 4 4의제곱
n=17 4
n=18 4
n=19 4
n=20 4
n=21 4
n=22 4
n=23 4
n=24 4
n=25 5 5의제곱
n=26 5
n=27 5
n=28 5
n=29 5
....
*/
// 제곱근으로 바로 구할 수 있다. O(1)
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	// for i := 1; i < 30; i++ {
	// 	fmt.Printf("n=%v %v\n", i, bulbSwitch2(i))
	// }
	fmt.Println(bulbSwitch(3))
	fmt.Println(bulbSwitch(0))
	fmt.Println(bulbSwitch(1))
	fmt.Println(bulbSwitch(10))
	fmt.Println(bulbSwitch(25))
}
