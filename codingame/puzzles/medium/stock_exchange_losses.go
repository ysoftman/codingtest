package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {

    var n int
    nums := []int64{}
    fmt.Scanf("%d", &n)
    fmt.Fprintln(os.Stderr, "n:", n)
    
    var data int64
    for i:=0;i<n;i++ {
        fmt.Scan(&data)
        nums = append(nums, data)
        fmt.Fprintln(os.Stderr, "data:", data)
    }

    var p int64
    p = 0
    t0, t1 := 0, 0
    
    var pre int64
    pre = 0
    for i := 0; i < n; i++ {    
        // 필요없는 반복은 없앤다.
        if pre > nums[i] {
            continue
        }
        for j := i+1; j < n; j++ {        
            // 손해가 극대화 되는 t0(buying time). t1(selling time) 찾기
            // 가장 작은 손해 케이스 찾기
            if nums[i] > nums[j] {
                if nums[i] - nums[j] > p {
                    p = nums[i] - nums[j]
                }
            }
        }
        pre = nums[i]
    }
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "t0[%v]:%v t1[%v]:%v\n", t0, nums[t0], t1, nums[t1])
    // fmt.Fprintf(os.Stderr, "min:%v max:%v\n", min, max)
    
    // 손해를 표시하기위해 -1 곱한다.
    p *= -1
    fmt.Println(p)
    

    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    // fmt.Println("answer")// Write answer to stdout
        
}
