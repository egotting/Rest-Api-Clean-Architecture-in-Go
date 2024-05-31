package exceptions

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field" `
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewNotFoundError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not found",
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal server error",
		Code:    http.StatusInternalServerError,
	}
}
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal server error",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}
