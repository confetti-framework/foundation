package flag_type

import (
	"strconv"
)

type Bool bool

func (b *Bool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		err = errParse
	}
	*b = Bool(v)
	return err
}

func (b *Bool) Get() interface{} {
	return bool(*b)
}

func (b *Bool) String() string {
	return strconv.FormatBool(bool(*b))
}

func (b *Bool) IsBoolFlag() bool {
	return true
}
