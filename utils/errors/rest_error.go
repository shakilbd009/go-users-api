package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Error:   "bad_request",
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Error:   "not_found",
		Status:  http.StatusNotFound,
	}
}
