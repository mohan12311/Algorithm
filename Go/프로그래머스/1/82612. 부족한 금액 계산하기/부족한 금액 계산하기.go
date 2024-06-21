func solution(price int, money int, count int) int64 {
    totalPrice := 0
    for i := 1; i <= count; i++ {
        totalPrice += price * i
    }
    
    if money >= totalPrice {
        return 0
    } else {
        return int64(totalPrice - money)
    }
}