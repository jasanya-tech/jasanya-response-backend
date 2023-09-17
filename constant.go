package jasanya_response_backend

type StringErrorHttp string

const S400 StringErrorHttp = "Invalid Request"
const S401 StringErrorHttp = "UNAUTHORIZED"
const S403 StringErrorHttp = "FORBIDDEN"
const S404 StringErrorHttp = "NOT FOUND"
const S408 StringErrorHttp = "REQUEST TIMEOUT"
const S422 StringErrorHttp = "unprocessable entity"
const S500 StringErrorHttp = "Internal Server Error"
