package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func main() {
    // lightX: the X position of the light of power
    // lightY: the Y position of the light of power
    // initialTX: Thor's starting X position
    // initialTY: Thor's starting Y position
    var lightX, lightY, initialTX, initialTY int
    fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)
    
    for {
        // remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
        var remainingTurns int
        fmt.Scan(&remainingTurns)
        
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        x := ""
        y := ""
        if (initialTY == lightY) {
            y = ""
        } else if (initialTY < lightY) {
            y = "S"
            initialTY++
        } else {
            y = "N";
            initialTY--
        }    
        if (initialTX == lightX) {
            x = ""
        } else if (initialTX < lightX) {
            x = "E"
            initialTX++
        } else {
            x = "W"
            initialTX--
        }
        
        fmt.Println(y+x)
        // fmt.Println("SE") // A single line providing the move to be made: N NE E SE S SW W or NW
    }
}