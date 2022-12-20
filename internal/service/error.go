package service

type (
	sError struct{}
)

var (
	insError = sError{}
)

func Error() *sError {
	return &insError
}
