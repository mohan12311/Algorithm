import (
    "strings"
    "unicode"
)

func solution(s string) string {
    var sb strings.Builder
    
    wasBlanked := true
    
    for _, char := range s {
        if unicode.IsSpace(char) {
            sb.WriteRune(char)
            wasBlanked = true
            continue
        }
        if wasBlanked {
            sb.WriteRune(unicode.ToUpper(char))
        } else {
            sb.WriteRune(unicode.ToLower(char))
        }
        wasBlanked = false
    }
    
    return sb.String()
}