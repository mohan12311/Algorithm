func solution(s string) int {
    count := 0
    l := len(s)
    
    for i := 0; i < l; i++ {
        spinned := spin(s, i)
        if stackSolution(spinned) {
            count++
        }
    }
    
    return count
}

func spin(s string, i int) string {
    return s[i:] + s[:i]
}

func stackSolution(s string) bool {
    stack := []rune{}
    pairMap := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    
    for _, char := range s {
        
        switch char {
            case '(', '{', '[':
            stack = append(stack, char)
            
            case ')', '}', ']':
            if len(stack) == 0 || stack[len(stack)-1] != pairMap[char] {
                return false
            }
            stack = stack[:len(stack)-1]
            
            default:
            panic("Something Wrong!")
        }
        
    }
    
    return len(stack) == 0
}