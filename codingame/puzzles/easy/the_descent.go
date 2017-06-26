package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    for {
        
        max := 0
        target := 0
        for i := 0; i < 8; i++ {
            // mountainH: represents the height of one mountain, from 9 to 0.
            var mountainH int
            fmt.Scan(&mountainH)
            
            if max < mountainH {
                max = mountainH
                target = i
            }
        }
        
        fmt.Println(target)
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        // fmt.Println("4") // The number of the mountain to fire on.
    }
}
