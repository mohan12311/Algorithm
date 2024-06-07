func solution(n int) int {
    var result int
    for i := 1; i < n+1; i++ {
        if n % i == 0 {
            result += i
        }
    }
    return result
}