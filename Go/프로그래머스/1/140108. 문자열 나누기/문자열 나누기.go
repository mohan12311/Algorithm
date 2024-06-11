func solution(s string) int {
    
    firstCounter, otherCounter := 0, 0
    var firstRune rune
    firstRuneSetted := false
    result := 0
    
    for _, char := range s {
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
        }
        
    }
    
    if firstCounter != 0 {
        result++
    }
    
    return result
}