package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var LON string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LON)

    var LAT string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LAT)
        
    fmt.Fprintf(os.Stderr, "LON:%v LAT:%v\n", LON, LAT)

    lonA := ConvertToFloat(LON)
    latA := ConvertToFloat(LAT)
    
    fmt.Fprintf(os.Stderr, "lonA:%v latA:%v\n", lonA, latA)
    
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    min := math.MaxFloat64
    result := ""
    for i := 0; i < N; i++ {
        scanner.Scan()
        
        DEFIB := scanner.Text()
        
        splitStr := strings.Split(DEFIB, ";")
        
    // 	for key, value := range splitStr {
    // 		fmt.Println("key:", key, "  value:", value)
    // 	}            
        
        lonB := ConvertToFloat(splitStr[4])
        latB := ConvertToFloat(splitStr[5])
	    
	    fmt.Fprintf(os.Stderr, "lonB:%v latB:%v\n", lonB, latB)
	    
	    // defib 까지의 거리 계산
	    temp := float64((latB - latA) / 2)
	    x := float32(lonB - lonA) * float32(math.Cos(temp))
	    x *= x
	    y := float32(latB - latA)
	    y *= y
	    d := float64(math.Sqrt(float64(x+y)) * 6371)
	    
	    fmt.Fprintf(os.Stderr, "x:%v y:%v d:%v result:%v\n", x, y, d, splitStr[1])
	    
	    if min > d {
	        min = d
	        result = splitStr[1]
	       // fmt.Fprintf(os.Stderr, "--------- min:%v result:%v\n", min, result)
	    }
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    // fmt.Println("answer")// Write answer to stdout
    
    
    fmt.Println(result)
}


// ConvertToFloat : , 를 소수점으로 계산해서 변환
func ConvertToFloat(str string) float32 {
	numberSplit := strings.Split(str, ",")
	high := numberSplit[0]
	low := numberSplit[1]

	out1, _ := strconv.Atoi(high)
	out2, _ := strconv.Atoi(low)
	result := float32(0.0)
	result += float32(out1)

	// 소수점으로 만들기
	out3 := float32(out2)
    for out3 > 1 {
		out3 = float32(out3) / float32(10)
        // fmt.Println("out1:", out1, " out2:", out2, " out3:", out3)
	}
	
	result += float32(out3)
	return result
}
