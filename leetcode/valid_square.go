/*
593. Valid Square
Medium
Given the coordinates of four points in 2D space p1, p2, p3 and p4, return true if the four points construct a square.
The coordinate of a point pi is represented as [xi, yi]. The input is not given in any order.
A valid square has four equal sides with positive length and four equal angles (90-degree angles).

Example 1:
Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
Output: true

Example 2:
Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
Output: false

Example 3:
Input: p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
Output: true

Constraints:
p1.length == p2.length == p3.length == p4.length == 2
-104 <= xi, yi <= 104
*/

package main

import "fmt"

func distance(p1, p2 []int) int {
	return ((p1[0] - p2[0]) * (p1[0] - p2[0])) + ((p1[1] - p2[1]) * (p1[1] - p2[1]))
}
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p := [][]int{}
	p = append(p, p1)
	p = append(p, p2)
	p = append(p, p3)
	p = append(p, p4)
	length := make(map[int]bool)
	// 4 점의 distance 가 모두 같아야 정사각형(square)이다.
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			dist := distance(p[i], p[j])
			if dist != 0 {
				length[dist] = true
			} else {
				return false
			}
		}
	}
	// 4점이 정사격형이면, 변길이와 사선의 길이 2개만 다른 값이 된다.
	if len(length) == 2 {
		return true
	}
	return false
}

func main() {
	p1 := []int{0, 0}
	p2 := []int{1, 1}
	p3 := []int{1, 0}
	p4 := []int{0, 1}
	fmt.Println(validSquare(p1, p2, p3, p4))
	p1 = []int{0, 0}
	p2 = []int{1, 1}
	p3 = []int{1, 0}
	p4 = []int{0, 12}
	fmt.Println(validSquare(p1, p2, p3, p4))
	p1 = []int{1, 0}
	p2 = []int{-1, 0}
	p3 = []int{0, 1}
	p4 = []int{0, -1}
	fmt.Println(validSquare(p1, p2, p3, p4))
	p1 = []int{0, 0}
	p2 = []int{1, 1}
	p3 = []int{1, 0}
	p4 = []int{1, 1}
	fmt.Println(validSquare(p1, p2, p3, p4))
}
