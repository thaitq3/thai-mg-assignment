package server

import (
	"github.com/gorilla/mux"
	"word-counting/handler"
)

func routing(router *mux.Router) {
	// POST wager
	router.
		HandleFunc("/words", handler.GetHandler(&handler.WordHandler{})).
		Methods("POST")
}
