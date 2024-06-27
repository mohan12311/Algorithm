func solution(n int, left int64, right int64) []int {
    leftX := left % int64(n) 
    leftY := left / int64(n) 
    rightX := right % int64(n) 
    rightY := right / int64(n)
    
    arr := []int{}
    for v := leftY; v <= rightY; v++ {
        for h := int64(0); h < int64(n); h++ {
            if v == leftY && h < leftX {
                continue
            }
            if v == rightY && h > rightX {
                continue
            }
            arr = append(arr, max(v + 1, h + 1))
        }
    }
    
    return arr
}

func max(a, b int64) int {
    if a > b {
        return int(a)
    }
    return int(b)
}