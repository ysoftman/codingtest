package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // nbFloors: number of floors
    // width: width of the area
    // nbRounds: maximum number of rounds
    // exitFloor: floor on which the exit is found
    // exitPos: position of the exit on its floor
    // nbTotalClones: number of generated clones
    // nbAdditionalElevators: ignore (always zero)
    // nbElevators: number of elevators
    var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
    fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)
    
    fmt.Fprintf(os.Stderr, 
        "nbFloors:%v width:%v nbRounds:%v exitFloor:%v exitPos:%v nbTotalClones:%v nbAdditionalElevators:%v nbElevators:%v\n",
        nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators)
        
    
    
    // 층별 엘레베이터 위치 정보
    var mapElevatorInfo map[int]int
    mapElevatorInfo = make(map[int]int)
    
    for i := 0; i < nbElevators; i++ {
        // elevatorFloor: floor on which this elevator is found
        // elevatorPos: position of the elevator on its floor
        var elevatorFloor, elevatorPos int
        fmt.Scan(&elevatorFloor, &elevatorPos)
        fmt.Fprintf(os.Stderr, "elevatorFloor:%v, &elevatorPos:%v\n", elevatorFloor, elevatorPos)
        
        // 층별 엘레베이터 정보 파악
        mapElevatorInfo[elevatorFloor] = elevatorPos
        fmt.Fprintf(os.Stderr, "mapElevatorInfo[%v] = %v\n", elevatorFloor, mapElevatorInfo[elevatorFloor])
    }
    
    
    
    for {
        // cloneFloor: floor of the leading clone
        // clonePos: position of the leading clone on its floor
        // direction: direction of the leading clone: LEFT or RIGHT
        var cloneFloor, clonePos int
        var direction string
        fmt.Scan(&cloneFloor, &clonePos, &direction)
        fmt.Fprintf(os.Stderr, "cloneFloor:%v, clonePos:%v, direction:%v\n", cloneFloor, clonePos, direction)
        
        nbRounds--
        fmt.Fprintln(os.Stderr, "Remain Rounds:", nbRounds)
        
        // 앞에가는 클로이 위험을 위치에서 블럭해야 한다.
        // 출구까지 갈 수 있도록 block 으로 경로 만든다.
        // 앞장선 클론이 위험지대를 파악하면 그자리에서 블럭하여 뒤따르는 클론의 방향을 반대로 안내한다.
        result := "WAIT"
        
        
        // 기본적으로 width 처음과 시작은 위험위치
        if clonePos == width-1 {
            result = "BLOCK"
        } else if clonePos  == 0 {
            result = "BLOCK"
        }
        
        
        // 현재 층에 출구가 없다면 위층으로 올라가야 한다.
        if cloneFloor != exitFloor {        
            // 뒤따라오는 클론들을 현재 층에 엘레베이터 위치로 유도해야 한다.
            if direction == "LEFT" && mapElevatorInfo[cloneFloor] - 1 ==  clonePos {
                result = "BLOCK"
                
            } else if direction == "RIGHT" && mapElevatorInfo[cloneFloor] + 1 ==  clonePos {
                result = "BLOCK"
            }
            
            if direction == "LEFT" && mapElevatorInfo[cloneFloor] > clonePos {
                result = "BLOCK"
            } else if direction == "RIGHT" && mapElevatorInfo[cloneFloor] < clonePos {
                result = "BLOCK"
            }
        // 현재 층에 출구가 있따면
        } else {
            if direction == "LEFT" && exitPos > clonePos {
                result = "BLOCK"
            } else if direction == "RIGHT" && exitPos < clonePos {
                result = "BLOCK"
            }
        }
            
        
        
            
        fmt.Println(result)
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        // fmt.Println("WAIT") // action: WAIT or BLOCK
    }
}