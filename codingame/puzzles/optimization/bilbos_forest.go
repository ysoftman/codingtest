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

// 처음에는 space 문자로 시작한다.
// space(공백) 문자는  Z -> space -> A 에 위치한다.
// space(32)  문자를 시작 문자로
const space_char int = 32  


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    magicPhrase := scanner.Text()
    
    fmt.Fprintln(os.Stderr, "magicPhrase:", magicPhrase)
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    // fmt.Println("+.>-.")// Write action to stdout
    
    result := ""
    
    
    var stones [30]int
    for i:=0; i<30; i++ {
        stones[i] = space_char
    }
    stone_idx := 0    
    
    for i:=0; i<len(magicPhrase); i++ {


        ForwardCnt, BackwardCnt := DirDist(stones[stone_idx], int(magicPhrase[i]))
        
        fmt.Fprintf(os.Stderr, "stones[stone_idx]:%s magicPhrase[i]:%s ForwardCnt:%d BackwardCnt:%d\n",
                    string(stones[stone_idx]), string(magicPhrase[i]), ForwardCnt, BackwardCnt)
        if ForwardCnt < BackwardCnt {
            for idx:=0; idx<ForwardCnt; idx++ {
                result += "+"
            }
        } else {
            for idx:=0; idx<BackwardCnt; idx++ {
                result += "-"
            }
        }
        
        result += "."
        
        stones[stone_idx] = int(magicPhrase[i])
        
        
        // 현재 스톤들에 적혀진 문자를 탐색해서 다음 입력 문자와 가장 까가운 거리(문자돌리는 비용)를 가진 스톤을 찾는다.
        mincost := 0
        min := 4000
        for j:=0; j<30; j++ {
            ForwardCnt, BackwardCnt := DirDist(stones[j], int(magicPhrase[i]))
            if min > ForwardCnt {
                min = ForwardCnt
                mincost = j
            }
            if min > BackwardCnt {
                min = BackwardCnt
                mincost = j
            }
        }
        
        
        fmt.Fprintln(os.Stderr, "mincost:", mincost, "stone_idx:", stone_idx)
        
        if mincost < stone_idx {
            for mincost != stone_idx {
                result += ">"   
                stone_idx++
                if stone_idx >= 30 {
                    stone_idx = 0
                }
            }
        } else {
            for mincost != stone_idx {
                result += "<"
                stone_idx--
                if stone_idx < 0 {
                    stone_idx = 29
                }
            }
        }
        
        
        // if i+1 < len(magicPhrase) {
        //     if stone_idx < 30 {
        //         result += ">"
        //     } else {
        //         for a:=stone_idx; a>=0; a-- {
        //             result += "<"
        //             if stones[a] == 0 {
        //                 break
        //             }
        //         }
        //     }
        // }
    }
    
    fmt.Println(result)
 
}



func DirDist(stone_sign, input_sign int) (ForwardCnt, BackwardCnt int) {
    // 앞쪽 방향으로 탐색했을때의 거리 체크
    temp := stone_sign
    ForwardCnt = 0
    for temp != input_sign {
        temp++
        ForwardCnt++
        if temp == 'Z' + 1 {
            temp = space_char
        } 
        if temp == space_char + 1  {
            temp = 'A'
        }
    }
    // 뒤쪽 방향으로 탐색했을때의 거리 체크
    temp = stone_sign
    BackwardCnt = 0
    for temp != input_sign {
        temp--
        BackwardCnt++
        if temp == 'A' - 1 {
            temp = space_char
        }
        if temp == space_char - 1 {
            temp = 'Z'
        }
    }
    
    // fmt.Fprintf(os.Stderr, "stone_sign:%s input_sign:%s ForwardCnt:%d BackwardCnt:%d\n",
    //     string(stone_sign), string(input_sign), ForwardCnt, BackwardCnt)
    
    
    
    return ForwardCnt, BackwardCnt
}

