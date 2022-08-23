/*
https://leetcode.com/problems/rectangle-area/
223. Rectangle Area
Medium
Given the coordinates of two rectilinear rectangles in a 2D plane, return the total area covered by the two rectangles.
The first rectangle is defined by its bottom-left corner (ax1, ay1) and its top-right corner (ax2, ay2).
The second rectangle is defined by its bottom-left corner (bx1, by1) and its top-right corner (bx2, by2).

Example 1:
Rectangle Area
Input: ax1 = -3, ay1 = 0, ax2 = 3, ay2 = 4, bx1 = 0, by1 = -1, bx2 = 9, by2 = 2
Output: 45

Example 2:
Input: ax1 = -2, ay1 = -2, ax2 = 2, ay2 = 2, bx1 = -2, by1 = -2, bx2 = 2, by2 = 2
Output: 16

Constraints:
-104 <= ax1 <= ax2 <= 104
-104 <= ay1 <= ay2 <= 104
-104 <= bx1 <= bx2 <= 104
-104 <= by1 <= by2 <= 104
*/
package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// rect1+rect2-(intersection)
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	rect1 := abs(ax2-ax1) * abs(ay2-ay1)
	rect2 := abs(bx2-bx1) * abs(by2-by1)

	// fmt.Println("rect1:", abs(ax2-ax1), abs(ay2-ay1))
	// fmt.Println("rect2:", abs(bx2-bx1), abs(by2-by1))

	x := 0
	includex := ""
	if ax1 <= bx1 && bx1 <= ax2 {
		if bx2 <= ax2 {
			x = abs(bx1 - bx2)
			includex = "b"
		} else if bx2 >= ax2 {
			x = abs(bx1 - ax2)
		}
	} else if bx1 <= ax1 && ax1 <= bx2 {
		if ax2 <= bx2 {
			x = abs(ax1 - ax2)
			includex = "a"
		} else if ax2 >= bx2 {
			x = abs(ax1 - bx2)
		}
	}

	y := 0
	includey := ""
	if ay1 <= by1 && by1 <= ay2 {
		if by2 <= ay2 {
			y = abs(by1 - by2)
			includey = "b"
		} else if by2 >= ay2 {
			y = abs(by1 - ay2)
		}
	} else if by1 <= ay1 && ay1 <= by2 {
		if ay2 <= by2 {
			y = abs(ay1 - ay2)
			includey = "a"
		} else if ay2 >= by2 {
			y = abs(ay1 - by2)
		}
	}
	// fmt.Println("intersection:", x, y)
	// rect1, rect2 둘 중 하나가 나머지 하나 전체를 포함하는 경우
	if len(includex) > 0 && includex == includey {
		if includex == "a" {
			return rect2
		}
		return rect1
	}

	// intersection(overlapped) 부분은 뺀다.
	return (rect1 + rect2) - (x * y)
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2))
	fmt.Println(computeArea(-2, -2, 2, 2, 3, 3, 4, 4))
	fmt.Println(computeArea(-2, -2, 2, 2, -3, -3, -2, -2))
}
