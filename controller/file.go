package controller

import (
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
	fileObj, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Fatal("Error uploading file")
		respondWithError(w, http.StatusBadRequest, "Error uploading file")
	}
	defer fileObj.Close()
	fileExt := strings.Split(fileHeader.Filename, ".")[1]
	fname := strings.Split(fileHeader.Filename, ".")[0] + "." + fileExt
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "files", fname)
	newFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()
	fileObj.Seek(0, 0)
	io.Copy(newFile, fileObj)
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
