package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GenError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationErr(errs validator.ValidationErrors) Response {
	var msgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			msgs = append(msgs, fmt.Sprintf("filed %s is required field ", err.Field()))

		default:
			msgs = append(msgs, fmt.Sprintf("filed %s is not valid ", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(msgs, ","),
	}
}
