package main

import "fmt"
import "os"
// import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/


type node struct {
    num int
    propagation_base int    // 전파를 시작하는 노드(중복 경로를 제외하는 용도로 사용)
    link []int
}



// type DataList map[int]node

// func (receiver DataList) Less(a, b int) bool {
//     if len(receiver[a].link) > len(receiver[b].link) {
//         return true
//     }
//     return false
// }

// func (receiver DataList) Len() int {
//     return len(receiver)
// }

// func (receiver DataList) Swap(a, b int) {
//     receiver[a], receiver[b] = receiver[b], receiver[a]
// }

func main() {
    // n: the number of adjacency relations
    var n int
    fmt.Scan(&n)
    
    fmt.Fprintln(os.Stderr, "n:", n)
    
    
    
    
    var graph map[int]node
    graph = make(map[int]node)

    // var graph_temp map[int]node
    // graph_temp = make(map[int]node)
    
    // 그래프 구조 생성
    for i := 0; i < n; i++ {
        // xi: the ID of a person which is adjacent to yi
        // yi: the ID of a person which is adjacent to xi
        var xi, yi int
        fmt.Scan(&xi, &yi)
        
        // fmt.Fprintln(os.Stderr, xi, yi)
        graph[xi] = node {
            num:xi,
            propagation_base:-1,
            link:append(graph[xi].link, yi),
        }
        graph[yi] = node {
            num:yi,
            propagation_base:-1,
            link:append(graph[yi].link, xi),
        }
        
        // graph_temp[xi] = node {
        //     num:xi,
        //     propagation_base:-1,
        //     link:append(graph_temp[xi].link, yi),
        // }
        // graph_temp[yi] = node {
        //     num:yi,
        //     propagation_base:-1,
        //     link:append(graph_temp[yi].link, xi),
        // }  
        
    }

    result := len(graph)
    


    // for i:=0;i<len(graph_temp);i++ {
    //     fmt.Fprintln(os.Stderr, "graph_temp[",i,"]:", graph_temp[i])
    // }
    
    // sort.Sort(DataList(graph_temp))
    
    // fmt.Fprintln(os.Stderr, "-----")
    
    // var sortedindex []int
    // for i:=0;i<len(graph_temp);i++ {
    //     // fmt.Fprintln(os.Stderr, "graph_temp[",i,"]:", graph_temp[i])
    //     sortedindex = append(sortedindex, graph_temp[i].num)
    // }
    
    // fmt.Fprintln(os.Stderr, "-----")
    
    // for i:=0;i<len(sortedindex);i++ {
    //     fmt.Fprintln(os.Stderr, "sortedindex[",i,"]:", sortedindex[i])
    // }

    depth := 0
    max := 0    
    bEnd := false
    for i:=0; i<len(graph); i++ {
        
        
        // 말단 노드로 탐색하면 최대 전파 경로를 알 수 있다.
        if len(graph[i].link) == 1 {
            TraverseGraph(graph, i, i, depth, &max, result, &bEnd)
            fmt.Fprintln(os.Stderr, "max:", max)
            break
        }
    }
    
    // 최대 전파 경로 카운트의 중간값이 가장 작은 전파 경로가 된다.
    result = int((float32(max) / float32(2)) + 0.5)
    
    // for i:=0; i<len(graph); i++ {
    // // for idx:=0; idx<len(sortedindex); idx++ {    
    //     // i:=sortedindex[idx]
        
    //     // fmt.Fprintln(os.Stderr, "graph[", i, "]:", graph[i])
        
    //     depth := 0
    //     max := 0
    //     bEnd := false
        
    //     // 말단 노드는 탐색 대상에 제외
    //     if len(graph[i].link) == 1 {
    //         continue
    //     }

    //     TraverseGraph(graph, i, i, depth, &max, result, &bEnd)
                
    //     // 가장 작은 전파 경로
    //     if max > 0 && max < result {
    //         result = max
    //         fmt.Fprintln(os.Stderr, "result:", result)
    //     }
    // }
    
    
    fmt.Println(result) // The minimal amount of steps required to completely propagate the advertisement
}

func TraverseGraph(graph map[int]node, base int, idx int, depth int, max *int, curmin int, bEnd *bool) {
    
    if *bEnd {
        return
    }
    
    depth++
    
    
    if curmin < depth {
        *bEnd = true
        // fmt.Fprintln(os.Stderr, "bEnd=true, idx:", idx, "curmin:", curmin, "depth:", depth)
        return
    }
    
    graph[idx] = node {
        link: graph[idx].link,
        propagation_base: base,
    }
    
    

    // 최소의 전파 비용 찾기
    // for i:=0; i<len(graph[idx].link); i++ {
    for i:=range graph[idx].link {
        
        if curmin < depth {
            *bEnd = true
            return
        }
        n := graph[idx].link[i]
        
        if graph[n].propagation_base == base {
            continue
        }
        // fmt.Fprintln(os.Stderr, idx, "->", n, "depth:", depth, " max:", *max, "graph[", idx, "].link:", graph[idx].link)
        
        if *max < depth {
            *max = depth        
        }
        
        TraverseGraph(graph, base, n, depth, max, curmin, bEnd)
    }
}
