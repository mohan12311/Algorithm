func solution(arr []int) []int {

    if len(arr) <= 1 {
        return []int{-1}
    }
    
    minIndex := 0
    
    for i, num := range arr {
        if num < arr[minIndex] {
            minIndex = i
        }
    }
    
    arr = append(arr[:minIndex], arr[minIndex+1:]...)
    
    return arr
}