package day5

type Stack[A any] []A

func (s *Stack[A]) Push(v A) {
	*s = append(*s, v)
}

func (s *Stack[A]) Pop() A {
	ds := *s
	res := ds[len(ds)-1]
	*s = ds[:len(ds)-1]
	return res
}

func (s *Stack[A]) Peek() A {
	res := (*s)[len(*s)-1]
	return res
}

func (s *Stack[A]) PushMany(values []A) {
	for _, v := range values {
		s.Push(v)
	}
}

func (s *Stack[A]) PopMany(n int) []A {
	buff := make([]A, 0, n)
	for i := 0; i < n; i++ {
		buff = append(buff, s.Pop())
	}

	return buff
}
