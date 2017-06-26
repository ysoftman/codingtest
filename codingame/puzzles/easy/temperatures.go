package main

import "fmt"
import "os"
import "bufio"
//import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)
    
    fmt.Fprintln(os.Stderr, scanner.Text())
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    // fmt.Println("result")// Write answer to stdout

    scanner.Split(bufio.ScanWords)
    scanner.Scan()
        
    abMin := 5526
    Min := 5526
    for i:=0; i<n; i++ {
        strTemp := scanner.Text() // the n temperatures expressed as integers ranging from -273 to 5526
        nTemp, _ := strconv.Atoi(strTemp)
        
        fmt.Fprintln(os.Stderr, nTemp)
        
        abTemp := 0
        if nTemp < 0 {
            abTemp = nTemp * -1
        } else {
            abTemp = nTemp
        }
        
        if abTemp < abMin {
            abMin = abTemp
            Min = nTemp
        } else if abTemp == abMin {
            if nTemp > Min {
                Min = nTemp
                abMin = abTemp
            }
        }
        
        scanner.Scan()
    }
    
    if Min > 5526 {
        Min = 5526
    }
    if Min < -273 {
        Min = -273
    }
    
    if n == 0 {
        Min = 0
    }
    
    fmt.Println(Min)// Write answer to stdout
}