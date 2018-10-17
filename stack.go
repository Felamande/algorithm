package main

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

func calc(expr string) (float64, error) {
	ops := new(Stack).Init("ops")
	val := new(Stack).Init("val")

outer:
	for _, s := range expr {
		switch {
		case isDigit(s):
			val.Push(float64(s - '0'))
		case s == '(':
			ops.Push(s)
		case s == ')':
			for {
				opr_, err := ops.Pop()
				if _, empty := err.(EmptyError); empty {
					return 0, ExprError{')'}
				}
				opr := opr_.(rune)
				if opr == '(' {
					break
				}

				//opr!='(' but empty
				if _, empty := err.(EmptyError); empty {
					return 0, ExprError{')'}
				}

				if !isOpr(opr) {
					return 0, ExprError{opr}
				}

				//opr is an operator
				v1, err1 := val.Pop()
				v2, err2 := val.Pop()
				v, err := OprCalc(v1, err1, v2, err2, opr) // v2 <opr> v1
				if err != nil {
					return 0, err
				}
				val.Push(v)
			}
			//40( 41) 42* 43+ 45- 47/
		case isOpr(s):
			oprIn := s
			for {
				oprTop_, err := ops.Peek()

				if _, empty := err.(EmptyError); empty {
					ops.Push(oprIn)
					continue outer
				}
				oprTop := oprTop_.(rune)

				if !isOpr(oprTop) {
					ops.Push(oprIn)
					continue outer
				}

				prior := isPrior(oprIn, oprTop)
				if prior { //连续出栈 直到oprIn > oprTop
					ops.Push(oprIn)
					continue outer
				}

				//
				ops.Pop()

				// v2 <oprTop> v1
				v1, err1 := val.Pop()
				v2, err2 := val.Pop()
				v, err := OprCalc(v1, err1, v2, err2, oprTop)
				if err != nil {
					return 0, err
				}
				val.Push(v)
			}
			ops.Push(oprIn)
		}
	}

	for !ops.Empty() {
		opr_, _ := ops.Pop()
		opr := opr_.(rune)

		v1, err1 := val.Pop()
		v2, err2 := val.Pop()
		v, err := OprCalc(v1, err1, v2, err2, opr) // v2 <opr> v1
		if err != nil {
			return 0, err
		}
		val.Push(v)
	}

	re, err := val.Pop()
	if err != nil {
		return 0, err
	}
	return re.(float64), nil
}

var oprPriorMap = map[rune]int{'+': 0, '-': 0, '*': 1, '/': 1}

// var oprFuncMap  =

func OprCalc(v1 interface{}, err1 error, v2 interface{}, err2 error, opr rune) (float64, error) {
	if _, empty := err1.(EmptyError); empty {
		return 0, ExprError{opr}
	}
	if _, empty := err2.(EmptyError); empty {
		return 0, ExprError{opr}
	}
	v := float64(0)
	switch opr {
	case '+':
		v = v2.(float64) + v1.(float64)
	case '-':
		v = v2.(float64) - v1.(float64)
	case '/':
		v = v2.(float64) / v1.(float64)
	case '*':
		v = v2.(float64) * v1.(float64)
	default:
		return 0, ExprError{opr}
	}
	return v, nil
}

//isprior 优先级大于不出栈
func isPrior(a, b rune) bool {
	return oprPriorMap[a] > oprPriorMap[b]
}

func isOpr(a rune) bool {
	_, exist := oprPriorMap[a]
	return exist
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

type EmptyError struct {
}

func (e EmptyError) Error() string {
	return "empty stack"
}

type ExprError struct {
	opr rune
}

func (e ExprError) Error() string {
	return "invalid symbol:" + string(e.opr)
}

type RandomStack struct {
	len   int
	elems map[int]interface{}
}
