package flag_type

import "strconv"

type Uint64 uint64

func (i *Uint64) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		err = numError(err)
	}
	*i = Uint64(v)
	return err
}

func (i *Uint64) Get() interface{} {
	return uint64(*i)
}

func (i *Uint64) String() string { return strconv.FormatUint(uint64(*i), 10) }
