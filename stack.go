package containers

type stack struct {
	l      []interface{}
	topPtr int
	maxLen int
}

// normalStack
func (s *stack) Push(i interface{}) bool {
	if s.topPtr == s.maxLen {
		return false
	}

	s.l[s.topPtr] = i
	s.topPtr++

	return true
}

func (s *stack) Pop() (bool, interface{}) {
	if s.topPtr == 0 {
		return false, nil
	}

	var res interface{}
	res, s.l[s.topPtr-1] = s.l[s.topPtr-1], 0
	s.topPtr--

	return true, res
}

func (s stack) Len() int {
	return s.topPtr
}

func (s stack) ToList() IList {
	return s.l
}

func (s stack) Top() interface{} {
	if s.topPtr == 0 {
		return nil
	}
	return s.l[s.topPtr-1]
}
