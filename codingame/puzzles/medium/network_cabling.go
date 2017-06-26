package main

import "fmt"
import "os"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var N int
    fmt.Scan(&N)
    
    fmt.Fprintln(os.Stderr, "num of buildings:", N)
    
    var listX []int
    var listY []int
    for i := 0; i < N; i++ {
        var X, Y int
        fmt.Scan(&X, &Y)
        // fmt.Fprintln(os.Stderr, X, Y)
        
        // X, Y 각각 리스트 구성
        listX = append(listX, X)
        listY = append(listY, Y)
    }
    
    // fmt.Fprintln(os.Stderr, "listX:", listX)
    // fmt.Fprintln(os.Stderr, "listY:", listY)
    
    // X, Y 각각 정렬
    sort.Ints(listX)
    sort.Ints(listY)
    
    // fmt.Fprintln(os.Stderr, "sorted listX:", listX)
    // fmt.Fprintln(os.Stderr, "sorted listY:", listY)
    
    var xCnt, yCnt uint64
    xCnt = 0
    yCnt = 0
    temp := 0
    
    // 가장 작은 케이블을 구성하는 main cable Y 위치 지정
    midY := listY[N/2]
    
    for j:=0;j<N;j++ {
        temp = midY - listY[j]
        if temp < 0 {
            temp *= -1
        }
        yCnt += uint64(temp)        
    }
    
    // x 방향 길이는 최대 최소 위치 차이만큼
    temp = listX[N-1] - listX[0]
    if temp < 0 {
        temp *= -1
    }
    xCnt = uint64(temp)    

    fmt.Fprintln(os.Stderr, "xCnt:", xCnt)
    fmt.Fprintln(os.Stderr, "yCnt:", yCnt)
    fmt.Println(xCnt+yCnt)
    
    // fmt.Println("answer")// Write answer to stdout
    
}
