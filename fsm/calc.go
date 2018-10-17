package fsm

type calculator struct {
	opr    *Stack
	num    *Stack
	digits *Stack
}

func NewCalc() *calculator {
	return &calculator{
		opr:    new(Stack).Init(""),
		num:    new(Stack).Init(""),
		digits: new(Stack).Init(""),
	}
}

//digitsToNum execute before every action
func (c *calculator) digitsToNum(curr, next interface{}) error {
	if isDigit(next.(rune)) {
		return nil
	}
	l := c.digits.Len()
	num := 0
	for i := 0; i < l; i++ {
		dgt_, _ := c.digits.Pop()
		dgt := dgt_.(int)
		num += dgt * (1 << uint(i))
	}
	c.num.Push(float64(num))
	return nil
}

func (c *calculator) getOpr(curr, next interface{}) error {
	oprIn := next.(rune)
	for {
		oprTop_, err := c.opr.Peek()
		if _, empty := err.(EmptyError); empty {
			c.opr.Push(oprIn)
			return nil
		}

		oprTop := oprTop_.(rune)

		if !isOpr(oprTop) {
			c.opr.Push(oprIn)
			return nil
		}

		if isPrior(oprIn, oprTop) {
			c.opr.Push(oprIn)
			return nil
		}

		c.opr.Pop()

		//
		v1, err1 := c.num.Pop()
		v2, err2 := c.num.Pop()
		v, err := oprCalc(v1, err1, v2, err2, oprTop)
		if err != nil {
			return err
		}
		c.num.Push(v)
	}
	c.opr.Push(oprIn)
	return nil
}

func (c *calculator) getDigit(curr, next interface{}) error {
	dgt := next.(rune) - '0'
	c.digits.Push(dgt)
	return nil
}

func (c *calculator) getLBracket(curr, next interface{}) error {
	c.opr.Push(next)
	return nil
}

func (c *calculator) getRBracket(curr, next interface{}) error {
	for {
		opr_, err := c.opr.Pop()
		if _, empty := err.(EmptyError); empty {
			return ExprError{')'}
		}

		oprTop := opr_.(rune)
		if oprTop == '(' {
			return nil
		}

		if !isOpr(oprTop) {
			return ExprError{oprTop}
		}

		v1, err1 := c.num.Pop()
		v2, err2 := c.num.Pop()
		v, err := oprCalc(v1, err1, v2, err2, oprTop)
		if err != nil {
			return err
		}
		c.num.Push(v)
	}
}

func (c *calculator) getSpace(curr, next interface{}) {
	return
}

func (c *calculator) Result() (float64, error) {
	for !c.opr.Empty() {
		opr_, _ := c.opr.Pop()
		opr := opr_.(rune)

		v1, err1 := c.num.Pop()
		v2, err2 := c.num.Pop()
		v, err := oprCalc(v1, err1, v2, err2, opr) // v2 <opr> v1
		if err != nil {
			return 0, err
		}
		c.num.Push(v)
	}
	re, err := c.num.Pop()
	if err != nil {
		return 0, err
	}
	return re.(float64), nil
}

func calc(expr string) (float64, error) {
	r := NewRunner()
	c := NewCalc()

	r.regEvent()

	for _, e := range expr {
		r.Handle(e)
	}

	return c.Result()
}
