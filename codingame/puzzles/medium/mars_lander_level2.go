package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // surfaceN: the number of points used to draw the surface of Mars.
    var surfaceN int
    fmt.Scan(&surfaceN)
    
    fmt.Fprintln(os.Stderr, "surfaceN:", surfaceN)
    
    preX, preY := 0, 0
    flatGroundX1, flatGroundX2 := 0, 0
    flatGroundY := 0
    
    var landlistX []int
    var landlistY []int
    
    for i := 0; i < surfaceN; i++ {
        // landX: X coordinate of a surface point. (0 to 6999)
        // landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
        var landX, landY int
        fmt.Scan(&landX, &landY)
        
        fmt.Fprintln(os.Stderr, "landX:", landX, "landY:", landY)
        
        // 지형 정보 파악
        landlistX = append(landlistX, landX)
        landlistY = append(landlistY, landY)
        
        if i == 0 {
            preX = landX
            preY = landY
            continue
        }
        
        // 평탄 지형 찾기
        if preY == landY {
            flatGroundX1 = preX
            flatGroundX2 = landX
            flatGroundY = landY
        }
        
        preX = landX
        preY = landY
    }
    
    
    fmt.Fprintln(os.Stderr, "landlistX:", landlistX)
    fmt.Fprintln(os.Stderr, "landlistY:", landlistY)
    
    
    // midX := 0
    for {
        // hSpeed: the horizontal speed (in m/s), can be negative.
        // vSpeed: the vertical speed (in m/s), can be negative.
        // fuel: the quantity of remaining fuel in liters.
        // rotate: the rotation angle in degrees (-90 to 90).
        // power: the thrust power (0 to 4).
        var X, Y, hSpeed, vSpeed, fuel, rotate, power int
        fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)
        
        
        // fmt.Fprintln(os.Stderr, X, Y, hSpeed, vSpeed, fuel, rotate, power)
        // fmt.Println("-20 3") // rotate power. rotate is the desired rotation angle. power is the desired thrust power.
        fmt.Fprintln(os.Stderr, "flatGroundX:", flatGroundX1, "~", flatGroundX2, " flatGroundY:", flatGroundY)
        // fmt.Fprintln(os.Stderr, "hSpeed:", hSpeed, "vSpeed:", vSpeed)
        
        // landY := 0
        // // 현재 지형 X 위치의 Y 값 파악
        // for i:=0; i<len(landlistX)-1; i++ {
        //     if X >= landlistX[i] && X < landlistX[i+1] {
        //         landY = landlistY[i]
        //         // fmt.Fprintln(os.Stderr, "i:", i, "landY:", landY)
        //         break
        //     }
        // }
        // fmt.Fprintln(os.Stderr, "vSpeed:", vSpeed, "Y:", Y, "landY:", landY)
                

        // 절대 속도
        hs, vs := 0, 0
        if hSpeed < 0 {
            hs = hSpeed * -1
        } else {
            hs = hSpeed
        }
        if vSpeed < 0 {
            vs = vSpeed * -1
        } else {
            vs = vSpeed
        }
        

        disty := (Y - flatGroundY)
        if disty < 0 { disty *= -1 }

        distx := (flatGroundX2 + flatGroundX1) / 2
        if distx > X {
            distx = distx - X
        } else {
            distx = X - distx
        }
        if distx < 0 { distx *= -1 }
        
        fmt.Fprintln(os.Stderr, "distx:", distx, "disty:", disty)
        
        
        // 현재 위치에서 착륙지점이 오른쪽에 있는 경우
        if X < flatGroundX1 {
            
            rotate = hSpeed - vs
                    
        // 현재 위치에서 착륙지점이 왼쪽에 있는 경우
        } else if X > flatGroundX2 {
            
            rotate = hSpeed + vs
            
            
        // 현재 위치가 착륙지점인 경우
        } else {            
            
            if hSpeed > 0 {
                if hs > 20 {
                    rotate++
                } else {
                    rotate--
                }
            } else {
                if hs > 20 {
                    rotate--                    
                } else {
                    rotate++
                }
            }
            
            
            if hs <= 20 {
                rotate = 0
            }
            
        }
        
        if vSpeed < -20 {
            power++
        } else {
            power--
        }
        
        
        if rotate < -90 { rotate = -90 }
        if rotate > 90 { rotate = 90 }    
        if power < 0 { power = 0}
        if power > 4 { power = 4}
        fmt.Println(rotate, power)        
    }
}
