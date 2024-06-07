func countingSort(arr [] int, max int) []int {
    countingArray := make([]int, max+1)
    
    for _, v := range arr {
        countingArray[v]++
    }
    
    var result []int
    
    for i := len(countingArray) - 1; i >= 0; i--{
        count := countingArray[i]
        for count > 0 {
            result = append(result, i)
            count--
        }
    }
    
    return result
}

func solution(k int, m int, score []int) int {
    sortedScore := countingSort(score, k)
    result := 0
    
    for i := m; i <= len(sortedScore); i += m {
        min := sortedScore[i-1]
        result += min * m
    }
    
    return result
}