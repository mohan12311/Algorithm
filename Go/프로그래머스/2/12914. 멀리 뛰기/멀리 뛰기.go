func solution(n int) int64 {
    return calc(n)
}

func calc(n int) int64 {
    n += 1
    
    if n < 2 {
        return int64(n)
    }
    
    results := make([]int, n+1)
    results[1] = 1
    
    for i := 2; i < n + 1; i++ {
        results[i] = (results[i - 2] + results[i - 1]) % 1234567
    }
    
    return int64(results[n])
}

