package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // n: the number of cards for player 1
    var n int
    fmt.Scan(&n)
    
    fmt.Fprintln(os.Stderr, "player1 cards:", n)
    
    var deck1 []string
    for i := 0; i < n; i++ {
        // cardp1: the n cards of player 1
        var cardp1 string
        fmt.Scan(&cardp1)
        deck1 = append(deck1, cardp1)
    }
    fmt.Fprintln(os.Stderr, "deck1:", deck1)
    
    // m: the number of cards for player 2
    var m int
    fmt.Scan(&m)
    
    fmt.Fprintln(os.Stderr, "player2 cards:", m)
    
    var deck2 []string
    for i := 0; i < m; i++ {
        // cardp2: the m cards of player 2
        var cardp2 string
        fmt.Scan(&cardp2)
        deck2 = append(deck2, cardp2)
    }
    fmt.Fprintln(os.Stderr, "deck2:", deck2)
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    round := 0
    winner := 0
    
    fmt.Fprintln(os.Stderr, "start game...")
    
    for {
        
        fmt.Fprintln(os.Stderr, "deck1:", deck1)
        fmt.Fprintln(os.Stderr, "deck2:", deck2)
        
        if len(deck1) == 0 {
            winner = 2
            break
        } else if len(deck2) == 0 {
            winner = 1
            break
        }

        // 플레이된 카드를 쌓아놓을 공간
        var playedcardlist1 []string
        var playedcardlist2 []string
        
        // 승자 결정
        winner = Fight(deck1[0], deck2[0])
        playedcardlist1 = append(playedcardlist1, deck1[0])
        playedcardlist2 = append(playedcardlist2, deck2[0])

        // fmt.Fprintln(os.Stderr, "round:", round, "c1:", c1, "n1:", n1)
        // fmt.Fprintln(os.Stderr, "round:", round, "c2:", c2, "n2:", n2)
        
        i := 1
        // 같은 카드라면
        for winner == 0 {
            fmt.Fprintln(os.Stderr, "WAR")
            // return 
            
            
            //  각 플레이어의 다음 3장의 카드를 쌓아두고
            for j:=0;j<3; j++ {
                
                // 다음 3장의 카드를 꺼낼 수 없으면 무승부 처
                if i>=len(deck1) || i>=len(deck2) {
                    fmt.Println("PAT")
                    return
                }
                
                playedcardlist1 = append(playedcardlist1, deck1[i])
                playedcardlist2 = append(playedcardlist2, deck2[i])
                i++
            }
            
            // 4번째 카드로 승자 결정(여기서 같은 카드라면 다시 루프
            winner = Fight(deck1[i], deck2[i])
            
            fmt.Fprintln(os.Stderr, "deck1[",i,"]:", deck1[i], "deck2[",i,"]:", deck2[i])
            
            playedcardlist1 = append(playedcardlist1, deck1[i])
            playedcardlist2 = append(playedcardlist2, deck2[i])
            i++
            if winner != 0 {
                break
            }
        }
        
        round++
        fmt.Fprintln(os.Stderr, "round:", round)
        
        
        if winner == 1 {
            // 플레이1, 2 순으로 추가해야 한다.
            deck1 = append(deck1, playedcardlist1...)
            deck1 = append(deck1, playedcardlist2...)
        } else {
            // 플레이1, 2 순으로 추가해야 한다.
            deck2 = append(deck2, playedcardlist1...)
            deck2 = append(deck2, playedcardlist2...)
        }
        // 플레이한 카드는 제외
        deck1 = append(deck1[i:])
        deck2 = append(deck2[i:])
    }
    
    fmt.Println(winner, round)
}

func Fight(card1, card2 string) (winner int) {
    
    // 앞 부분 숫자만 사용하기
    c1 := card1[:len(card1)-1]
    c2 := card2[:len(card2)-1]
    n1 := CardStringToNum(c1)
    n2 := CardStringToNum(c2)
    
    if n1 > n2 {
        winner = 1
    } else if n1 < n2 {
        winner = 2
    } else {
        winner = 0
    }
    return winner
}

func CardStringToNum(s string) int {
    switch s {
        case "2": return 2
        case "3": return 3
        case "4": return 4
        case "5": return 5
        case "6": return 6
        case "7": return 7
        case "8": return 8
        case "9": return 9
        case "10": return 10
        case "J" : return 11
        case "Q" : return 12
        case "K" : return 13
        case "A" : return 14
    }
    fmt.Fprintln(os.Stderr, "err!")
    return -1
}

