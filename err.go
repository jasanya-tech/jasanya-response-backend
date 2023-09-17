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
	code, codeCompany := getCode(s)
	httpError := &HttpError{
		Message:     s,
		Err:         m,
		Code:        code,
		CodeCompany: codeCompany,
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

func getCode(s StringErrorHttp) (int, int) {
	var code int
	var codeCompany int
	switch s {
	case S400:
		code = 400
		codeCompany = 400
	case S401:
		code = 401
		codeCompany = 401
	case S403:
		code = 403
		codeCompany = 403
	case S404:
		code = 404
		codeCompany = 404
	case S408:
		code = 408
		codeCompany = 408
	case S422:
		code = 422
		codeCompany = 422
	case S500:
		code = 500
		codeCompany = 500
	default:
		code = 500
		codeCompany = 500
	}

	return code, codeCompany
}
