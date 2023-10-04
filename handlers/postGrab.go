package handlers

import (
	"abc_api/models"
	"encoding/json"
	"log"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-type", "application/json")
}

func PostGrab(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Post new vars..")
	var vars models.Vars
	err := json.NewDecoder(request.Body).Decode(&vars)
	if err != nil {
		msg := models.Message{
			Message: "provided json file is invalid",
		}
		writer.WriteHeader(40)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(vars)
	vars.Check()
}
