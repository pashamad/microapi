package handler

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	id      string
	code    int32
	message string
	detail  string
	hint    string
}

type MicroError struct {
	Id     string
	Code   int32
	Detail string
	Status string
}

func (e *MicroError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func New(id, detail string, code int32) error {
	return &Error{
		id:      id,
		code:    code,
		message: http.StatusText(int(code)),
		detail:  detail,
		hint:    "err.unknown",
	}
}

func FromMicroError(err MicroError) *Error {
	return &Error{
		id:      err.Id,
		code:    err.Code,
		message: err.Status,
		detail:  err.Detail,
		hint:    "err.unknown",
	}
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	// todo
	if verr, ok := err.(*MicroError); ok && verr != nil {
		return FromMicroError(*verr)
	}

	return Parse(err.Error())
}

func Parse(err string) *Error {
	me := new(MicroError)
	errr := json.Unmarshal([]byte(err), me)
	if errr != nil {
		me.Detail = err
	}

	return FromMicroError(*me)
}
