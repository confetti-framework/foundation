package flag_type

import "time"

type Duration time.Duration

func (d *Duration) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		err = errParse
	}
	*d = Duration(v)
	return err
}

func (d *Duration) Get() interface{} {
	return time.Duration(*d)
}

func (d *Duration) String() string { return (*time.Duration)(d).String() }
