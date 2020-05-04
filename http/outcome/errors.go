package outcome

func Error(errors error) Response {
	return Response{}
}

func Errors(errors []error) Response {
	return Response{}
}
