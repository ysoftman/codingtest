package main

import "fmt"
import "os"
import "bufio"
// import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    text := scanner.Text()    
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    nLen := len(text)
    fmt.Fprintln(os.Stderr, "text:", text, "  len:", nLen)
    
    
    
    // 입력 메시지 이진화 후 모두 연결하
    strBinary := ""
    for i:=0; i<nLen; i++ {
        // 문자를 2진수로 변경
        n := int64(text[i])
		temp := strconv.FormatInt(n, 2)
        // 7비트로 맞추기
		if len(temp) < 7 {
            for i := 0; i < 7-len(temp); i++ {
				strBinary += "0"
			}
		}
        strBinary += temp
        
        fmt.Fprintf(os.Stderr, "text: %c  integer: %v  binary: %v\n", text[i], n, strBinary)
    }
    
    output := ""
    prevalue := uint8(99)
    // fmt.Fprintf(os.Stderr, "uint8(48): %v\n", uint8(48))
    // fmt.Fprintf(os.Stderr, "uint8(49): %v\n", uint8(49))
    // fmt.Fprintf(os.Stderr, "strBinary[0]: %v\n", strBinary[0])
    // fmt.Fprintf(os.Stderr, "strBinary[1]: %v\n", strBinary[1])
    for j:=0; j<len(strBinary); j++ {
        
        if strBinary[j] != prevalue {
            if j > 0 {
                output += " "
            }
            
            // 0 (48 ascii)
            if strBinary[j] == 48 {
                output += "00 "
            // 1 (49 ascii)
            } else {
                output += "0 "
            }
            
        }
        output += "0"
        
        prevalue = strBinary[j]
    }
    
    
    fmt.Println(output)// Write answer to stdout
}