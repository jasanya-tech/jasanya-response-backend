package jasanya_response_backend

import (
	"testing"
)

type biodata struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func example(b bool) (*biodata, error) {
	if b {
		return nil, HttpErrString("postgres error", S500)
	}
	return &biodata{
		Name: "rama",
		Age:  15,
	}, nil
}

func TestError(t *testing.T) {
	_, err := example(true)

	res := ToResponseError(err)
	t.Logf("%#v", res)
}

func TestHttpMapOfSlices(t *testing.T) {
	err := HttpErrMapOfSlices(ErrorLoginMessage("email", "password"), S400)

	res := ToResponseError(err)
	t.Logf("%#v", res)
}

func TestRegisterErrMapOfSlices(t *testing.T) {
	errMap := RegisterErrMapOfSlices("payment_id", "invalid payment id")
	errMap2 := RegisterErrMapOfSlices("payment_id2", "invalid payment id2")

	err := HttpErrMapOfSlices(MergeMapOfSlices(errMap2, errMap), S400)

	res := ToResponseError(err)
	t.Logf("%#v", res)
}

func TestSucc(t *testing.T) {
	data, _ := example(false)

	res := ToResponseSuccess(data, "success created")
	t.Logf("%#v", res)
}
