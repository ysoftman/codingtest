package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // road: the length of the road before the gap.
    var road int
    fmt.Scan(&road)
    
    // gap: the length of the gap.
    var gap int
    fmt.Scan(&gap)
    
    // platform: the length of the landing platform.
    var platform int
    fmt.Scan(&platform)
    
    for {
        // speed: the motorbike's speed.
        var speed int
        fmt.Scan(&speed)
        
        // coordX: the position on the road of the motorbike.
        var coordX int
        fmt.Scan(&coordX)
        
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        // fmt.Println("SPEED") // A single line containing one of 4 keywords: SPEED, SLOW, JUMP, WAIT.
        
        fmt.Fprintln(os.Stderr, "road:", road, " gap:", gap, "       speed:", speed, " coordX:", coordX)
        
        // if coordX + 1 == road {
        //     fmt.Println("JUMP")            
        // } else if speed > gap {
        //     fmt.Println("SLOW")
        // } else if road < coordX {
        //     fmt.Println("SLOW")        
        // } else if speed <= gap {
        //     fmt.Println("SPEED")
        // } else {
        //     fmt.Println("WAIT")
        // }
        
        if coordX == road - 1 {
            fmt.Println("JUMP")
        } else if speed > gap + 1 || coordX > road {
            fmt.Println("SLOW")
        } else if speed < gap + 1 {
            fmt.Println("SPEED")
        } else {
            fmt.Println("WAIT")
        }
        
        
        
    }
}

