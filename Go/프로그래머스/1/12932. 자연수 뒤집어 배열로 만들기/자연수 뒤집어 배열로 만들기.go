func solution(n int64) []int {
    var result []int
    for n > 0 {
        result = append(result, int(n%10))
        n /= 10
    }
    return result
}