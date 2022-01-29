package handler

import (
	"net/http"

	"github.com/golobby/container/v3"

	"word-counting/utils"
)

func GetHandler(handler Handler) func(http.ResponseWriter,
	*http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := container.Fill(handler)
		if err != nil {
			utils.OutputInternalServerError(writer, err)
		} else {
			handler.Handle(writer, request)
		}
	}
}
