import "strconv"

func solution(n int64) int64 {
    str := strconv.Itoa(int(n))
    b := []byte(str)
    for i := 1; i < len(b); i++ {
        for q := i ; q > 0; q-- {
            if b[q] < b[q-1] {
                break
            }
            b[q], b[q-1] = b[q-1], b[q]
        }
    }
    result, _ := strconv.Atoi(string(b))
    return int64(result)
}