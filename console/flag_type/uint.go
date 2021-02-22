package flag_type

import "strconv"

type Uint uint

func (i *Uint) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = Uint(v)
	return err
}

func (i *Uint) Get() interface{} {
	return uint(*i)
}

func (i *Uint) String() string { return strconv.FormatUint(uint64(*i), 10) }
