package main

import "fmt"
import "os"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 *
*/
func main() {
    var L, H int
    fmt.Scan(&L, &H)
    
    fmt.Fprintln(os.Stderr, L, H)
    
    // 마야 문자 데이터 저장
    var mayanNum map[int]string
    mayanNum = make(map[int]string)
    for i := 0; i < H; i++ {
        var numeral string
        fmt.Scan(&numeral)
        
        fmt.Fprintln(os.Stderr, numeral)
        
        num := 0
        pre := 0
        for j:=L; j<len(numeral); j+=L {
            mayanNum[num] += numeral[pre:j]
            pre = j
            num++
        }
        mayanNum[num] += numeral[pre:len(numeral)]
    }
    
    // 문자 확인
    // for i:=0; i<len(mayanNum); i++ {
    //     fmt.Fprintln(os.Stderr, i)
    //     str := mayanNum[i]
    //     for j:=range str {
    //         if j > 0 && j % L == 0 {
    //             fmt.Fprintln(os.Stderr, "")
    //         }             
    //         fmt.Fprintf(os.Stderr, "%c", str[j])
    //     }
    //     fmt.Fprintln(os.Stderr, "")
    // }
    
    var S1 int
    fmt.Scan(&S1)
    
    fmt.Fprintln(os.Stderr, S1)
    
    num1 := 0
    str1 := ""
    temp1 := (S1/H)-1
    // operand1 숫자 10진수로 파악
    for i := 1; i <= S1; i++ {
        var num1Line string
        fmt.Scan(&num1Line)
        // fmt.Fprintln(os.Stderr, "num1Line:", num1Line)
        
        str1 += num1Line
        if i % H == 0 || i == S1 {
            // 마야문자 순자로 치환
            n := MayanToNum(mayanNum, str1)
            fmt.Fprintln(os.Stderr, "n:", n)
            // 10진수로 변환
            if temp1 > 0 {
                num1 += n * int(math.Pow(20,float64(temp1)))
            } else {
                num1 += n * 1
            }
            temp1--
            fmt.Fprintln(os.Stderr, "str1:", str1)
            fmt.Fprintln(os.Stderr, "num1:", num1)
            str1 = ""
        }
    }
    
    var S2 int
    fmt.Scan(&S2)
    
    fmt.Fprintln(os.Stderr, S2)
    
    num2 := 0
    str2 := ""
    temp2 := (S2/H)-1
    // operand2 숫자 10진수로 파악
    for i := 1; i <= S2; i++ {
        var num2Line string
        fmt.Scan(&num2Line)
        // fmt.Fprintln(os.Stderr, "num2Line:", num2Line)
        
        str2 += num2Line
        if i % H == 0 || i == S2 {
            // 마야문자 순자로 치환
            n := MayanToNum(mayanNum, str2)
            fmt.Fprintln(os.Stderr, "n:", n)
            // 10진수로 변환
            if temp2 > 0 {
                num2 += n * int(math.Pow(20,float64(temp2)))
            } else {
                num2 += n * 1
            }
            temp2--
            // fmt.Fprintln(os.Stderr, "str2:", str2)
            fmt.Fprintln(os.Stderr, "num2:", num2)
            str2 = ""
        }
    }
    
    
    // 연산자 파악
    var operation string
    fmt.Scan(&operation)
    
    fmt.Fprintln(os.Stderr, operation)
    
     
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    // fmt.Println("result")// Write answer to stdout
    result := 0
    switch operation {
        case "+":
        result = num1 + num2
        case "-":
        result = num1 - num2
        case "*":
        result = num1 * num2
        case "/":
        result = num1 / num2
    }
    // 10진수 결과
    fmt.Fprintln(os.Stderr, "result:", result)
    
    
    // resultstr := strconv.Itoa(result)
    // for i:=len(resultstr)-1; i>=0; i-- {
    // // for i:=0; i<len(resultstr); i++ {
    //     n, _ := strconv.Atoi(string(resultstr[i]))
    //     PrintMayan(mayanNum, L, n)
    // }
    
    // 마야숫자(20진수)로 변경
    var remainder []int
    for {
        remainder = append(remainder, result % 20)
        // fmt.Fprintln(os.Stderr, result % 20)
        result /= 20
        if result <= 0 {
            break
        }
    }
    // 각 숫자 출력
    for i:=len(remainder)-1;i>=0;i-- {
        fmt.Fprintln(os.Stderr, "remainder[i]:", remainder[i])
        PrintMayan(mayanNum, L, remainder[i])
    }
}



// 마야문자 -> 숫자
func MayanToNum(mayanNum map[int]string, str string) int {
    for i:= range mayanNum {
        if str == mayanNum[i] {
            return i
        }
    }
    fmt.Fprintln(os.Stderr, "err!")
    return -1
}


// 숫자에 해당하는 마야문자 출력하기
func PrintMayan(mayanNum map[int]string, width int, num int) {
    str := mayanNum[num]
    for j:=range str {
        if j>0 && j%width == 0 {
            fmt.Println("")
        }
        fmt.Printf("%c", str[j])
    }
    fmt.Println("")
}
