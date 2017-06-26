package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

// 노드 자료구조
type node struct {
    link []int
}

func main() {
    // n: the number of relationships of influence
    var n int
    fmt.Scan(&n)
    
    fmt.Fprintln(os.Stderr, "the number of relationships of influence:", n)
    
    // 그래프 구성하기
    var graph map[int](node)
    graph = make(map[int](node))
    for i := 0; i < n; i++ {
        // x: a relationship of influence between two people (x influences y)
        var x, y int
        fmt.Scan(&x, &y)
        
        fmt.Fprintln(os.Stderr, x, y)
        
        // 노드에 링크 정보 추가
        graph[x] = node{append(graph[x].link, y)}
        
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")    
    // fmt.Println("2") // The number of people involved in the longest succession of influences
    
    // for debug
    // for i:= range graph {
    //     fmt.Fprintln(os.Stderr, "for debug, graph[", i, "].link:", graph[i].link)
    // }
 
    depth := 0
    max := 0
    result := 0
    for i := range graph {
        depth, result = 1, 1
        
        fmt.Fprintln(os.Stderr, "graph[", i, "]")
        search(graph, graph[i], depth, &result)
        
        if max < result {
            max = result
        }
    }
    
    fmt.Println(max)
       
}

// 노드 링크 따라가면서 최대 길이 체크
func search(graph map[int](node), n node, depth int, result *int) {
    
    depth++
    
    // 최대 depth 를 기억해둔다.
    if *result < depth {
        *result = depth
    }
    
    fmt.Fprintln(os.Stderr, "depth:", depth, "node:", n)
    
    for j:= range n.link {
        value := n.link[j]
        
        // fmt.Fprintln(os.Stderr, "value:", value)
        
        _, exist := graph[value]
        if exist {
            fmt.Fprintln(os.Stderr, "graph[", value, "] is exist")
            search(graph, graph[value], depth, result)
        } else {
            fmt.Fprintln(os.Stderr, "graph[", value, "] is not exist")
        }
    }
}
