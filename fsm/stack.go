package fsm

import (
	"fmt"
)

type Stack struct {
	elems  map[int]interface{}
	length int
	name   string
}

func (s *Stack) Init(name string) *Stack {
	s.length = 1
	s.name = name
	s.elems = make(map[int]interface{})
	return s
}

func (s *Stack) Push(v interface{}) {
	// if len(s.elems) == s.Len() {
	// s.elems = append(s.elems, v)
	// } else {
	s.elems[s.Len()] = v
	// }
	s.length++
	// s.Print()
}

func (s *Stack) Pop() (v interface{}, err error) {
	if s.Len() == 0 {
		return nil, EmptyError{}
	}

	v = s.elems[s.Len()-1]
	delete(s.elems, s.Len()-1)
	s.length--
	// s.Print()
	return v, nil
}

func (s *Stack) Peek() (v interface{}, err error) {
	if s.Len() == 0 {
		return nil, EmptyError{}
	}
	v = s.elems[s.Len()-1]
	return v, nil
}

func (s *Stack) Len() int {
	return s.length - 1
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Print() {
	print(s.name + ":")
	for i := 0; i < s.Len(); i++ {
		fmt.Printf("%v ", s.elems[i])
	}
	println()
}
