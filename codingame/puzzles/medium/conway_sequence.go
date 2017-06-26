package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var R int
    fmt.Scan(&R)

    fmt.Fprintln(os.Stderr, "Original Number:", R)    
    
    var L int
    fmt.Scan(&L)
    
    fmt.Fprintln(os.Stderr, "Line:", L)
    
    
    var prelist []int
    prelist = append(prelist, R)
    
    num := 0
    cnt := 0
    for i:=1;i<L;i++ {
        
        var curlist []int
        cnt = 0
        
        // fmt.Fprintln(os.Stderr, "(", i, ") ", prelist)
        
        // 이전 리스트에서의 숫자 및 숫자 카운트
        for j := range prelist {
            num = prelist[j]
            cnt++
            
            // fmt.Fprintln(os.Stderr, "j:", j, "prelist[j]:", prelist[j])
            // 리스트의 맨 마지막일때
            if j+1 >= len(prelist) {
                curlist = append(curlist, cnt)
                curlist = append(curlist, num)
                break
            }
            // 숫자가 달라지면 현재 숫자에 대한 카운트 기록
            if prelist[j] != prelist[j+1] {
                curlist = append(curlist, cnt)
                curlist = append(curlist, num)
                cnt = 0
            }
        }
        
        // fmt.Fprintln(os.Stderr, "(", i, ") ", prelist)

        prelist = nil
        prelist = curlist[:]
        curlist = nil
    }

    // 결과 출력
    for i := range prelist {
        fmt.Print(prelist[i])
        if i < len(prelist)-1 {
            fmt.Print(" ")
        }
    }
}
