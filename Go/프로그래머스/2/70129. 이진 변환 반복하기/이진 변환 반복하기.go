import (
    "strings"
    "strconv"
)

func solution(s string) []int {
    count, zeros := recursion(s, 0, 0)
    return []int{count, zeros}
}

func recursion(binary string, zeros int, processCount int) (int, int){
    if len(binary) == 1 {
        return processCount, zeros
    }
    
    processCount++
    zeros += strings.Count(binary, "0")
    newBinary := toBinary(binary)
    
    return recursion(newBinary, zeros, processCount)
}

func toBinary(num string) string {
    num = strings.Replace(num, "0", "", -1)
    c := len(num)
    return strconv.FormatInt(int64(c), 2)
}