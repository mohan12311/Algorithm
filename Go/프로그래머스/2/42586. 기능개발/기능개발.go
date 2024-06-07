func solution(progresses []int, speeds []int) []int {
    
    toComplete := make([]int, len(progresses))
    for i := 0; i < len(progresses); i++ {
        remaining := 100 - progresses[i]
        toComplete[i] = (remaining + speeds[i] - 1) / speeds[i]
    }
    
    var result []int
    curMax := toComplete[0]
    count := 1
    
    for i := 1; i < len(toComplete); i++ {
        if toComplete[i] <= curMax {
            count++
        } else {
            result = append(result, count)
            curMax = toComplete[i]
            count = 1
        }
    }
    
    result = append(result, count)
    return result
    
}
