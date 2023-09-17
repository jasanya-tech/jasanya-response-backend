package _error

import (
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
)

func HttpErrMapOfSlices(m map[string][]string, s response.CodeCompany) error {
	httpError := &response.HttpError{
		Err:         m,
		CodeCompany: s,
	}

	return httpError
}

func HttpErrString(m string, s response.CodeCompany) error {
	httpError := &response.HttpError{
		Err:         m,
		CodeCompany: s,
	}

	return httpError
}

//
// func Http401(m string) error {
// 	httpError := &HttpError{
// 		Message:     S401,
// 		Err:         m,
// 		Code:        401,
// 		CodeCompany: cc401,
// 	}
//
// 	return httpError
// }
//
// func Http403(m string) error {
// 	httpError := &HttpError{
// 		Message:     S403,
// 		Err:         m,
// 		Code:        403,
// 		CodeCompany: cc403,
// 	}
//
// 	return httpError
// }
//
// func Http404(m string) error {
// 	httpError := &HttpError{
// 		Message:     S404,
// 		Err:         m,
// 		Code:        404,
// 		CodeCompany: cc404,
// 	}
//
// 	return httpError
// }
//
// func Http408(m string) error {
// 	httpError := &HttpError{
// 		Message:     S408,
// 		Err:         m,
// 		Code:        408,
// 		CodeCompany: cc408,
// 	}
//
// 	return httpError
// }
//
// func Http422(m string) error {
// 	httpError := &HttpError{
// 		Message:     S422,
// 		Err:         m,
// 		Code:        422,
// 		CodeCompany: cc422,
// 	}
//
// 	return httpError
// }
//
// func Http500(m string) error {
// 	httpError := &HttpError{
// 		Message:     S500,
// 		Err:         m,
// 		Code:        500,
// 		CodeCompany: cc500,
// 	}
//
// 	return httpError
// }
