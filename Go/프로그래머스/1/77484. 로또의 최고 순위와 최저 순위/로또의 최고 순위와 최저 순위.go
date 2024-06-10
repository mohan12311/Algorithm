func solution(lottos []int, win_nums []int) []int {
    
    winnerMap := map[int]int{
        0: 6,
        1: 6,
        2: 5,
        3: 4,
        4: 3,
        5: 2,
        6: 1,
    }
    
    zeroCount := 0
    for _, zero := range lottos {
        if zero == 0 {
            zeroCount++
        }
    }
    
    count := 0
    for _, num := range lottos {
        if num == 0 {
            continue
        }
        for _, win := range win_nums {
            if num == win {
                count++
                break
            }
        }
    }
    
    return []int{winnerMap[count+zeroCount], winnerMap[count]}
}