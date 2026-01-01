package easy

type MyStack struct {
	q1, q2 []int
}

func Constructor() MyStack {
	return MyStack{
		q1: []int{},
		q2: []int{},
	}
}

func (s *MyStack) Push(x int) {
	s.q2 = append(s.q2, x)
	for len(s.q1) > 0 {
		s.q2 = append(s.q2, s.q1[0])
		s.q1 = s.q1[1:]
	}

	s.q1, s.q2 = s.q2, []int{}
}

func (s *MyStack) Pop() int {
	x := s.q1[0]
	s.q1 = s.q1[1:]

	return x
}

func (s *MyStack) Top() int {
	return s.q1[0]
}

func (s *MyStack) Empty() bool {
	return len(s.q1) == 0
}
