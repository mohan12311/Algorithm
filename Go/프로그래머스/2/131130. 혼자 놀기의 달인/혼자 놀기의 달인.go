import "sort"

func solution(cards []int) int {
    n := len(cards)
    opened := make([]bool, n)
    var groups []int
    
    for i := 0; i < n; i++ {
        if !opened[i] {
            size := calc(cards, opened, i)
            groups = append(groups, size)
        }
    }
    
    if len(groups) < 2 {
        return 0
    }
    
    sort.Slice(groups, func(i, j int) bool {
        return groups[i] > groups[j]
    })
    
    return groups[0] * groups[1]
}

func calc(cards []int, opened []bool, index int) int {
    count := 0
    cur := index
    
    for !opened[cur] {
        opened[cur] = true
        count++
        cur = cards[cur] - 1
    }
    
    return count
}