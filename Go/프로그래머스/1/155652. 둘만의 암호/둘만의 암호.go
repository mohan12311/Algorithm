import "strings"

func solution(s string, skip string, index int) string {
    
    //var result string
    var result strings.Builder
    //runeSlice := []rune(skip)
    skipMap := make(map[rune]bool)
    
    for _,char := range skip {
        skipMap[char] = true
    }
    
    for _, char := range s {
        count := 0
        processed := char
        
        for count < index {
            processed++
            
            if processed > 'z' {
                processed -= 26
            }
            
            if skipMap[processed] {
                continue
            }
            count++
        }
        
        result.WriteRune(processed)
    }
    
    return result.String()
}

// func check(target rune, checkSlice []rune) bool {
//     for _, v := range checkSlice {
//         if v == target { return true }
//     }
//     return false
// }