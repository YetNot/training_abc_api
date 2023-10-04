package handlers

import (
	"abc_api/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetVars(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get var..")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
