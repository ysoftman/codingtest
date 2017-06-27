// Spiral Array
// 문제는 다음과 같다:

// 6 6

//   0   1   2   3   4   5
//  19  20  21  22  23   6
//  18  31  32  33  24   7
//  17  30  35  34  25   8
//  16  29  28  27  26   9
//  15  14  13  12  11  10
// 위처럼 6 6이라는 입력을 주면 6 X 6 매트릭스에 나선형 회전을 한 값을 출력해야 한다.

// ysoftman
package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	fmt.Println(m, n)

	// 2차원 배열 생성
	var matrix [][]int
	matrix = make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	// // -1 초기화
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		matrix[i][j] = -1
	// 	}
	// }

	cnt := 0
	left, top, right, bottom := 0, 0, n, m
	x, y := 0, 0
	for left < right && top < bottom {
		// 오른쪽 방향으로 고고
		for x < right {
			matrix[y][x] = cnt
			x++
			cnt++
		}
		top++
		y = top
		x = right - 1
		// 아래쪽 방향으로 고고
		for y < bottom {
			matrix[y][x] = cnt
			y++
			cnt++
		}
		right--
		x = right - 1
		y = bottom - 1
		// 왼쪽 방향으로 고고
		for x >= left {
			matrix[y][x] = cnt
			x--
			cnt++
		}
		bottom--
		y = bottom - 1
		x = left
		// 위쪽 방향으로 고고
		for y >= top {
			matrix[y][x] = cnt
			y--
			cnt++
		}
		left++
		x = left
		y = top
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}

}
