package jasanya_response_backend

import (
	"errors"
)

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}

func ToResponseError(err error) HttpResponse {
	var errHTTP *HttpError

	ok := errors.As(err, &errHTTP)
	if !ok {
		err = HttpErrString("postgres error", S500)
		errors.As(err, &errHTTP)
	}

	httpResp := HttpResponse{
		Status:  errHTTP.CodeCompany,
		Message: string(errHTTP.Message),
		Errors:  errHTTP.Err,
		Data:    nil,
	}
	return httpResp
}

func ToResponseSuccess(data any, message string) HttpResponse {

	httpResp := HttpResponse{
		Status:  200,
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	return httpResp
}
