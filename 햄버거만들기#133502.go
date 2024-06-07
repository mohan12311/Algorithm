// 링크드 리스트로 구현해보자

/*
* 
* 테스트 1 〉	통과 (0.00ms, 3.78MB)
* 테스트 2 〉	통과 (0.00ms, 4.2MB)
* 테스트 3 〉	통과 (2763.11ms, 15.3MB)
* 테스트 4 〉	실패 (시간 초과)
* 테스트 5 〉	실패 (시간 초과)
* 테스트 6 〉	통과 (6442.41ms, 21.6MB)
* 테스트 7 〉	실패 (시간 초과)
* 테스트 8 〉	통과 (7218.19ms, 22.4MB)
* 테스트 9 〉	통과 (3752.96ms, 17.9MB)
* 테스트 10 〉	통과 (2.72ms, 4.16MB)
* 테스트 11 〉	통과 (2063.65ms, 13.9MB)
* 테스트 12 〉	실패 (시간 초과)
* 테스트 13 〉	통과 (0.00ms, 4.21MB)
* 테스트 14 〉	통과 (0.00ms, 3.54MB)
* 테스트 15 〉	통과 (0.00ms, 4.23MB)
* 테스트 16 〉	통과 (0.00ms, 3.57MB)
* 테스트 17 〉	통과 (0.00ms, 3.55MB)
* 테스트 18 〉	통과 (0.00ms, 4.22MB)
* 
* 효율성 문제로 실패...
* 
*/

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
