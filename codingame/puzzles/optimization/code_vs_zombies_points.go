package main

import "fmt"
import "os"

/**
 * Save humans, destroy zombies!
 **/


type human struct {
    id int
    x int
    y int
    minZombieID int
    minZombieDistX int
    minZombieDistY int
}

type zombie struct {
    id int
    x int
    y int
    minShooterDistX int
    minShooterDistY int    
}


func main() {
    for {
        var shooterX, shooterY int
        fmt.Scan(&shooterX, &shooterY)
        
        var humanCount int
        fmt.Scan(&humanCount)
        
        
        var hlist []human
        for i := 0; i < humanCount; i++ {
            var humanId, humanX, humanY int
            fmt.Scan(&humanId, &humanX, &humanY)
            humanobj := human {
                id:humanId,
                x:humanX,
                y:humanY,
            }
            hlist = append(hlist, humanobj)
        }
        
        // fmt.Fprintln(os.Stderr, "hlist:", hlist)
        
        var zombieCount int
        fmt.Scan(&zombieCount)
        
        var zlist []zombie
        for i := 0; i < zombieCount; i++ {
            var zombieId, zombieX, zombieY, zombieXNext, zombieYNext int
            fmt.Scan(&zombieId, &zombieX, &zombieY, &zombieXNext, &zombieYNext)
            
            // fmt.Println(zombieX, zombieY)
            
            zombieobj := zombie {
                id:zombieId,
                x:zombieX,
                y:zombieY,
            }
            zlist = append(zlist, zombieobj)
        }

        ZHMin := 16000 + 9000
        // SZMin := 16000 + 9000
        SHMin := 16000 + 9000
            
        
        // 사람마다 가장 가까운 좀비 정보 기록
        for i:=range hlist {
            ZHDistx, ZHDisty := 0, 0
            ZHMin = 16000 + 9000
            for j:=range zlist {
                if hlist[i].x < zlist[j].x {
                    ZHDistx = zlist[j].x - hlist[i].x
                } else {
                    ZHDistx = hlist[i].x - zlist[j].x
                }
                if hlist[i].y < zlist[j].y {
                    ZHDisty = zlist[j].y - hlist[i].y
                } else {
                    ZHDisty = hlist[i].y - zlist[j].y
                }
                if ZHDistx + ZHDisty < ZHMin {
                    ZHMin = ZHDistx + ZHDisty
                    hlist[i].minZombieID = zlist[j].id
                    hlist[i].minZombieDistX = ZHDistx
                    hlist[i].minZombieDistY = ZHDisty
                }
            }
        }
        
        // // 슈터 와 좀비 거리
        // for i:=range zlist {
        //     SZDistx, SZDisty := 0, 0
        //     if shooterX < zlist[i].x {
        //         SZDistx = zlist[i].x - shooterX
        //     } else {
        //         SZDistx = shooterX - zlist[i].x
        //     }
        //     if shooterY < zlist[i].y {
        //         SZDisty = zlist[i].y - shooterY
        //     } else {
        //         SZDisty = shooterY - zlist[i].y
        //     }
        //     zlist[i].minShooterDistX = SZDistx
        //     zlist[i].minShooterDistX = SZDisty
        // }
        
        
        
        targetX, targetY := 0, 0
        
        // 슈터와 사람 거리
        SHDistx, SHDisty := 0, 0
        for i:=range hlist {
            if hlist[i].x < shooterX {
                SHDistx = shooterX - hlist[i].x
            } else {
                SHDistx = hlist[i].x - shooterX
            }
            if hlist[i].y < shooterY {
                SHDisty = shooterY - hlist[i].y
            } else {
                SHDisty = hlist[i].y - shooterY
            }
            
            fmt.Fprintln(os.Stderr, "SHDistx:", SHDistx, "SHDisty:", SHDisty)
            fmt.Fprintln(os.Stderr, "hlist[",i,"].minZombieDistX:", hlist[i].minZombieDistX, "hlist[",i,"].minZombieDistY", hlist[i].minZombieDistY)
            // 슈터와 사람 거리, 사람과 종비거리를 고려
            if (SHDistx + SHDisty) - (hlist[i].minZombieDistX + hlist[i].minZombieDistY) < SHMin {
                SHMin = (SHDistx + SHDisty) - (hlist[i].minZombieDistX + hlist[i].minZombieDistY)
                targetX = hlist[i].x
                targetY = hlist[i].y
            }
        }
        
        fmt.Println(targetX, targetY)
    }
}
