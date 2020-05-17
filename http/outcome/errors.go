package outcome

import "github.com/lanvard/contract/inter"

func Error(errors error) inter.Response {
	return &Response{}
}

func Errors(errors []error) inter.Response {
	return &Response{}
}
