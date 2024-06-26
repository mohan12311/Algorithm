func solution(n int, a int, b int) int {
    answer := 0

    for a != b {
        a = (a + 1) / 2
        b = (b + 1) / 2
        answer++
    }

    return answer
}