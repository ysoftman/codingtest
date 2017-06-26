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

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var L, C int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L, &C)
    
    fmt.Fprintln(os.Stderr, L, C)
    
    var themap [][]string
    themap = make([][]string, L)
    
    var teleporty []int
    var teleportx []int
    startx, starty := 0, 0
    // 전체 지도 구성하기
    for i := 0; i < L; i++ {
        themap[i] = make([]string, C)
        scanner.Scan()
        row := scanner.Text()
        for j:=range row {
            themap[i][j] = string(row[j])
            // 시작위치 파악해 놓기
            if themap[i][j] == "@" {
                starty = i
                startx = j
                fmt.Fprintln(os.Stderr, "start(y,x):", starty, startx)
            }
            
            // T 위치 파악해 두기
            if themap[i][j] == "T" {
                teleporty = append(teleporty, i)
                teleportx = append(teleportx, j)
            }
        }
    }
    
    // for debug
    for i:=range themap {
        for j:=range themap[i] {
            fmt.Fprint(os.Stderr, themap[i][j])
        }
        fmt.Fprintln(os.Stderr, "")
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    // fmt.Println("answer")// Write answer to stdout
    
    i:=starty
    j:=startx
    curdir:="SOUTH"
    bBreakerMode:=false
    bInvertMode:=false
    var path []string
    var pathstr []string
    for {
        // 출구 찾으면 끝내기
        if themap[i][j] == "$" {
            fmt.Fprintln(os.Stderr, "exit~")
            break
        }
        
        // teleport 이면 반대쪽 teleport 를 찾아서 이동
        if themap[i][j] == "T" {
            idx := 0
            for idx = range teleporty {
                if teleporty[idx] != i {
                    break
                }
            }
            i = teleporty[idx]
            j = teleportx[idx]        
            fmt.Fprintln(os.Stderr, "teleport:", i, j)
        }
        // I  로 invert mode 시작 종료
        if themap[i][j] == "I"{
            if bInvertMode {
                bInvertMode = false
            } else {
                bInvertMode = true
            }   
        }
        
        // B 로 breaker mode 시작 종료
        if themap[i][j] == "B"{
            if bBreakerMode {
                bBreakerMode = false
            } else {
                bBreakerMode = true
            }   
        }
        
        // breaker 모드의 경우 다음 방향위체 X 는 부셔버릴 수 있다.
        if bBreakerMode {
            bDestory := false
            if curdir == "SOUTH" && themap[i+1][j] == "X" {
                i++
                bDestory = true
            } else if curdir == "EAST" && themap[i][j+1] == "X" {
                j++
                bDestory = true
            } else if curdir == "NORTH" && themap[i-1][j] == "X" {
                i--
                bDestory = true
            } else if curdir == "WEST" && themap[i][j-1] == "X" {
                j--
                bDestory = true
            }
            if bDestory {
                themap[i][j] = " "
                fmt.Fprintln(os.Stderr, i, j)
                // fmt.Println(curdir)
                pathstr = append(pathstr, curdir)
                curxy := fmt.Sprintf("%s %s", i, j)
                path = append(path, curxy)                
                continue
            }
        }
        
        if themap[i][j] == " " || themap[i][j] == "@" || themap[i][j] == "B" || themap[i][j] == "I" || themap[i][j] == "T" {

            if bInvertMode {
                if curdir == "WEST" && IsNotBlock(themap[i][j-1]) {
                    j--
                } else if curdir == "NORTH" && IsNotBlock(themap[i-1][j]) {
                    i--
                } else if curdir == "EAST" && IsNotBlock(themap[i][j+1]) {
                    j++                
                } else if curdir == "SOUTH" && IsNotBlock(themap[i+1][j]) {
                    i++
                } else if IsNotBlock(themap[i][j-1]) {
                    curdir = "WEST"
                    j--
                } else if IsNotBlock(themap[i-1][j]) {
                    curdir = "NORTH"
                    i--
                } else if IsNotBlock(themap[i][j+1]) {
                    curdir = "EAST"
                    j++
                } else if IsNotBlock(themap[i+1][j]) {
                    curdir = "SOUTH"
                    i++                
                }
            } else {
                if curdir == "SOUTH" && IsNotBlock(themap[i+1][j]) {
                    i++
                } else if curdir == "EAST" && IsNotBlock(themap[i][j+1]) {
                    j++
                } else if curdir == "NORTH" && IsNotBlock(themap[i-1][j]) {
                    i--
                } else if curdir == "WEST" && IsNotBlock(themap[i][j-1]) {
                    j--
                } else if IsNotBlock(themap[i+1][j]) {
                    curdir = "SOUTH"
                    i++
                } else if IsNotBlock(themap[i][j+1]) {
                    curdir = "EAST"
                    j++
                } else if IsNotBlock(themap[i-1][j]) {
                    curdir = "NORTH"
                    i--
                } else if IsNotBlock(themap[i][j-1]) {
                    curdir = "WEST"
                    j--
                }
            }
        } else if themap[i][j] == "S" && IsNotBlock(themap[i+1][j]) {
            i++
            curdir="SOUTH"
        } else if themap[i][j] == "E" && IsNotBlock(themap[i][j+1]) {
            j++
            curdir="EAST"
        } else if themap[i][j] == "N" && IsNotBlock(themap[i-1][j]) {
            i--
            curdir="NORTH"
        } else if themap[i][j] == "W" && IsNotBlock(themap[i][j-1]) {
            j--
            curdir="WEST"
        }
        
        fmt.Fprintln(os.Stderr, i, j)
        // fmt.Println(curdir)
        pathstr = append(pathstr, curdir)


        // loop 상태라면 끝내기 
        if len(path) > L * C {
            pathstr = nil
            pathstr = append(pathstr, "LOOP")
            break
        }
        curxy := fmt.Sprintf("%s %s", i, j)
        path = append(path, curxy)
        
    }
    
    for idx := range pathstr {
        fmt.Println(pathstr[idx])
    }
}


func IsNotBlock(temp string) bool {
    if temp == "#" || temp == "X" {
        return false
    }
    return true
}

