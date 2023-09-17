package jasanya_response_backend

import (
	"fmt"
)

type HttpError struct {
	Message     StringErrorHttp
	Err         any
	Code        int
	CodeCompany int
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%v. %d", e.Message, e.Code)
}

func HttpErrString(m string, s StringErrorHttp) error {
	httpError := &HttpError{
		Message:     s,
		Err:         m,
		Code:        401,
		CodeCompany: 401,
	}

	return httpError
}

func HttpErrMapOfSlices(m map[string][]string, s StringErrorHttp) error {
	httpError := &HttpError{
		Message:     s,
		Err:         m,
		Code:        400,
		CodeCompany: 400,
	}

	return httpError
}
