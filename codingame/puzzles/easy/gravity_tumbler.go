package main

import "fmt"
import "os"
import "bufio"

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func countHeavyBits(scanline string) int {
	cnt := 0
	for i := 0; i < len(scanline); i++ {
		if scanline[i] == '#' {
			cnt++
		}
	}
	return cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var width, height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)

	// 90도 반시계 방향으로 덤블링한 횟수
	var count int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &count)

	outmem := []int{}
	for i := 0; i < height; i++ {
		scanner.Scan()
		//raster := scanner.Text()
		// fmt.Println(scanner.Text())
		outmem = append(outmem, countHeavyBits(scanner.Text()))
	}
	count %= 4
	switch count {
	case 1, 3:
		for w := 0; w < width; w++ {
			for i := 0; i < len(outmem); i++ {
				if outmem[i] >= width-w {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	case 0, 2:
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				if w >= width-outmem[h] {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	}
}
