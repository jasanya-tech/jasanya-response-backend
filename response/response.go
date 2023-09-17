package response

import (
	"fmt"
	"net/http"
	"strconv"
)

type HttpResponse struct {
	Code    int    `json:"-"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}

type HttpError struct {
	Err         any
	CodeCompany CodeCompany
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%v. %s", CodeCompanyName[e.CodeCompany], e.CodeCompany)
}

const CM00 CodeCompany = "00"
const CM01 CodeCompany = "01"
const CM02 CodeCompany = "02"
const CM03 CodeCompany = "03"
const CM04 CodeCompany = "04"
const CM05 CodeCompany = "05"
const CM06 CodeCompany = "06"
const CM07 CodeCompany = "07"
const CM08 CodeCompany = "08"
const CM09 CodeCompany = "09"
const CM10 CodeCompany = "10"
const CM11 CodeCompany = "11"
const CM12 CodeCompany = "12"
const CM99 CodeCompany = "99"

type CodeCompany string

var CodeCompanyName = map[CodeCompany]string{
	"00": "success",
	"01": "not found",
	"02": "has ben deleted",
	"03": "required unique",
	"04": "Invalid Authorization",
	"05": "FORBIDDEN",
	"06": "BAD REQUEST",
	"07": "",
	"08": "TIMEOUT",
	"09": "BAD GATEWAY",
	"10": "SERVER DOWN",
	"11": "UNPROCESSABLE ENTITY",
	"12": "Invalid email OR password",
	"99": "SYSTEM ERROR",
}

func getCode(s CodeCompany) int {
	var code int
	switch s {
	case CM00:
		code = http.StatusOK
	case CM01:
		code = http.StatusNotFound
	case CM02:
		code = http.StatusNotFound
	case CM03:
		code = http.StatusBadRequest
	case CM04:
		code = http.StatusUnauthorized
	case CM05:
		code = http.StatusForbidden
	case CM06:
		code = http.StatusBadRequest
	case CM07:
		code = 0
	case CM08:
		code = http.StatusRequestTimeout
	case CM09:
		code = http.StatusBadGateway
	case CM10:
		code = http.StatusBadGateway
	case CM11:
		code = http.StatusUnprocessableEntity
	case CM12:
		code = http.StatusBadRequest
	case CM99:
		code = http.StatusInternalServerError
	default:
		code = http.StatusInternalServerError
	}

	return code
}

func Error(err HttpError) HttpResponse {
	codeCompanyInt, _ := strconv.Atoi(string(err.CodeCompany))
	code := getCode(err.CodeCompany)
	return HttpResponse{
		Code:    code,
		Status:  codeCompanyInt,
		Message: CodeCompanyName[err.CodeCompany],
		Errors:  err.Err,
		Data:    nil,
	}
}

func Success(data any, codeCompany CodeCompany) HttpResponse {
	code := getCode(codeCompany)
	codeCompanyInt, _ := strconv.Atoi(string(codeCompany))
	return HttpResponse{
		Code:    code,
		Status:  codeCompanyInt,
		Message: CodeCompanyName[codeCompany],
		Errors:  nil,
		Data:    data,
	}
}
