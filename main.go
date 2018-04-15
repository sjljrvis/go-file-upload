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
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
