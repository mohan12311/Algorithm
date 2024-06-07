func solution(n int, lost []int, reserve []int) int {

    losted := make([]bool, n)
    reserved := make([]bool, n)

    for _, v := range lost {
        losted[v-1] = true
    }

    for _, v := range reserve {
        reserved[v-1] = true
    }

    for i := 0; i < n; i++ {
        if losted[i] && reserved[i] {
            losted[i] = false
            reserved[i] = false
        }
    }

    for i := 0; i < n; i++ {
        if reserved[i] {
            if i > 0 && losted[i-1] {
                losted[i-1] = false
            } else if i < n-1 && losted[i+1] {
                losted[i+1] = false
            }
        }
    }

    for _, v := range losted {
        if v {
            n--
        }
    }

    return n

}