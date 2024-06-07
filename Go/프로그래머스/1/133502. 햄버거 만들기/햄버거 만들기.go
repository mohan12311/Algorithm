func solution(ingredient []int) int {
    stack := []int{}
    count := 0
    
    for _, v := range ingredient {
        stack = append(stack, v)
        
        if n := len(stack); n >= 4 {
            if stack[n-4] == 1 && stack[n-3] == 2 && stack[n-2] == 3 && stack[n-1] == 1 {
                stack = stack[:n-4]
                count++
            }
        }
    } 
    
    return count
}