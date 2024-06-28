import (
    "sort"
    "strings"
    "strconv"
)

func solution(numbers []int) string {
    strNumbers := make([]string, len(numbers))
    for i, num := range numbers {
       strNumbers[i] = strconv.Itoa(num)
    }
    
    sort.Slice(strNumbers, func (i, j int) bool {
        return strNumbers[i] + strNumbers[j] > strNumbers[j] + strNumbers[i]
    })
    
    res := strings.Join(strNumbers, "")
    
    if res[0] == '0' { return "0"}
    
    return res
}