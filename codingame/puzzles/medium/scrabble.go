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

    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    fmt.Fprintln(os.Stderr, "N:", N)
    
    var wordlist []string
    for i := 0; i < N; i++ {
        scanner.Scan()
        W := scanner.Text()
        wordlist = append(wordlist, W)
    }
    // fmt.Fprintln(os.Stderr, "wordlist:", wordlist)
    
    scanner.Scan()
    letters := scanner.Text()
    fmt.Fprintln(os.Stderr, "letters:", letters)
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    // fmt.Println("invalid word")// Write answer to stdout
        
    
    maxpoints := 0
    maxword := ""
    for i := range wordlist {
        
        str := wordlist[i]
        
        // fmt.Fprintln(os.Stderr, "str:", str)
        
        // string 은 변경불가능하기때문에 byte 슬라이스로 변환해서 변경하도록한다.
        temp := []byte(letters)
        
        // letter 로 현재 word 를 만들수 있는지 파악
        bCheck := true    
        for x := range str {
            bFind := false
            for y := range temp {
                if str[x] == temp[y] {
                    bFind = true
                    // 한번 비교해서 있는 단어는 중복될 수 없다.
                    temp[y] = '_'
                    break
                }
            }
            if bFind == false {
                bCheck = false
                break
            }
            
        }
        
        // letters 로 만들수 없는 단어는 스
        if bCheck == false {
            // fmt.Fprintln(os.Stderr, "skip")
            continue
        }
        // fmt.Fprintln(os.Stderr, "check")
        
        points := 0
        for j := range str {
            switch str[j] {
                case 'e', 'a', 'i', 'o', 'n', 'r', 't', 'l', 's', 'u':
                points += 1
                case 'd', 'g':
                points += 2
                case 'b', 'c', 'm', 'p':
                points += 3
                case 'f', 'h', 'v', 'w','y':
                points += 4
                case 'k':
                points += 5
                case 'j', 'x':
                points += 8
                case 'q', 'z':
                points += 10
            }
        }
        
        fmt.Fprintln(os.Stderr, "word:", str, " points:", points)
        
        if maxpoints < points {
            maxpoints = points
            maxword = str
            // fmt.Fprintln(os.Stderr, "word:", str, " points:", points)
        }
    }
        
    fmt.Fprintln(os.Stderr, "word:", maxword, " points:", maxpoints)
    fmt.Println(maxword)
    
}
