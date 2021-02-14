package jinrou

import "errors"

type Stack []*Player

func (s *Stack) Push(p *Player) {
	*s = append(*s, p)
}

func (s *Stack) Pop() (*Player, error) {
	if s.Empty() {
		return nil, errors.New("stack size is not enough")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s *Stack) Duplicate() error {
	if s.Empty() {
		return errors.New("stack size is not enough")
	}
	*s = append(*s, (*s)[len(*s)-1])
	return nil
}

func (s *Stack) Top() (*Player, error) {
	if s.Empty() {
		return nil, errors.New("stack size is not enough")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) Exchange() error {
	if len(*s) < 2 {
		return errors.New("stack size is not enough")
	}
	(*s)[len(*s)-1], (*s)[len(*s)-2] = (*s)[len(*s)-2], (*s)[len(*s)-1]
	return nil
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
