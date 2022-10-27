/*
https://leetcode.com/problems/image-overlap/
835. Image Overlap
Medium

You are given two images, img1 and img2, represented as binary, square matrices of size n x n. A binary matrix has only 0s and 1s as values.

We translate one image however we choose by sliding all the 1 bits left, right, up, and/or down any number of units. We then place it on top of the other image. We can then calculate the overlap by counting the number of positions that have a 1 in both images.

Note also that a translation does not include any kind of rotation. Any 1 bits that are translated outside of the matrix borders are erased.

Return the largest possible overlap.

Example 1:
Input: img1 = [[1,1,0],[0,1,0],[0,1,0]], img2 = [[0,0,0],[0,1,1],[0,0,1]]
Output: 3
Explanation: We translate img1 to right by 1 unit and down by 1 unit.
The number of positions that have a 1 in both images is 3 (shown in red).

Example 2:
Input: img1 = [[1]], img2 = [[1]]
Output: 1

Example 3:
Input: img1 = [[0]], img2 = [[0]]
Output: 0

Constraints:
n == img1.length == img1[i].length
n == img2.length == img2[i].length
1 <= n <= 30
img1[i][j] is either 0 or 1.
img2[i][j] is either 0 or 1.
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cntOverlapImg(shifty, shiftx int, img1 [][]int, img2 [][]int) int {
	leftShiftCnt1 := 0
	rightShiftCnt2 := 0
	i := 0
	for y := shifty; y < len(img1); y++ {
		j := 0
		for x := shiftx; x < len(img1); x++ {
			if img1[y][x] == 1 && img2[i][j] == 1 {
				leftShiftCnt1++
			}
			if img1[y][j] == 1 && img2[i][x] == 1 {
				rightShiftCnt2++
			}
			j++
		}
		i++
	}
	return max(leftShiftCnt1, rightShiftCnt2)
}

// 이미지 2개를 겹쳤을때 둘다 1이면 1로 계산해서 가장 많은 1이 나온 경우 찾기
func largestOverlap(img1 [][]int, img2 [][]int) int {
	r := 0
	for y := 0; y < len(img1); y++ {
		for x := 0; x < len(img2); x++ {
			// img1 을 이동시켜 오버랩했을때 카운트
			r = max(r, cntOverlapImg(y, x, img1, img2))
			// img2 를 이동시켜 오버랩했을때 카운트
			r = max(r, cntOverlapImg(y, x, img2, img1))
		}
	}
	return r
}

func main() {
	fmt.Println(largestOverlap([][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}))
	fmt.Println(largestOverlap([][]int{{1}}, [][]int{{1}}))
	fmt.Println(largestOverlap([][]int{{0}}, [][]int{{0}}))
	fmt.Println(largestOverlap([][]int{{1, 0}, {0, 0}}, [][]int{{0, 1}, {1, 0}}))
}
