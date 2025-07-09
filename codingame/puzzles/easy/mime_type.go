package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// N: Number of elements which make up the association table.
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	// Q: Number Q of file names to be analyzed.
	var Q int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &Q)

	// ext, mt 맵으로 저장해둠
	var extmap map[string]string
	extmap = make(map[string]string)

	for i := 0; i < N; i++ {
		// EXT: file extension
		// MT: MIME type.
		var EXT, MT string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &EXT, &MT)

		EXT = strings.ToLower(EXT)
		extmap[EXT] = MT
	}

	// 디버깅 확인용
	for key, value := range extmap {
		fmt.Fprintf(os.Stderr, "key:%v value:%v\n", key, value)
	}
	fmt.Fprintf(os.Stderr, "------------------\n")

	// 입력 값 판단
	for i := 0; i < Q; i++ {
		scanner.Scan()
		filename := scanner.Text() // One file name per line.

		// 입력 파일 확장자 파악
		// splitedstring := strings.Split(filename, ".")
		// fileExt := splitedstring[len(splitedstring)-1]
		idx := strings.LastIndex(filename, ".")
		fileExt := ""
		if idx >= 0 {
			fileExt = filename[idx+1:]
		}

		fileExt = strings.ToLower(fileExt)
		fmt.Fprintln(os.Stderr, "input>", filename, fileExt)

		bFind := false

		// 타입 선택
		for key, value := range extmap {
			if key == fileExt {

				fmt.Println(value)
				bFind = true
				break
			}
		}

		if bFind == false {
			fmt.Println("UNKNOWN")
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// fmt.Println("UNKNOWN") // For each of the Q filenames, display on a line the corresponding MIME type. If there is no corresponding type, then display UNKNOWN.
}
