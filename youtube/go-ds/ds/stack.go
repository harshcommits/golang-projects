package ds

type Stack struct {
	items []int
}

// Push adds items to the top of the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop removes items from the top of the stack
func (s *Stack) Pop() int {
	last_value := len(s.items) - 1
	removed_item := s.items[last_value]
	s.items = s.items[:last_value]
	return removed_item
}
