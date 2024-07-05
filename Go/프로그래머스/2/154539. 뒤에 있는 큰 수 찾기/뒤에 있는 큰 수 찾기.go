type Stack struct {
    Value []int
    Index []int
}

const (
    ERR = -1
)

func newStack() *Stack {
    return &Stack{[]int{}, []int{}}
}

func (s *Stack) Push(v, i int) {
    s.Value = append(s.Value, v)
    s.Index = append(s.Index, i)
}

func (s *Stack) Pop() {
    if s.IsEmpty() {
        return
    }
    l := len(s.Value)
    s.Value = s.Value[:l-1]
    s.Index = s.Index[:l-1]
}

func (s *Stack) Peek() (int, int) {
    if s.IsEmpty() {
        return ERR, ERR
    }
    l := len(s.Value)
    return s.Value[l-1], s.Index[l-1]
}

func (s *Stack) IsEmpty() bool {
    return len(s.Value) == 0
}

func bigNumber(a, b int) bool {
    return a < b
}

func solution(numbers []int) []int {
    stack := newStack()
    
    for i, num := range numbers {
        
        for !stack.IsEmpty() {
            prevV, prevI := stack.Peek()
            if !bigNumber(prevV, num) {
                break
            }
            numbers[prevI] = num
            stack.Pop()
        }
        
        stack.Push(num, i)
    }
    
    for !stack.IsEmpty() {
        _, i := stack.Peek()
        numbers[i] = -1
        stack.Pop()
    }
    
    return numbers
    
}