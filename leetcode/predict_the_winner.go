/*
https://leetcode.com/problems/predict-the-winner/
486. Predict the Winner
Medium
You are given an integer array nums. Two players are playing a game with this array: player 1 and player 2.

Player 1 and player 2 take turns, with player 1 starting first. Both players start the game with a score of 0. At each turn, the player takes one of the numbers from either end of the array (i.e., nums[0] or nums[nums.length - 1]) which reduces the size of the array by 1. The player adds the chosen number to their score. The game ends when there are no more elements in the array.

Return true if Player 1 can win the game. If the scores of both players are equal, then player 1 is still the winner, and you should also return true. You may assume that both players are playing optimally.

Example 1:
Input: nums = [1,5,2]
Output: false
Explanation: Initially, player 1 can choose between 1 and 2.
If he chooses 2 (or 1), then player 2 can choose from 1 (or 2) and 5. If player 2 chooses 5, then player 1 will be left with 1 (or 2).
So, final score of player 1 is 1 + 2 = 3, and player 2 is 5.
Hence, player 1 will never be the winner and you need to return false.

Example 2:
Input: nums = [1,5,233,7]
Output: true
Explanation: Player 1 first chooses 1. Then player 2 has to choose between 5 and 7. No matter which number player 2 choose, player 1 can choose 233.
Finally, player 1 has more score (234) than player 2 (12), so you need to return True representing player1 can win.

Constraints:
1 <= nums.length <= 20
0 <= nums[i] <= 107
*/
package main

import "fmt"

/*
단순히 left, right 둘중 큰수를 택하면 player2 더 큰수를 가져잘 수 있어 optimal 한 방법이 아니다.
min-max algorithm
그림 설명
https://leetcode.com/problems/predict-the-winner/solutions/127634/Figures/486/486_Predict_the_winner_new.PNG
*/
func turnValue(nums []int, left, right, turn int) int {
	if left == right {
		return turn * nums[left]
	}
	// left/right 선택을때의 값 = player1 turn 일때 + 와 다음 player2 turn 일때(player1 입장에선 가져올 수 없는 값이니) - 로 해서 합산
	chooseLeft := turn*nums[left] + turnValue(nums, left+1, right, -turn)
	chooseRight := turn*nums[right] + turnValue(nums, left, right-1, -turn)

	return turn * max(turn*chooseLeft, turn*chooseRight)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func PredictTheWinner(nums []int) bool {
	return turnValue(nums, 0, len(nums)-1, 1) >= 0
}

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}
