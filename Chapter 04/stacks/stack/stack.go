package stack

type Stack struct {
	stack        []int
	stackPointer int
}

func (s *Stack) Push(v int) {
	s.stack = append(s.stack, v)
	s.stackPointer = len(s.stack) - 1
}

func (s *Stack) Pop() int {
	if s.stackPointer == -1 {
		return -1
	}
	element := s.stack[s.stackPointer]
	s.stack = s.stack[:s.stackPointer]
	s.stackPointer--
	return element
}
