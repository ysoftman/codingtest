package main

import "fmt"
import "os"
import "bufio"
//import "strings"
import "strconv"

/**
 * Don't let the machines win. You are humanity's last hope...
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // width: the number of cells on the X axis
    var width int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&width)
    
    // height: the number of cells on the Y axis
    var height int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&height)
    
    
    fmt.Fprintf(os.Stderr, "width:%v height:%v\n", width, height)
    
    // 노드를 저장할 2차원 배열 생성    
    var node [][]int
    node = make([][]int, height)
    for i:=0; i<height; i++ {
        node[i] = make([]int, width)
    }
    
    for i := 0; i < height; i++ {
        scanner.Scan()
        line := scanner.Text() // width characters, each either 0 or .
        // 노드가 존재하면 0 존재하지 않으면 . 으료 표시
        fmt.Fprintln(os.Stderr, line)
        // 노드 정보 2차원으로 저장
        for j:=0; j<width; j++ {
            if string(line[j]) == "." {
                node[i][j] = -1
            } else {
                num, _ := strconv.Atoi(string(line[j]))
                node[i][j] = num
            }
            
        }
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Fprintln(os.Stderr, node)

    for y:=0; y<height;y++ {
        for x:=0; x<width; x++ {
            
            if node[y][x] < 0 {
                continue
            }
                
            // 현재 노드에서 x+1, y+1 위치에 노드 존재여부 판단하여 기록
            output := strconv.Itoa(x) + " " + strconv.Itoa(y)
                        
            x2 := -1
            y2 := -1
            for i:=1; x+i<width; i++ {
                if node[y][x+i] == 0 {
                    x2 = x+i
                    y2 = y
                    break
                }
            }
            output += " " + strconv.Itoa(x2) + " " + strconv.Itoa(y2)
            
        
            x3 := -1
            y3 := -1
            for i:=1; y+i<height; i++ {
                if node[y+i][x] == 0 {
                    x3 = x
                    y3 = y+i
                    break
                }
            }
            output += " " + strconv.Itoa(x3) + " " + strconv.Itoa(y3)
            
            // 현재 노드에서의 결과 출력            
            fmt.Println(output)
            
        }
    }
    
    // fmt.Println("0 0 1 0 0 1") // Three coordinates: a node, its right neighbor, its bottom neighbor
    
}