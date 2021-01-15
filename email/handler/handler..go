package handler

import (
	"emailservice/email"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/send-email", email.SendOne).Methods("POST")
	return router
}
