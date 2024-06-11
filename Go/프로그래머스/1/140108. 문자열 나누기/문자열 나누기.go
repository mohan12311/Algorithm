func solution(s string) int {
    
    firstCounter, otherCounter := 0, 0
    var firstRune rune
    firstRuneSetted := false
    result := 0
    
    for i, char := range s {
        if !firstRuneSetted {
            firstRune = char
            firstRuneSetted = true
        }
        
        if char == firstRune {
            firstCounter++ 
        } else {
            otherCounter++
        }
        
        if firstCounter == otherCounter {
            result++
            firstCounter, otherCounter = 0, 0
            firstRuneSetted = false
            firstRune = '0'
        } else if i == len(s) - 1 {
            result++
            break
        }
        
    }
    
    return result
}