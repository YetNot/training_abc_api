package utils

import (
	"abc_api/handlers"

	"github.com/gorilla/mux"
)

func PostVars(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.PostGrab).Methods("POST")
}

func GetSolve(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetVars).Methods("GET")
}