import "sort"

func solution(k int, tangerine []int) int {
    quantityMap := map[int]int{}
    for _, v := range tangerine {
        quantityMap[v]++
    }
    
    tSlice := make([]int, 0, len(tangerine))
    
    for _, v := range quantityMap {
        tSlice = append(tSlice, v)
    }
    
    sort.Slice(tSlice, func(i, j int) bool {
        return tSlice[i] > tSlice[j]
    })
    
    count := 0
    
    for _, v := range tSlice {
        count++
        k -= v
        if k <= 0 {
            break
        }
    }
    
    return count
}