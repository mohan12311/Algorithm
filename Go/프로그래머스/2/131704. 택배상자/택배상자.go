func solution(order []int) int {
    stack := []int{}
    loaded := 0
    currentBox := 1
    
    for _, orderBox := range order {
        for currentBox < orderBox {
            stack = append(stack, currentBox)
            currentBox++
        }
        
        if currentBox == orderBox {
            loaded++
            currentBox++
        } else if check(stack, orderBox) {
            stack = stack[:len(stack)-1]
            loaded++
        } else {
            break
        }
    }
    
    return loaded
}

func check(stack []int, target int) bool {
    l := len(stack)
    if l == 0 {
        return false
    }
    if stack[l - 1] != target {
        return false
    }
    return true
}

