package main

import "fmt"
import "os"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

// 노드 자료 구조
type nodeinfo2 struct {
    num int
    node []nodeinfo2
}

// AddChildNode 자식 노드 추가
func AddChildNode(parent *nodeinfo2, num int) {
	parent.node = append(parent.node, nodeinfo2{num, nil})
// 	fmt.Fprintf(os.Stderr, "add parent.node[%v].num:%v\n", len(parent.node)-1, parent.node[len(parent.node)-1].num)
}

// TraverseNode 노드 탐색
func TraverseNode(node *nodeinfo2, cnt *int, depth int) {
	for i := 0; i < len(node.node); i++ {
		(*cnt)++
		for j := 0; j < depth; j++ {
			fmt.Fprintf(os.Stderr, "    ")
		}
		fmt.Fprintf(os.Stderr, "node[%v].num:%v\n", i, node.node[i].num)
		TraverseNode(&node.node[i], cnt, depth+1)
	}
}



func main() {
    var N int
    fmt.Scan(&N)
    
    var rootnode nodeinfo2
    rootnode.num = -1
    rootnode.node = nil
    

    for i := 0; i < N; i++ {
        var telephone string
        fmt.Scan(&telephone)
        
        fmt.Fprintln(os.Stderr, telephone)


        mynode := &rootnode
        
        // 입력된 폰번호를 링크 노드로 구성
        for i:=0; i<len(telephone);i++ {
            

            n, _ := strconv.Atoi(string(telephone[i]))

            bAdd := true
            for j:=0; j<len(mynode.node); j++ {
                
                // 이미 있는 숫자는 추가하지 않는다
                if mynode.node[j].num == n {
                    bAdd = false
                    break
                }
            }
            
            // 없으면 자식 노드로 추가
            if bAdd {
                AddChildNode(mynode, n)
            }
            
            
            // 현재 노드위치 파악
            for j:=0; j<len(mynode.node); j++ {
                // fmt.Fprintf(os.Stderr, "mynode.node[%v].num:%v == %v\n", j, mynode.node[j].num, n)
                if mynode.node[j].num == n {
                    mynode = &mynode.node[j]
                    break
                }
            }
            
        }
    }
    
    // 구성된 노드들 개수 카운트
	nodecnt := 0
	TraverseNode(&rootnode, &nodecnt, 0)
	fmt.Println(nodecnt)
	
	
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    // fmt.Println("number") // The number of elements (referencing a number) stored in the structure.
}
