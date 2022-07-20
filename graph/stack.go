package graph

type Stack struct {
	stack [][]*GraphNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(d []*GraphNode) {
	s.stack = append(s.stack, d)
}

func (s *Stack) Pop() []*GraphNode {
	if s.Size() <= 0 {
		return nil
	}

	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *Stack) Size() int {
	return len(s.stack)
}
