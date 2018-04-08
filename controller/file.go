package controller

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// File struct
type File struct {
	Name string
	Size int
	URL  string
}

// Files Array
type Files []File

// Response struct
type Response struct {
	Status  bool   `bson:"status" json:"status"`
	Message string `bson:"message" json:"message"`
}

// UploadFile : Upload a file to container
func UploadFile(w http.ResponseWriter, r *http.Request) {
	f, fh, err := r.FormFile("file")
	if err != nil {
		log.Fatal("Error uploading file")
		respondWithError(w, http.StatusBadRequest, "Error uploading file")
	}
	defer f.Close()
	fileExt := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, f)
	fname := strings.Split(fh.Filename, ".")[0] + "." + fileExt
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "files", fname)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	f.Seek(0, 0)
	io.Copy(nf, f)
	respondWithJSON(w, http.StatusOK, Response{true, "File uploaded successfully"})
}

// GetFile :-GET request to get all files in container respective to user
func GetFile(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Header.Get("Authorization"))
	respondWithJSON(w, http.StatusOK, Response{true, "File uploaded successfully"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})

}
