package flag_type

import "strconv"

type Int int

func (i *Int) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = Int(v)
	return err
}

func (i *Int) Get() interface{} {return int(*i) }

func (i *Int) String() string { return strconv.Itoa(int(*i)) }
