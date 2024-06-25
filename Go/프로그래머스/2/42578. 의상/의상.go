func solution(clothes [][]string) int {
    clothMap := map[string][]string{}
    for _, row := range clothes {
        clothMap[row[1]] = append(clothMap[row[1]], row[0])
    }
    
    count := 1
    for _, v := range clothMap {
        count *= len(v) + 1     // 입지 않는 경우
    }
    count--                     // 벌거 벗을 순 없음
    
    return count
}
