import "strings"

func solution(X string, Y string) string {
    XCounter := make(map[rune]int)
    YCounter := make(map[rune]int)
    
    for _, v := range X {
        XCounter[v]++
    }
    
    for _, v := range Y {
        YCounter[v]++
    }
    
    allZeros := true
    var sb strings.Builder
    for i := '9'; i >= '0'; i-- {
        xCount := XCounter[i]
        yCount := YCounter[i]
        if xCount == 0 || yCount == 0 {
            continue
        }
        if allZeros && i != '0' {
            allZeros = false
        }
        toUpdate := 0
        if xCount < yCount {
            toUpdate = xCount
        } else {
            toUpdate = yCount
        }
        for toUpdate > 0 {
            sb.WriteRune(i)
            toUpdate--
        }
    }
    
    result := sb.String()
    
    if result == "" {
        return "-1"
    }
    
    if allZeros {
        return "0"
    }
    
    return result
    
    
}

// import "sort"
// import "strings"

// func solution(X string, Y string) string {
//     runeX := []rune(X)
//     runeY := []rune(Y)
    
//     pair := []rune{}
    
//     for _, x := range runeX {
//         for i, y := range runeY{
//             if x == y {
//                 pair = append(pair, x)
//                 runeY = append(runeY[:i], runeY[i+1:]...)
//                 break
//             }
//         }
//     }
    
//     var sb strings.Builder
//     if len(pair) == 0 {
//         return "-1"
//     } else {
//         sort.Slice(pair, func(i, j int) bool {
//             return pair[i] > pair[j]
//         })
        
//         allZeros := true
//         for _, z := range pair {
//             if z != '0' {
//                 allZeros = false
//             }
//             sb.WriteRune(z)
//         }
        
//         if allZeros {
//             return "0"
//         }
        
//         return sb.String()
//     }
// }