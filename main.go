package main

import (
	"abc_api/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	prefix string = "/api/v1"
)

var (
	port         string
	postGrabVars string = prefix + "/grab"
	getSolve     string = prefix + "/solve"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env")
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server ON:", port)
	router := mux.NewRouter()
	utils.PostVars(router, postGrabVars)
	utils.GetSolve(router, getSolve)
	log.Fatal(http.ListenAndServe(":"+port, router))
}