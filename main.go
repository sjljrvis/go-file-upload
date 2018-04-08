package main

import (
	"github.com/gorilla/mux"
	. "github.com/sjljrvis/go-file/controller"
	"log"
	"net/http"
)

func init() {
	log.Println("Server Started")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/file", UploadFile).Methods("POST")
	r.HandleFunc("/file", GetFile).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
