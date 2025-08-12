package resterr

import "net/http"

type Resterr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Resterr) Error() string {
	return r.Message
}

func NewResterr(message, err string, code int, causes []Causes) *Resterr {
	return &Resterr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *Resterr {
	return &Resterr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Resterr {
	return &Resterr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *Resterr{
	return &Resterr{
		Message: message,
		Err: "internal server error",
		Code: http.StatusInternalServerError,
	}
}

func NewForbiddenError(message string) *Resterr {
	return &Resterr{
		Message: message,
		Err: "forbidden",	
		Code: http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *Resterr {
     return &Resterr{
		Message: message,
		Err: "not found",
		Code: http.StatusNotFound,
	 }
}

func NewMethodNotAllowed(message string) *Resterr {
	return &Resterr{
		Message: message,
		Err: "not Allowed Method",
		Code: http.StatusMethodNotAllowed,
	}
}