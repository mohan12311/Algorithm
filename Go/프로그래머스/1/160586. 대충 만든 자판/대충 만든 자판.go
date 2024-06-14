func solution(keymap []string, targets []string) []int {
    alphabetMap := make(map[rune]int)
    for _, keys := range keymap {
        
        for i, char := range keys {
            if alphabetMap[char] == 0 || alphabetMap[char] > i+1 {
                alphabetMap[char] = i + 1
            }
        }
        
    }
    
    result := make([]int, len(targets))
    
    for i, target := range targets {
        
        count := 0
        err := false
        for _, char := range target {
            v, ok := alphabetMap[char]
            if !ok {
                err = true
                break
            }
            count += v
        }
        
        if err {
            result[i] = -1
        } else {
            result[i] = count
        }
        
    }
    
    return result
}