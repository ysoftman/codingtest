package main

import "fmt"
import "os"
import "strconv"
/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

// 노드의 링크 정보
type nodeinfo struct {
    name string
    links []int
}

func main() {
    // N: the total number of nodes in the level, including the gateways
    // L: the number of links
    // E: the number of exit gateways
    var N, L, E int
    fmt.Scan(&N, &L, &E)
    
    
    // 노드에 대한 링크정보를 저장할 데이터구조
    var node map[int](nodeinfo)
    node = make(map[int](nodeinfo))
    
    for i := 0; i < L; i++ {
        // N1: N1 and N2 defines a link between these nodes
        var N1, N2 int
        fmt.Scan(&N1, &N2)
        // fmt.Fprintln(os.Stderr, N1, N2)
        // N1 노드에 N2 링크 추가
        node[N1] = nodeinfo {
            "node",
            append(node[N1].links, N2),
        }
        // N2 노드에 N1 링크 추가
        node[N2] = nodeinfo {
            "node",
            append(node[N2].links, N1),
        }
    }
    for i := 0; i < E; i++ {
        // EI: the index of a gateway node
        var EI int
        fmt.Scan(&EI)
        // fmt.Fprintln(os.Stderr, EI)
        // gateway 명시
        node[EI] = nodeinfo {
            "gateway",
            node[EI].links,
        }
    }
    
    fmt.Fprintln(os.Stderr, node)
    
    for {
        // SI: The index of the node on which the Skynet agent is positioned this turn
        var SI int
        fmt.Scan(&SI)
        
        departlink := -1
        destlink := -1
        ln := -1
        bstop := false
        for idx := range node {    
            
            // gateway 노드를 기준으로 바이러스위치와 가까운 링크 찾기            
            if node[idx].name == "gateway" {
                
                for ln = range node[idx].links {
                    // fmt.Fprintf(os.Stderr, "ln %v\n", ln)
                    // fmt.Fprintf(os.Stderr, "node[%v].links %v\n", idx, node[idx].links)
                    if node[idx].links[ln] < 0 {
                        // fmt.Fprintf(os.Stderr, "continue\n")
                        continue
                    }
                    departlink = idx
                    destlink = node[idx].links[ln]
                    // fmt.Fprintf(os.Stderr, "node[%v].links[%v] %v\n", idx, ln, node[idx].links[ln])
                    // 바이러스가 바로 이웃하고 있다면 더이상 탐색하지 현재 링크로 끝내기
                    if node[idx].links[ln] == SI {
                        bstop = true
                        break
                    }
                }
                
            } else {
                continue
            }
            
            if bstop {
                break
            }
        }
        
        fmt.Println(strconv.Itoa(departlink) + " " + strconv.Itoa(destlink))

        // 제거된 링크는 노드에서 제외(-1로 표시)
        node[departlink].links[ln] = -1
        for ln = range node[destlink].links {
            if node[destlink].links[ln] == departlink {
                break
            }
        }
        node[destlink].links[ln] = -1
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        // fmt.Println("0 1") // Example: 0 1 are the indices of the nodes you wish to sever the link between
    }
}