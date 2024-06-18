import (
    "strings"
    "strconv"
)

func solution(s string) string {
    numbers := strings.Split(s, " ")
    
    min, _ := strconv.Atoi(numbers[0])
    max, _ := strconv.Atoi(numbers[0])
    for _, num := range numbers {
        n, _ := strconv.Atoi(num)
        switch {
        case n < min:
            min = n
        case n > max:
            max = n
        }
    }
    
    return strconv.Itoa(min) + " " + strconv.Itoa(max)
}