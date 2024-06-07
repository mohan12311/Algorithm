// 주문한 순서대로 아래에서 위로 쌓인다.
// -> 큐 구현문제

// 정해진 순서대로 빵 - 야채 - 고기 - 빵
// 함수로 처리?

// 1 빵
// 2 야채
// 3 고기
// [1 2 3 1]의 연속적인 배열이 있어야 함.

// index=2에서 '1'을 발견!
// index += 1에서 '2'를 발견!
// index += 2에서 '3'를 발견!
// index += 3에서 '3'를 발견!

// founded!++ 후,
// index += 4를 index자리에 집어넣습니다 ...
// 최대 100만까지면 불가능할듯.

// 링크드 리스트로 구현해보자

type LinkedList struct {
	Head  *Node
	Tail  *Node
	Count int
}

type Node struct {
	Next  *Node
	Value int
}

func New() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func NewNode(v int) *Node {
	return &Node{nil, v}
}

func (l *LinkedList) Add(v int) {
	newNode := NewNode(v)
	if l.Tail != nil {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
	if l.Head == nil {
		l.Head = newNode
	}
	l.Count++
}

func (l *LinkedList) Process() bool {
    cur := l.Head
	var prev *Node
    
	for cur != nil && cur.Next != nil && cur.Next.Next != nil && cur.Next.Next.Next != nil {
		if cur.Value == 1 &&
			cur.Next.Value == 2 &&
			cur.Next.Next.Value == 3 &&
			cur.Next.Next.Next.Value == 1 {
			if prev == nil {
				l.Head = cur.Next.Next.Next.Next
			} else {
				prev.Next = cur.Next.Next.Next.Next
			}
			if cur.Next.Next.Next.Next == nil {
				l.Tail = prev
			}
			cur = cur.Next.Next.Next.Next
			l.Count -= 4
			return true
		} else {
			prev = cur
			cur = cur.Next
		}
	}
    return false
}

func solution(ingredient []int) int {
	linkedList := New()

	for _, v := range ingredient {
		linkedList.Add(v)
	}
    
    result := 0
    
    for linkedList.Process() {
        result++
    }

	return result
}
