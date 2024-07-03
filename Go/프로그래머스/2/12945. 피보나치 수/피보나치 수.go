func solution(n int) int {
    return altFibo(n) % 1234567
}

func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    a, b := 0, 1
    
    for i := 2; i <= n; i++ {
        a, b = b, a + b
    }
    
    return b
}

func altFibo(n int) int {
    if n < 2 {
        return n
    }
    
    fiboResults := make([]int, n+1)
    fiboResults[1] = 1
    
    for i := 2; i < n + 1; i++ {
        fiboResults[i] = (fiboResults[i-2] + fiboResults[i-1]) % 1234567
    }
    
    return fiboResults[n]
}