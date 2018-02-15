package reflect

import (
	"fmt"
)

type ErrCanNotConvertType struct {
	value interface{}
	from  Type
	to    Type
}

func (e ErrCanNotConvertType) Error() string {
	return fmt.Sprintf(
		"Can not convert '%#v' of type '%s' to '%s'",
		e.value,
		e.from,
		e.to,
	)
}

func NewErrCanNotConvertType(value interface{}, from Type, to Type) ErrCanNotConvertType {
	return ErrCanNotConvertType{
		value: value,
		from:  from,
		to:    to,
	}
}
