package main

import "fmt"
import "os"


// myIntAbs : int 형 절대값 구하기
func myIntAbs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}


/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // W: width of the building.
    // H: height of the building.
    var W, H int
    fmt.Scan(&W, &H)
    
    // N: maximum number of turns before game over.
    var N int
    fmt.Scan(&N)
    
    var X0, Y0 int
    fmt.Scan(&X0, &Y0)
    
    toX, toY := X0, Y0
    
    // 탐색 범위
    left, top, right, bottom := 0, 0, W-1, H-1
    for {
        // bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
        var bombDir string
        fmt.Scan(&bombDir)
        
        fmt.Fprintf(os.Stderr, "before) left:%v top:%v right:%v bottom:%v toX:%v toY:%v\n", left, top, right, bottom, toX, toY)
        
        fmt.Fprintln(os.Stderr, "bombDir:", bombDir)    
        
        // 현재위치에서 폭탄 방향으로 움직일 위치 찾기
    	switch bombDir {
    	case "U":
    		// 다음턴 탐색 범위 줄여놓기
    		bottom = toY
    		left = toX
    		right = toX
    		
    		toY -= int(float32(myIntAbs(toY-top))/2 + 0.5)

    	case "UR":
    		bottom = toY
    		left = toX    		

    		toY -= int(float32(myIntAbs(toY-top))/2 + 0.5)
    		toX += int(float32(myIntAbs(toX-right))/2 + 0.5)
    		
    	case "R":
    		left = toX
    		top = toY
    		bottom = toY    	
    		toX += int(float32(myIntAbs(toX-right))/2 + 0.5)

    	case "DR":
    		top = toY
    		left = toX    	
    		toY += int(float32(myIntAbs(toY-bottom))/2 + 0.5)
    		toX += int(float32(myIntAbs(toX-right))/2 + 0.5)

    	case "D":
    		top = toY
    		left = toX
    		right = toX
    		toY += int(float32(myIntAbs(toY-bottom))/2 + 0.5)
    		
    	case "DL":
            top = toY
    		right = toX    	
    		
    		toY += int(float32(myIntAbs(toY-bottom))/2 + 0.5)
    		toX -= int(float32(myIntAbs(toX-left))/2 + 0.5)
    		
    	case "L":
    		right = toX
    		top = toY
    		bottom = toY
    		toX -= int(float32(myIntAbs(toX-left))/2 + 0.5)
    		
    	case "UL":
    	    bottom = toY
    	    right = toX
    		toY -= int(float32(myIntAbs(toY-top))/2 + 0.5)
    		toX -= int(float32(myIntAbs(toX-left))/2 + 0.5)    	    
    	} 
        
        fmt.Fprintf(os.Stderr, "after) left:%v top:%v right:%v bottom:%v toX:%v toY:%v\n", left, top, right, bottom, toX, toY)
        
        fmt.Println(toX, toY)
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        // fmt.Println("0 0") // the location of the next window Batman should jump to.
    }
}

