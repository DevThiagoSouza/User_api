package exeption

import (
	"net/http"
)

type Erro struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Erro) Error() string {
	return r.Message
}

func NewError(message, err string, code int, causes []Causes) *Erro {
	return &Erro{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequest(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidation(message string, causes []Causes) *Erro {
	return &Erro{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFound(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbidden(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
