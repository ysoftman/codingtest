// LCD Display

// 한 친구가 방금 새 컴퓨터를 샀다. 그 친구가 지금까지 샀던 가장 강력한 컴퓨터는 공학용 전자 계산기였다. 그런데 그 친구는 새 컴퓨터의 모니터보다 공학용 계산기에 있는 LCD 디스플레이가 더 좋다며 크게 실망하고 말았다. 그 친구를 만족시킬 수 있도록 숫자를 LCD 디스플레이 방식으로 출력하는 프로그램을 만들어보자.

// 입력
// 입력 파일은 여러 줄로 구성되며 표시될 각각의 숫자마다 한 줄씩 입력된다. 각 줄에는 s와 n이라는 두개의 정수가 들어있으며 n은 출력될 숫자( 0<= n <= 99,999,999 ), s는 숫자를 표시하는 크기( 1<= s < 10 )를 의미한다. 0 이 두 개 입력된 줄이 있으면 입력이 종료되며 그 줄은 처리되지 않는다.

// 출력
// 입력 파일에서 지정한 숫자를 수평 방향은 '-' 기호를, 수직 방향은 '|'를 이용해서 LCD 디스플레이 형태로 출력한다. 각 숫자는 정확하게 s+2개의 열, 2s+3개의 행으로 구성된다. 마지막 숫자를 포함한 모든 숫자를 이루는 공백을 스페이스로 채워야 한다. 두 개의 숫자 사이에는 정확하게 한 열의 공백이 있어야 한다.

// 각 숫자 다음에는 빈 줄을 한 줄 출력한다. 밑에 있는 출력 예에 각 숫자를 출력하는 방식이 나와있다.

// 예
// 입력 예
// 2 12345
// 3 67890
// 0 0

// 출력 예
//       --   --        --
//    |    |    | |  | |
//    |    |    | |  | |
//       --   --   --   --
//    | |       |    |    |
//    | |       |    |    |
//       --   --        --

//  ---   ---   ---   ---   ---
// |         | |   | |   | |   |
// |         | |   | |   | |   |
// |         | |   | |   | |   |
//  ---         ---   ---
// |   |     | |   |     | |   |
// |   |     | |   |     | |   |
// |   |     | |   |     | |   |
//  ---         ---   ---   ---

// ysoftman
package main

import "fmt"
import "strings"

var table map[int]string

func main() {
	var s int
	var n string
	fmt.Scanf("%d %s", &s, &n)
	if s == 0 && n == "0" {
		return
	}

	table = make(map[int]string)
	table[0] = "23567890"
	table[1] = "456890"
	table[2] = "12347890"
	table[3] = "2345689"
	table[4] = "2680"
	table[5] = "134567890"
	table[6] = "2356890"

	maxline := (2 * s) + 3
	fmt.Println("maxline :", maxline)
	for row := 0; row < maxline; row++ {
		for i := 0; i < len(n); i++ {
			idx1, idx2 := getTableIndex(row, maxline)
			if idx1 >= 0 {
				if strings.Count(table[idx1], string(n[i])) > 0 {
					fmt.Printf("%s", makeMark(idx1, s))
				} else {
					for a := 0; a < s+1; a++ {
						fmt.Printf(" ")
					}
					if idx2 < 0 {
						fmt.Printf(" ")
					}
				}
			}
			if idx2 >= 0 {
				if strings.Count(table[idx2], string(n[i])) > 0 {
					fmt.Printf("%s", makeMark(idx2, s))
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}

func getTableIndex(row int, maxline int) (int, int) {
	if row == 0 {
		return 0, -1
	} else if row == maxline-1 {
		return 6, -1
	} else if row == maxline/2 {
		return 3, -1
	} else if row < maxline/2 {
		return 1, 2
	} else if row > maxline/2 {
		return 4, 5
	}
	return -1, -1
}

func makeMark(idx int, size int) string {
	mark := ""
	if idx == 0 || idx == 3 || idx == 6 {
		mark = " "
		for i := 1; i < size+1; i++ {
			mark += "-"
		}
		mark += " "
	} else if idx == 1 || idx == 4 {
		mark = "|"
		for i := 1; i < size+1; i++ {
			mark += " "
		}
	} else if idx == 2 || idx == 5 {
		mark = "|"
	}
	return mark
}
