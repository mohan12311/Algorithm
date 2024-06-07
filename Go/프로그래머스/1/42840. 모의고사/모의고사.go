func solution(answers []int) []int {
    
	pattern1 := []int{1, 2, 3, 4, 5}
	pattern2 := []int{2, 1, 2, 3, 2, 4, 2, 5}
	pattern3 := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
    
    scores := []int{0, 0, 0}
    
    for i, answer := range answers {
        if answer == pattern1[i%len(pattern1)] {
            scores[0]++
        }        
        if answer == pattern2[i%len(pattern2)] {
            scores[1]++
        }        
        if answer == pattern3[i%len(pattern3)] {
            scores[2]++
        }
    }
    
    maxScore := max(scores)
    var result []int
    
    for i, score := range scores {
        if score == maxScore {
            result = append(result, i+1)
        }
    }
    
    return result
}

func max(arr []int) int {
    maxVal := arr[0]
    for _, value := range arr {
        if value > maxVal {
            maxVal = value
        }
    }
    return maxVal
}