package data_structures

type set struct {
	m map[interface{}]int
}

// normalSet
func (s *set) Add(i interface{}) bool {
	s.m[i] = 0

	return true
}

func (s *set) Adds(is ...interface{}) bool {
	for _, i := range is {
		s.m[i] = 0
	}

	return true
}

func (s *set) Delete(i interface{}) bool {
	if _, exist := s.m[i]; !exist {
		return false
	}
	delete(s.m, i)

	return true
}

func (s set) IsExist(i interface{}) bool {
	if _, exist := s.m[i]; !exist {
		return false
	}

	return true
}

func (s set) ToList() IList {
	l := make([]interface{}, len(s.m))
	for k := range s.m {
		l = append(l, k)
	}

	return IList(l)
}
