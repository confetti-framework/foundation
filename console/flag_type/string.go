package flag_type

import "fmt"

type String string

func (s *String) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *String) Set(value string) error {
	*s = String(value)
	return nil
}

func (s *String) Get() interface{} {
	return string(*s)
}
