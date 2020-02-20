package data

type elem struct {
	X          int
	Prev, Next *elem
}

// implemented with double linked cycle list and hash map
// Constant time add, remove, check
// Iterable, and able to insert and delete easily while iterating
type IntSet struct {
	hm         map[int]*elem
	head, curr *elem
}

func (s *IntSet) Init(ary []int) {
	s.hm = make(map[int]*elem)
	if ary == nil {
		return
	}
	for _, n := range ary {
		s.Add(n)
	}
}

func (s *IntSet) Add(n int) {
	if _, ok := s.hm[n]; !ok {
		nn := &elem{
			X:    n,
			Prev: nil,
			Next: nil,
		}
		s.hm[n] = nn
		if s.head == nil {
			s.head = nn
			nn.Prev = nn
			nn.Next = nn
		} else {
			s.head.Prev.Next = nn
			nn.Prev = s.head.Prev
			nn.Next = s.head
			s.head.Prev = nn
		}
	}
}

func (s *IntSet) Contains(n int) bool {
	_, ok := s.hm[n]
	return ok
}

func (s *IntSet) Remove(n int) {
	e, ok := s.hm[n]
	if !ok {
		return
	}
	if e.Next == e {
		s.head = nil
		s.curr = nil
	} else {
		if s.head == e {
			s.head = e.Next
		}
		e.Next.Prev = e.Prev
		e.Prev.Next = e.Next
	}
	delete(s.hm, n)
}

func (s *IntSet) Get() int {
	if s.curr == nil {
		s.curr = s.head
	}
	return s.curr.X
}

// return whether next s.Get() call is still within the same iteration
func (s *IntSet) Next() bool {
	if s.curr == nil {
		s.curr = s.head
	}
	s.curr = s.curr.Next
	return s.curr != s.head
}

// insert before s.curr
func (s *IntSet) Insert(n int) {
	if _, ok := s.hm[n]; !ok && s.curr != nil {
		nn := &elem{
			X:    n,
			Prev: s.curr.Prev,
			Next: s.curr,
		}
		s.hm[n] = nn
		s.curr.Prev.Next = nn
		s.curr.Prev = nn
	}
}

// delete current element under s.curr and move it to next elem
func (s *IntSet) Delete() {
	if s.curr == nil {
		return
	}
	v := s.curr.X
	s.curr = s.curr.Next
	s.Remove(v)
}

func (s *IntSet) Size() int {
	return len(s.hm)
}
