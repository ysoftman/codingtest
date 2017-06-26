package main

import "fmt"
import "os"
// import "bufio"
//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // W: number of columns.
    // H: number of rows.
    var W, H int
    fmt.Scan(&W)
    fmt.Scan(&H)
    fmt.Fprintln(os.Stderr, "W:", W, "H:", H)
    
    // 타일을 저장할 그리드 구성
    var grid [][]int
    grid = make([][]int, H)
    for i := 0; i < H; i++ {
        grid[i] = make([]int, W)
    }
    for i := 0; i < H; i++ {
        for j := 0; j< W; j++ {
            tile := 0
            fmt.Scan(&tile)
            // fmt.Fprintln(os.Stderr, "tile:", tile)
            grid[i][j] = tile
        }
        
    }
    // 출구 X 위치
    ex := 0
    fmt.Scan(&ex)
    fmt.Fprintln(os.Stderr, "grid:", grid)
    
    for {
        var XI, YI int
        var POS string
        fmt.Scan(&XI)
        fmt.Scan(&YI)
        fmt.Scan(&POS)
        fmt.Fprintf(os.Stderr, "XI:%v YI:%v POS:%s\n", XI, YI, POS)
        nextX, nextY := 0, 0
          
        // 현재 위치에서 다음 타일 위치 선택하
        switch grid[YI][XI] {
            
            case 1, 3, 7, 8, 9, 12, 13:
                nextX = XI
                nextY = YI + 1
            case 2:
                if POS == "LEFT" {
                    nextX = XI + 1
                    nextY = YI
                } else {
                    nextX = XI - 1
                    nextY = YI
                }
            case 4:
                if POS == "TOP" {
                    nextX = XI -1
                    nextY = YI
                } else {
                    nextX = XI
                    nextY = YI + 1
                }
            case 5:
                if POS == "TOP" {
                    nextX = XI + 1
                    nextY = YI
                } else {
                    nextX = XI
                    nextY = YI + 1
                }
            case 6:
                if POS == "LEFT" {
                    nextX = XI + 1
                    nextY = YI
                } else {
                    nextX = XI - 1
                    nextY = YI
                }
            case 10:
                nextX = XI - 1
                nextY = YI
            case 11:
                nextX = XI + 1
                nextY = YI
        } 
        
        if nextX < 0 || nextX >= W || nextY < 0 || nextY >= H {
            fmt.Fprintln(os.Stderr, "out of range")
        }
        
        fmt.Println(nextX, nextY)
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        // fmt.Println("0 0") // One line containing the X Y coordinates of the room in which you believe Indy will be on the next turn.
    }
}
