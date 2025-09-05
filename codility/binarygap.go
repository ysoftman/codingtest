/*
https://app.codility.com/programmers/trainings/9/binary_gap/

A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001
content_copy and contains a binary gap of length 2. The number 529 has binary representation 1000010001
content_copy and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100
content_copy and contains one binary gap of length 1. The number 15 has binary representation 1111
content_copy and has no binary gaps. The number 32 has binary representation 100000
content_copy and has no binary gaps.

class Solution { public int solution(int N); }

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.
*/
package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// Implement your solution here
	start := false
	end := false
	max := 0
	gap := 0
	for N > 1 {
		if N%2 == 1 {
			start = true
		}
		if start && N%2 == 0 {
			gap++
		}
		if start && gap > 0 && N%2 == 1 {
			end = true
		}

		if end {
			if max < gap {
				max = gap
			}
			end = false
			gap = 0
		}
		N = N / 2
	}
	if N == 1 {
		if max < gap {
			max = gap
		}
	}
	return max
}

func Solution2(N int) int {
	max := 0
	gap := 0

	start := false
	// 오른쪽shift 로 간편하게 binary 로 체크할 수 있다.
	for N > 0 {
		// 0000..1  로  마지막 자리 하나가 0 or 1 인지 판단
		if N&1 == 1 {
			if start && max < gap {
				max = gap
			}
			gap = 0
			start = true
		} else if start {
			gap++
		}
		N >>= 1
	}

	return max
}

func main() {
	for i := 0; i < 10_000; i++ {
		fmt.Println(Solution(i), Solution2(i))
		if Solution(i) != Solution2(i) {
			fmt.Println(i, "different answer")
		}
	}
}
