func drag(ldx, luy, rdx, rdy int) int {
    x := ldx - rdx
    y := luy - rdy
    if x < 0 {
        x *= -1
    }
    if y < 0 {
        y *= -1
    }
    return x + y
}

func count(array []bool) (int, int) {
    var first, last int
    for i, v := range array {
        if v {
            first = i
            break
        }
    }
    
    for i := len(array) - 1; i >= 0; i-- {
        if array[i] {
            last = i + 1
            break
        }
    }
    
    return first, last
}

func solution(wallpaper []string) []int {
    var top, bottom, left, right int
    y := make([]bool, len(wallpaper))
    x := make([]bool, len(wallpaper[0]))
    for i, line := range wallpaper {
        for j, char := range line {
            if char == '#' {
                y[i] = true
                x[j] = true
            }
        }
    }
    
    top, bottom = count(y)
    left, right = count(x)
    
    
    
    return []int{top, left, bottom, right}
}