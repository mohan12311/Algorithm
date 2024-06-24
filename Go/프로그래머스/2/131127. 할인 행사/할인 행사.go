type Bucket struct {
    Items map[string]int
}

func createBucket(w []string, n []int) *Bucket {
    b := Bucket{ Items: map[string]int{} }
    for i := 0; i < len(w); i++ {
        b.Items[w[i]] = n[i]
    }
    
    return &b
}

func (b *Bucket) Check() bool {
    for _, v := range b.Items {
        if v != 0 {
            return false
        }
    }
    return true
}

func (b *Bucket) Init(discount []string) {
    for i := 0; i < 10; i++ {
        object := discount[i]
        b.Items[object]--
    } 
}

func solution(want []string, number []int, discount []string) int {
    bucket := createBucket(want, number)
    bucket.Init(discount)
    count := 0
    if bucket.Check() {
        count++
    }
    
    for i := 10; i < len(discount); i++ {
        head := discount[i - 10]
        tail := discount[i]
        bucket.Items[head]++
        bucket.Items[tail]--
        if bucket.Check() {
            count++
        }
    }
    
    return count
}