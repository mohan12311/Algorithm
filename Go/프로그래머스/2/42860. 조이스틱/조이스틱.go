func solution(name string) int {
    n := len(name)
    count := 0
    move := n - 1 // 오른쪽으로만 이동했을 때
    
    for i, char := range name {
        if char > 'A' {
            up := int(char - 'A')
            down := int('Z' - char + 1)
            count += min(up, down)
        }
        
        next := i + 1 
        for next < len(name) && name[next] == 'A' { // 다음 목표 위치를 탐색
            next++
        }
        
        // i  : 현재 위치까지 오는 이동거리
        // n - next : 원점에서 목표 위치까지의 이동 거리
        // min(i, n - next) : 원점으로 돌아가기 vs 진행방향으로 가기
        // 더 효율적인 이동방향으로 업데이트 합니다.
        move = min(move, i + n - next + min(i, n - next))
    }
    
    return count + move
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}