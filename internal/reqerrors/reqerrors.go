package reqerrors

type Error interface {
	Error() string
	StatusCode() int
}

type reqError struct {
	message    string
	statusCode int
}

func New(statusCode int, message string) Error {
	return &reqError{
		message:    message,
		statusCode: statusCode,
	}
}

func (e *reqError) Error() string {
	return e.message
}

func (e *reqError) StatusCode() int {
	return e.statusCode
}
