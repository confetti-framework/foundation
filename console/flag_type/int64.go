package flag_type

import "strconv"

type Int64 int64

func (i *Int64) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		err = numError(err)
	}
	*i = Int64(v)
	return err
}

func (i *Int64) Get() interface{} {
	return int64(*i)
}

func (i *Int64) String() string { return strconv.FormatInt(int64(*i), 10) }
