package main

import "fmt"
import "os"
import "bufio"
import "strings"
// import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var L int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L)
    
    var H int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&H)
    
    scanner.Scan()
    T := scanner.Text()
    
    
    
    
    T = strings.ToUpper(T)    
    fmt.Fprintln(os.Stderr, "L:", L, "  H:", H, "  T:", T)


    nlength := len(T)
    fmt.Fprintln(os.Stderr, "Text Length:" , nlength)
    

        

    for i := 0; i < H; i++ {
        
        scanner.Scan()
        ROW := scanner.Text()
        
        
        RowString := ""
        for numOfChars:=0; numOfChars<nlength; numOfChars++ {
            
            startX := int(T[numOfChars]) - 65
            //  fmt.Fprintln(os.Stderr, "T[numOfChars]:", T[numOfChars], "  startX:", startX)

            if startX < 0 || startX > 26 {
                // ASCII 가 아닌 경우 ? 로 표시
                startX = 26
            }
            
            RowString += ROW[startX*L:startX*L+L]
        }
        
        fmt.Println(RowString)
    }
    

    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    
    
    // fmt.Println("answer")// Write answer to stdout
}