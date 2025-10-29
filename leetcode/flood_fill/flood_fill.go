/*
https://leetcode.com/problems/flood-fill/
733. Flood Fill
Easy
An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].
To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with newColor.
Return the modified image after performing the flood fill.

Example 1:
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

Example 2:
Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
Output: [[2,2,2],[2,2,2]]
*/
package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	lenY := len(image)
	lenX := len(image[0])
	startColor := image[sr][sc]
	stack := [][]int{}
	stack = append(stack, []int{sr, sc})
	for len(stack) > 0 {
		// top
		y := stack[len(stack)-1][0]
		x := stack[len(stack)-1][1]
		// pop
		stack = stack[:len(stack)-1]
		image[y][x] = newColor
		if y-1 >= 0 && image[y-1][x] == startColor && image[y-1][x] != newColor {
			image[y-1][x] = newColor
			stack = append(stack, []int{y - 1, x})
		}
		if x-1 >= 0 && image[y][x-1] == startColor && image[y][x-1] != newColor {
			image[y][x-1] = newColor
			stack = append(stack, []int{y, x - 1})
		}
		if y+1 < lenY && image[y+1][x] == startColor && image[y+1][x] != newColor {
			image[y+1][x] = newColor
			stack = append(stack, []int{y + 1, x})
		}
		if x+1 < lenX && image[y][x+1] == startColor && image[y][x+1] != newColor {
			image[y][x+1] = newColor
			stack = append(stack, []int{y, x + 1})
		}
	}
	return image
}

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 0, 0, 2))
}
