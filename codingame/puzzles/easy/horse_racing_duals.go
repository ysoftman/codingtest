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
    
    numlist := make([]int, N)
    
    for i := 0; i < N; i++ {
        var Pi int
        fmt.Scan(&Pi)
        
        numlist[i] = Pi
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    // fmt.Println("answer")// Write answer to stdout
    
    
    fmt.Fprintln(os.Stderr, "before sort:", numlist)
    
    // 우선 정렬하자
    sort.Ints(numlist)
    
    fmt.Fprintln(os.Stderr, "after sort:", numlist)
    
    // 정렬된 상태에서 한번씩만 차이을 비교
    min := numlist[N-1]
    for i:=0; i<N-1; i++ {
        if min > numlist[i+1] - numlist[i] {
            min = numlist[i+1] - numlist[i]
        }
    }
    
    fmt.Println(min)
}