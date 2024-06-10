import "container/heap"


// 적이 많은 라운드에서 무적을 써야 최고득점이 가능함
// 최대힙으로 구현
type MaxHeap []int

// heap구조체 인터페이스들

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 최대힙이므로 i < j 가 아닌 i > j
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// 인터페이스 구현 끝
 
func solution(n int, k int, enemy []int) int {
    h := &MaxHeap{}
    heap.Init(h)
    
    round := 0    
    for _, v := range enemy {
        // 일단 무조건 힙에 집어넣기
        heap.Push(h, v)
        n -= v
        if n < 0 {                  // 병사만으로 싸움 불가
            if k > 0 {              // 무적권을 사용할 수 있다면?
                k--
                n += heap.Pop(h).(int)  // 가장 어려운 라운드에서 소모한 병사값을 다시 불러옴
            } else {
                break
            }
        }
        round++
    }
    
    return round
}