package main

import "fmt"

// permutation(순열, 순서상관있음)
// nPr : n 에서 r번 카운트 -1 하면서 곱하기 or n!/(n-r)!
// 5P3 : 5개 중 3개를 줄세우기: 5x4x3 = 60 or 5!/3! = 5x4x3x2x1/2x1 = 120/2 = 60
// 5P2 : 5개 중 2개를 줄세우기: 5x4 = 20
func permute(n, r int) int {
	// 참고 0! = 1
	result := 1
	cnt := 0
	for i := n; i > 1; i-- {
		if cnt >= r {
			break
		}
		cnt++
		result *= i
	}
	return result
}

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)
}

// n!/(n-r)! 공식 사용
func permute2(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// combination(조합,순서상관없음)
// nCr : or n!/r!(n-r)!
// 5C3 : 5개 중 3개를 조합: 5!/3!*2! = 5x4x3x2x1/(3x2x1)x(2x1) = 120/12 = 10
// 5C2 : 5개 중 2개를 조합: 5!/2!*3! = 5x4x3x2x1/(2x1)x(3x2x1) = 120/12 = 10
// 10C7 : 10개 중 7개를 조합: 10!/7!*3! = 10x9x8x7x6x5x4x3x2x1/(7x6x5x4x3x2x1)x(3x2x1) = 10x9x8/3x2x1 = 720/6 = 120
func combine(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	fmt.Println(permute(5, 3))
	fmt.Println(permute(5, 2))
	fmt.Println(permute(10, 7))
	fmt.Println(permute2(5, 3))
	fmt.Println(permute2(5, 2))
	fmt.Println(permute2(10, 7))
	fmt.Println(combine(5, 3))
	fmt.Println(combine(5, 2))
	fmt.Println(combine(10, 7))
}
