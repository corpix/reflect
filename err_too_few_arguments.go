package reflect

import (
	"fmt"
)

type ErrTooFewArguments struct {
	Want int
	Got  int
}

func (e *ErrTooFewArguments) Error() string {
	return fmt.Sprintf(
		"Too few arguments, want '%d', got '%d'",
		e.Want,
		e.Got,
	)
}

func NewErrTooFewArguments(want int, got int) *ErrTooFewArguments {
	return &ErrTooFewArguments{
		Want: want,
		Got:  got,
	}
}
