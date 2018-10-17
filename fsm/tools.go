package fsm

var oprPriorMap = map[rune]int{'+': 0, '-': 0, '*': 1, '/': 1}

// var oprFuncMap  =

func oprCalc(v1 interface{}, err1 error, v2 interface{}, err2 error, opr rune) (float64, error) {
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
