package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
)

func JsonResponse(w http.ResponseWriter, success bool, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		StatusCode: statusCode,
		Success:    success,
		Message:    message,
		Data:       data,
	})
}

func Unmarshall(r io.Reader, struct_ interface{}) error {
	structType := reflect.TypeOf(struct_)
	if structType.Kind() != reflect.Ptr || structType.Elem().Kind() != reflect.Struct {
		return errors.New("struct_ must be pointer to a struct")
	}

	err := json.NewDecoder(r).Decode(struct_)
	if err != nil {
		return err
	}
	return nil
}
