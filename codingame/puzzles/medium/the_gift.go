package main

import "fmt"
import "os"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var N int
    fmt.Scan(&N)
    
    fmt.Fprintln(os.Stderr, "participants:", N)
    
    
    var C int
    fmt.Scan(&C)
    
    fmt.Fprintln(os.Stderr, "cost:", C)
    
    var listBudget []int
    sumBudget := 0
    for i := 0; i < N; i++ {
        var B int
        fmt.Scan(&B)
        listBudget = append(listBudget, B)
        sumBudget += B
    }
    // 예산 정렬해놓기
    sort.Ints(listBudget)
    
    // 현재 참여자들의 예산으로는 선물을 살수 없는 경우
    if sumBudget < C {
        fmt.Println("IMPOSSIBLE")// Write answer to stdout
        return
    }
    
    
    // 선물비용에 대해서 평균 내야야할 돈을 계산
    mean := C / N
    fmt.Fprintln(os.Stderr, "mean:", mean)
    
    curN := 0
    remainder := 0
    var output []int

    for i:=0;i<N;i++ {

        // 남은 비용에 대해서 현재 남음 참여자 기준으로 평균 돈을 계산한다.
        mean = C / (N - curN)
        if (N - curN) > 0 {
            remainder = C % (N - curN)
        }
        
        // 평균 돈보다 작은 예산은 예산을 모두 낸다.
        if mean >= listBudget[i] {
            output = append(output, listBudget[i])
            // 비용 다시 계산
            C -= listBudget[i]
            
        // 평균 돈보다 큰 예산은 평균 돈만 낸다.
        } else {
            output = append(output, mean)
            // 비용 다시 계산
            C -= mean
        }
        // 참여완료된 개수
        curN++

    }
    
    // 평균 금액 계산 후 나머지 금액이 있으면 맨 마지막에 더해주기
    if remainder > 0 {
        output[N-1] += remainder
    }
    
    // sort.Ints(output)
    
    // fmt.Println(output)
    for i:=0;i<N;i++ {
        fmt.Println(output[i])
    }
}

