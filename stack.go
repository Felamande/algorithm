package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	elems  []interface{}
	length int
}

func (s *Stack) Init() *Stack {
	s.length = 1
	return s
}

func (s *Stack) Push(v interface{}) {
	if len(s.elems) == s.Len() {
		s.elems = append(s.elems, v)
	} else {
		s.elems[s.Len()] = v
	}
	s.length++
}

func (s *Stack) Pop() (v interface{}) {

	v = s.elems[s.Len()-1]
	s.length--
	return v
}

func (s *Stack) Peek() (v interface{}) {
	return s.elems[s.Len()-1]
}

func (s *Stack) Len() int {
	return s.length - 1
}

func calc(expr string) (v int64, err error) {
	ops := new(Stack).Init()
	vals := new(Stack).Init()
	v = 0
	for _, s := range expr {
		switch {
		case s == ' ':
			continue
		case s == '(':
			continue
		case s >= '0' && s <= '9':
			vals.Push(int64(s - '0'))
		case s == '+' || s == '-' || s == '*' || s == '/':
			ops.Push(s)
		case s == ')':
			op := ops.Pop().(rune)
			v = vals.Pop().(int64)
			switch op {
			case '+':
				v = vals.Pop().(int64) + v
			case '-':
				v = vals.Pop().(int64) - v
			case '*':
				v = vals.Pop().(int64) * v
			case '/':
				v = vals.Pop().(int64) / v
			}
			vals.Push(v)
		default:
			return 0, errors.New("invalid symbol: " + string(s))
		}
	}
	return v, nil
}

func parenthese(str string) bool {
	s := new(Stack).Init()

	for _, w := range str {
		switch w {
		case '(', '{', '[':
			s.Push(w)
		case ')':
			if s.Len() == 0 {
				return false
			}
			lb := s.Pop().(rune)
			if lb != '(' {
				return false
			}
		case ']':
			if s.Len() == 0 {
				return false
			}
			lb := s.Pop().(rune)
			if lb != '[' {
				return false
			}
		case '}':
			if s.Len() == 0 {
				return false
			}
			lb := s.Pop().(rune)
			if lb != '{' {
				return false
			}
		default:
			return false
		}
	}
	if s.Len() != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(parenthese("[[[}}"))
}
