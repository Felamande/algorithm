package fsm

import (
	"fmt"
)

type event struct {
	curr, next interface{}
}
type actionFn func(interface{}, interface{}) error
type runner struct {
	curr      interface{}
	preProcs  []actionFn
	postProcs []actionFn
	stMap     map[event]actionFn
}

func NewRunner() *runner {
	return &runner{
		stMap: make(map[event]actionFn),
	}
}

func (r *runner) regEvent(curr, next interface{}, f actionFn) {
	r.stMap[event{curr, next}] = f
}

func (r *runner) Handle(state interface{}) error {
	for _, f := range r.preProcs {
		err := f(r.curr, state)
		if err != nil {
			if _, isContinue := err.(Continue); isContinue {
				return nil
			}
			return err
		}
	}

	if f, exist := r.stMap[event{r.curr, state}]; exist {
		err := f(r.curr, state)
		if err != nil {
			return err
		}
	} else {
		return UknownStateError{state}
	}

	for _, f := range r.postProcs {
		err := f(r.curr, state)
		if err != nil {
			if _, isContinue := err.(Continue); isContinue {
				return nil
			}
			return err
		}
	}

	r.curr = state
	return nil
}

// func (r *runner) Result() (interface{}, error) {

// }

type UknownStateError struct {
	curr interface{}
}

func (e UknownStateError) Error() string {
	return fmt.Sprintf("Unknown state:%v", e.curr)
}

type Continue struct{}

func (e Continue) Error() string {
	return ""
}
