import (
    "strings"
    "strconv"
    "math"
)

type stack struct {
    value []string
    count int
}

func newStack() *stack {
    return &stack{[]string{}, 0}
}

func (s *stack) Push(char rune) {
    s.value = append(s.value, string(char))
}

func (s *stack) Pop() {
    if s.IsEmpty() {
        return
    }
    number, _ := strconv.Atoi(strings.Join(s.value, ""))
    s.value = []string{}
    if isPrime(number) {
        s.count++
    }
}

func (s *stack) Len() int {
    return len(s.value)
}

func (s *stack) IsEmpty() bool {
    return len(s.value) == 0
}

func (s *stack) Process(char rune) {
    if char == '0' && !s.IsEmpty() {
        s.Pop()
    } else {
        s.Push(char)
    }
} 

func solution(n int, k int) int {
    s := converter(n, k)
    stack := newStack()
    
    for _, char := range s {
        stack.Process(char)
    }
    
    if !stack.IsEmpty() {
        stack.Pop()
    }
    
    return stack.count
}

func isPrime(num int) bool {
    
    if num < 2 {
        return false
    }
    
    sqrt := int(math.Sqrt(float64(num)))
    
    for i := 2; i <= sqrt; i++ {
        if num % i == 0 { return false }
    }
    
    return true
}

func converter(n, k int) string {
    res := []string{}
    
	for n > 0 {
		temp := n % k
		res = append([]string{strconv.Itoa(temp)}, res...)
		n /= k
	}

	return strings.Join(res, "")
}