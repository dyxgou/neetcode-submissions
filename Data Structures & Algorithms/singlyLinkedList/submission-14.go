import (
	"iter"
)

type Node struct {
	val  int
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

type LinkedList struct {
	head *Node
	len  int
}

func NewLinkedList() *LinkedList {
	return new(LinkedList)
}

func (l *LinkedList) InsertHead(val int) {
	n := NewNode(val)
	l.len++

	if l.head == nil {
		l.head = n
		return
	}

	n.next = l.head
	l.head = n
}

func (l *LinkedList) InsertTail(val int) {
	if l.len == 0 {
		l.head = NewNode(val)
		l.len++
		return
	}

	for n := l.head; n != nil; n = n.Next() {
		if n.next == nil {
			n.next = NewNode(val)
			l.len++
			return
		}
	}
}

func (l *LinkedList) Get(i int) int {
	var idx int

	for n := range l.Iter() {
		if idx == i {
			return n.val
		}

		idx++
	}

	return -1
}
func (l *LinkedList) Remove(i int) bool {
	if l.len == 0 {
		return false
	}

	if i == 0 {
		if l.len == 1 {
			l.head = nil
		} else {
			l.head = l.head.next
		}

		l.len--
		return true
	}


	if i > l.len-1 {
		return false
	}

	next, stop := iter.Pull(l.Iter())
	defer stop()

	for range i - 1 {
		next()
	}

	prev, ok := next()
	if !ok {
		return false
	}

	prev.next = prev.next.next
	l.len--

	return true
}


func (l *LinkedList) GetValues() []int {
	buf := make([]int, 0, 10)

	for n := range l.Iter() {
		buf = append(buf, n.val)
	}

	return buf
}

func (l *LinkedList) Iter() iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		for n := l.head; n != nil; n = n.Next() {
			if !yield(n) {
				return
			}
		}
	}
}